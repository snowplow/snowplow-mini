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
	"encoding/json"
	"io/ioutil"
)

type IgluInfo struct {
	Name         string
	Priority     int
	VendorPrefix string
	Uri          string
	Apikey       string
}

type ExternalIgluServer struct {
	ConfigPath string
	IgluInfo
}

func (e ExternalIgluServer) addExternalIgluServer() error {

	jsonFile, err := ioutil.ReadFile(e.ConfigPath)
	if err != nil {
		return err
	}

	var httpConf map[string]string
	if e.IgluInfo.Apikey != "" {
		httpConf = map[string]string{"uri": e.IgluInfo.Uri, "apikey": e.IgluInfo.Apikey}
	} else {
		httpConf = map[string]string{"uri": e.IgluInfo.Uri}
	}

	newIgluRepo := RepoConf{
		Name:           e.IgluInfo.Name,
		Priority:       e.IgluInfo.Priority,
		VendorPrefixes: []string{e.IgluInfo.VendorPrefix},
		Conn: ConnectionConf{
			Http: httpConf,
		},
	}

	var igluConf IgluConf
	json.Unmarshal(jsonFile, &igluConf)

	igluConf.Data.Repos = append(igluConf.Data.Repos, newIgluRepo)
	jsonBytes, err := json.MarshalIndent(igluConf, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(e.ConfigPath, jsonBytes, 0644)
}
