/**
 * Copyright (c) 2016-present Snowplow Analytics Ltd. All rights reserved.
 *
 * This software is made available by Snowplow Analytics, Ltd.,
 * under the terms of the Snowplow Limited Use License Agreement, Version 1.0
 * located at https://docs.snowplow.io/limited-use-license-1.0
 * BY INSTALLING, DOWNLOADING, ACCESSING, USING OR DISTRIBUTING ANY PORTION
 * OF THE SOFTWARE, YOU AGREE TO THE TERMS OF SUCH LICENSE AGREEMENT.
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
