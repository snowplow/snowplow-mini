/**
 * Copyright (c) 2016-2017 Snowplow Analytics Ltd.
 * All rights reserved.
 *
 * This program is licensed to you under the Apache License Version 2.0,
 * and you may not use this file except in compliance with the Apache
 * License Version 2.0.
 * You may obtain a copy of the Apache License Version 2.0 at
 * http://www.apache.org/licenses/LICENSE-2.0.
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the Apache License Version 2.0 is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied.
 *
 * See the Apache License Version 2.0 for the specific language
 * governing permissions and limitations there under.
 */

package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

var configPath string
var config ControlPlaneConfig

func main() {
	configFlag := flag.String("config", "", "Control Plane API config file")
	flag.Parse()
	configPath = *configFlag

	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		panic(err)
	}

	http.HandleFunc("/restart-services", restartServices)
	http.HandleFunc("/enrichments", uploadEnrichments)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func restartServices(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "PUT" {
		err := restartSPServices()
		if err != nil {
			http.Error(resp, err.Error(), 500)
		} else {
			resp.WriteHeader(http.StatusOK)
			io.WriteString(resp, "OK")
		}
	} else {
		// Return 404 for other methods
		http.Error(resp, "", 404)
	}
}

func uploadEnrichments(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// maxMemory bytes of body's file parts are stored in memory,
		// with the remainder stored on disk in temporary files
		var maxMemory int64 = 32 << 20
		err := req.ParseMultipartForm(maxMemory)
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}
		file, handler, err := req.FormFile("enrichmentjson")
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}
		defer file.Close()

		fileContentBytes, err := ioutil.ReadAll(file)
		fileContent := string(fileContentBytes)

		if !isJSON(fileContent) {
			http.Error(resp, "JSON is not valid", 400)
			return
		}

		f, err := os.OpenFile(config.Dirs.Enrichments+"/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}
		defer f.Close()

		io.WriteString(f, fileContent)

		err = restartService("streamEnrich")
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}

		resp.WriteHeader(http.StatusOK)
		io.WriteString(resp, "uploaded successfully")
	} else {
		http.Error(resp, "", 404)
	}
}
