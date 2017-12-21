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
		if strings.Contains(igluUri, "localhost") ||
			strings.Contains(igluUri, "127.0.0.1") {

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

	_, err := db.Exec("DELETE FROM apikeys")
	if err != nil {
		return err
	}

	_, err = db.Exec("INSERT INTO apikeys " +
		"(uid, vendor_prefix, permission, createdat) " +
		"VALUES " +
		"('" + li.IgluApikey + "','*','super',current_timestamp)")

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
