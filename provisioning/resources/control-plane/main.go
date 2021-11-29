/**
 * Copyright (c) 2016-2018 Snowplow Analytics Ltd.
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
	"errors"
	"flag"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/BurntSushi/toml"
)

var configPath string
var config ControlPlaneConfig

func main() {
	configFlag := flag.String("config", "/home/ubuntu/snowplow/configs/control-plane-api.toml",
		"Control Plane API config file")
	flag.Parse()
	configPath = *configFlag

	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		panic(err)
	}

	http.HandleFunc("/restart-services", restartServices)
	http.HandleFunc("/restart-service", restartService)
	http.HandleFunc("/enrichments", uploadEnrichments)
	http.HandleFunc("/iglu-config", uploadIgluConfig)
	http.HandleFunc("/external-iglu", addExternalIgluServer)
	http.HandleFunc("/local-iglu-apikey", addLocalIgluApikey)
	http.HandleFunc("/credentials", changeUsernameAndPassword)
	http.HandleFunc("/domain-name", addDomainName)
	http.HandleFunc("/version", getSpminiVersion)
	http.HandleFunc("/telemetry", manageTelemetry)
	http.HandleFunc("/reset-service", resetService)
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

func restartService(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "PUT" {
		req.ParseForm()
		if serviceNameArr, ok := req.Form["service_name"]; ok {
			err, status := restartSPService(serviceNameArr[0])
			if err != nil {
				http.Error(resp, err.Error(), status)
			} else {
				resp.WriteHeader(http.StatusOK)
				io.WriteString(resp, "OK")
			}
		} else {
			http.Error(resp, "Missing key service_name", 400)
			return
		}
	} else {
		http.Error(resp, "Only PUT is supported", 404)
		return
	}
}

func resetService(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		err := req.ParseForm()
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}
		if serviceNameArr, ok := req.Form["service_name"]; ok {
			serviceName := serviceNameArr[0]
			if serviceName == "elasticsearch" {
				err = resetElasticsearch("http://localhost:9200/_all")
				if err != nil {
					http.Error(resp, err.Error(), 500)
					return
				}
				err = createESIndices()
				if err != nil {
					http.Error(resp, err.Error(), 500)
					return
				}
				err = createKibanaIndexPatterns()
				if err != nil {
					http.Error(resp, err.Error(), 500)
					return
				}
				resp.WriteHeader(http.StatusOK)
				io.WriteString(resp, "Both Elasticsearch & Kibana are reset including the data and index mappings")
			} else {
				http.Error(resp, serviceName+" can't be reset. Only elasticsearch can be reset.", 400)
				return
			}
		} else {
			http.Error(resp, "Missing key service_name", 400)
			return
		}
	} else {
		http.Error(resp, "Only POST is supported", 400)
		return
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

		// Uploaded enrichment can be shorter than the existing one
		// Truncating to 0 bytes and seeking I/O offset to the beginning
		// Prevents the possibility of corrupted json
		f.Truncate(0)
		f.Seek(0, 0)
		// Now we can write to file in peace
		io.WriteString(f, fileContent)

		err, status := restartSPService("enrich")
		if err != nil {
			http.Error(resp, err.Error(), status)
			return
		}

		resp.WriteHeader(http.StatusOK)
		io.WriteString(resp, "uploaded successfully")
	} else {
		http.Error(resp, "", 404)
	}
}

func uploadIgluConfig(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		// maxMemory bytes of body's file parts are stored in memory,
		// with the remainder stored on disk in temporary files
		var maxMemory int64 = 32 << 20
		err := req.ParseMultipartForm(maxMemory)

		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}

		file, _, err := req.FormFile("igluserverhocon")
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}
		defer file.Close()

		fileContentBytes, err := ioutil.ReadAll(file)
		fileContent := string(fileContentBytes)
		f, err := os.OpenFile(config.Dirs.Config+"/"+config.ConfigNames.IgluServer, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}
		defer f.Close()

		// Uploaded Iglu Server configuration can be shorter than existing one
		// Which would make iglu server configuration invalid
		// Truncating to 0 bytes and seeking I/O offset to the beginning
		// Prevents that possibility
		f.Truncate(0)
		f.Seek(0, 0)
		// Now we can write to config file in peace
		io.WriteString(f, fileContent)

		err, status := restartSPService("iglu")
		if err != nil {
			http.Error(resp, err.Error(), status)
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
		if !(checkVendor && checkUri && checkName && checkPriority) {
			http.Error(resp, "missing parameter", 400)
			return
		}
		uri := uriArr[0]
		apikey := ""
		if checkApikey {
			apikey = apikeyArr[0]
		}
		vendorPrefix := vendorPrefixArr[0]
		name := nameArr[0]
		priority, err := strconv.Atoi(priorityArr[0])
		if err != nil {
			http.Error(resp, "Priority must be an integer", 400)
			return
		}

		if !isURLReachable(uri) {
			http.Error(resp, "Given URL is not reachable", 400)
			return
		}
		if apikey != "" && !isValidUUID(apikey) {
			http.Error(resp, "Given apikey is not a valid UUID.", 400)
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

		err, status := restartSPService("enrich")
		if err != nil {
			http.Error(resp, err.Error(), status)
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

		if !isValidUUID(igluApikey) {
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

		err, status := restartSPService("enrich")
		if err != nil {
			http.Error(resp, err.Error(), status)
			return
		}
		resp.WriteHeader(http.StatusOK)
		io.WriteString(resp, "added successfully")
	} else {
		http.Error(resp, "", 404)
	}
}

func changeUsernameAndPassword(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		err := req.ParseForm()
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}

		var newUsername string
		var newPassword string

		if newUsernameArr, ok := req.Form["new_username"]; ok {
			newUsername = newUsernameArr[0]
		} else {
			http.Error(resp, "missing parameter new_username", 400)
			return
		}

		if newPasswordArr, ok := req.Form["new_password"]; ok {
			newPassword = newPasswordArr[0]
		} else {
			http.Error(resp, "missing parameter new_password", 400)
			return
		}

		err, status := changeCredentials(
			config.ConfigNames.Caddy,
			newUsername,
			newPassword,
		)
		if err != nil {
			http.Error(resp, err.Error(), status)
			return
		}
		err, status = restartSPService("caddy")
		if err != nil {
			http.Error(resp, err.Error(), status)
			return
		}
		resp.WriteHeader(http.StatusOK)
		io.WriteString(resp, "changed successfully")
	} else {
		http.Error(resp, "", 404)
	}
}

func addDomainName(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		req.ParseForm()

		domainNameArr, checkDomainName := req.Form["domain_name"]
		if !checkDomainName {
			http.Error(resp, "missing parameter", 400)
			return
		}
		domainName := domainNameArr[0]

		err := checkHostDomainName(domainName)
		if err != nil {
			http.Error(resp, err.Error(), 405)
			return
		}

		err = changeDomainName(
			config.Dirs.Config+"/"+config.ConfigNames.Caddy,
			domainName,
		)
		if err != nil {
			http.Error(resp, err.Error(), 405)
			return
		}

		err, status := restartSPService("caddy")
		if err != nil {
			http.Error(resp, err.Error(), status)
			return
		}

		resp.WriteHeader(http.StatusOK)
		io.WriteString(resp, "added successfully")
	} else {
		http.Error(resp, "", 404)
	}
}

func getSpminiVersion(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		versionBytes, err := ioutil.ReadFile(config.VersionFilePath)
		if err != nil {
			http.Error(resp, err.Error(), 500)
			return
		}
		versionStr := string(versionBytes)
		resp.WriteHeader(http.StatusOK)
		io.WriteString(resp, versionStr)
	} else {
		http.Error(resp, "", 404)
	}
}

func manageTelemetry(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		telemetryDisable, err := getTelemetryDisable(config.Dirs.Config + "/" + config.ConfigNames.Collector)
		handleError(resp, err, http.StatusInternalServerError)

		resp.WriteHeader(http.StatusOK)
		_, err = io.WriteString(resp, "{\"collector.telemetry.disable\": "+telemetryDisable+"}")
		handleError(resp, err, http.StatusInternalServerError)
	} else if req.Method == "PUT" {
		err := req.ParseForm()
		handleError(resp, err, http.StatusInternalServerError)

		disable := req.PostFormValue("disable")
		if disable == "" {
			handleError(resp, errors.New("no key named disable"), http.StatusBadRequest)
		} else if disable == "true" || disable == "false" {
			err = setTelemetryDisable(config.Dirs.Config+"/"+config.ConfigNames.Collector, disable)
			handleError(resp, err, http.StatusInternalServerError)
			err, status := restartSPService("collector")
			handleError(resp, err, status)
		} else {
			handleError(resp, errors.New("set disable key to either true or false"), http.StatusBadRequest)
		}
	} else {
		handleError(resp, errors.New("http method not supported"), http.StatusMethodNotAllowed)
	}
}

func handleError(w http.ResponseWriter, err error, status int) {
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
}
