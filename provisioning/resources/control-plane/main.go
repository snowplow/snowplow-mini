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
	http.HandleFunc("/external-iglu", addExternalIgluServer)
	http.HandleFunc("/local-iglu-apikey", addLocalIgluApikey)
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

func addExternalIgluServer(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()

		vendorPrefixArr, checkVendor := req.Form["vendor_prefix"]
		uriArr, checkUri := req.Form["uri"]
		apikeyArr, checkApikey := req.Form["apikey"]
		nameArr, checkName := req.Form["name"]
		priorityArr, checkPriority := req.Form["priority"]
		if !(checkVendor && checkUri && checkApikey && checkName && checkPriority) {
			http.Error(resp, "missing parameter", 400)
			return
		}
		uri := uriArr[0]
		apikey := apikeyArr[0]
		vendorPrefix := vendorPrefixArr[0]
		name := nameArr[0]
		priority, err := strconv.Atoi(priorityArr[0])
		if err != nil {
			http.Error(resp, "priority must be integer", 400)
			return
		}

		if !isUrlReachable(uri) {
			http.Error(resp, "Given URL is not reachable", 400)
			return
		}
		if !isValidUuid(apikey) {
			http.Error(resp, "Given apikey is not valid UUID.", 400)
			return
		}

		externalIgluServer := ExternalIgluServer{
			ConfigPath: config.Dirs.Config + "/" +
				config.ConfigNames.IgluResolver,
			IgluInfo: IgluInfo{
				VendorPrefix: vendorPrefix,
				Uri:          uri,
				Apikey:       apikey,
				Name:         name,
				Priority:     priority,
			},
		}

		err = externalIgluServer.addExternalIgluServer()
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}

		err = restartService("streamEnrich")
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}
		resp.WriteHeader(http.StatusOK)
		io.WriteString(resp, "added successfully")
	} else {
		http.Error(resp, "", 404)
	}
}

func addLocalIgluApikey(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()

		igluApikeyArr, checkApikey := req.Form["local_iglu_apikey"]
		if !checkApikey {
			http.Error(resp, "missing parameter", 400)
			return
		}
		igluApikey := igluApikeyArr[0]

		if !isValidUuid(igluApikey) {
			http.Error(resp, "Given apikey is not valid UUID", 400)
			return
		}

		psqlInfos := PsqlInfos{
			User:     config.Psql.User,
			Password: config.Psql.Password,
			Database: config.Psql.Database,
			Addr:     config.Psql.Addr,
		}

		localIglu := LocalIglu{
			ConfigPath: config.Dirs.Config + "/" +
				config.ConfigNames.IgluResolver,
			IgluApikey: igluApikey,
			Psql:       psqlInfos,
		}
		err := localIglu.addApiKey()
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}

		err = restartService("streamEnrich")
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}
		resp.WriteHeader(http.StatusOK)
		io.WriteString(resp, "added successfully")
	} else {
		http.Error(resp, "", 404)
	}
}
