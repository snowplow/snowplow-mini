/**
 * Copyright (c) 2016-present Snowplow Analytics Ltd. All rights reserved.
 *
 * This software is made available by Snowplow Analytics, Ltd.,
 * under the terms of the Snowplow Limited Use License Agreement, Version 1.1
 * located at https://docs.snowplow.io/limited-use-license-1.1
 * BY INSTALLING, DOWNLOADING, ACCESSING, USING OR DISTRIBUTING ANY PORTION
 * OF THE SOFTWARE, YOU AGREE TO THE TERMS OF SUCH LICENSE AGREEMENT.
 */

package main

import (
	"encoding/json"
	"gopkg.in/pg.v5"
	"io/ioutil"
	"strings"
)

type PsqlInfos struct {
	User     string
	Password string
	Database string
	Addr     string
}

type LocalIglu struct {
	ConfigPath string
	IgluApikey string
	Psql       PsqlInfos
}

func (li LocalIglu) addApiKeyToConfig() error {

	jsonFile, err := ioutil.ReadFile(li.ConfigPath)
	if err != nil {
		return err
	}

	var igluConf IgluConf
	json.Unmarshal(jsonFile, &igluConf)

	for i, repo := range igluConf.Data.Repos {
		igluUri := repo.Conn.Http["uri"]
		if strings.Contains(igluUri, "iglu-server") {
			igluConf.Data.Repos[i].Conn.Http["apikey"] = li.IgluApikey
		}
	}

	jsonBytes, err := json.MarshalIndent(igluConf, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(li.ConfigPath, jsonBytes, 0644)
}

func (li LocalIglu) insertApiKeyToDb() error {

	db := pg.Connect(&pg.Options{
		User:     li.Psql.User,
		Password: li.Psql.Password,
		Database: li.Psql.Database,
		Addr:     li.Psql.Addr,
	})
	defer db.Close()

	_, err := db.Exec("DELETE FROM iglu_permissions")
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO iglu_permissions " +
		"VALUES " +
		"('" + li.IgluApikey + "', '', TRUE, 'CREATE_VENDOR'::schema_action, '{\"CREATE\", \"DELETE\"}'::key_action[])")

	if err != nil {
		return err
	}

	return nil
}

func (li LocalIglu) addApiKey() error {
	err := li.addApiKeyToConfig()
	if err != nil {
		return err
	}

	err = li.insertApiKeyToDb()
	if err != nil {
		return err
	}

	return nil
}
