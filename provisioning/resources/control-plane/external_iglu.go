/**
 * Copyright (c) 2016-present Snowplow Analytics Ltd. All rights reserved.
 *
 * This program is licensed to you under the Snowplow Community License Version 1.0,
 * and you may not use this file except in compliance with the Snowplow Community License Version 1.0.
 * You may obtain a copy of the Snowplow Community License Version 1.0 at https://docs.snowplow.io/community-license-1.0
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
