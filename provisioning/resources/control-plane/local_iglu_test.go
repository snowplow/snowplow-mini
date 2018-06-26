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
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestAddApiKeyToConfig(t *testing.T) {
	assert := assert.New(t)

	igluConfigOne :=
		`{
  "schema": "iglu:com.snowplowanalytics.iglu/resolver-config/jsonschema/1-0-1",
  "data": {
    "cacheSize": 500,
    "repositories": [
      {
        "name": "Iglu Server",
        "priority": 0,
        "vendorPrefixes": [
          "com.snowplowanalytics"
        ],
        "connection": {
          "http": {
            "uri": "http://localhost:8081/api",
            "apikey": "PLACEHOLDER"
          }
        }
      }
    ]
  }
}`
	expectedIgluConfigOne :=
		`{
  "schema": "iglu:com.snowplowanalytics.iglu/resolver-config/jsonschema/1-0-1",
  "data": {
    "cacheSize": 500,
    "repositories": [
      {
        "name": "Iglu Server",
        "priority": 0,
        "vendorPrefixes": [
          "com.snowplowanalytics"
        ],
        "connection": {
          "http": {
            "uri": "http://localhost:8081/api",
            "apikey": "iglu_apikey"
          }
        }
      }
    ]
  }
}`

	igluConfigTwo :=
		`{
  "schema": "iglu:com.snowplowanalytics.iglu/resolver-config/jsonschema/1-0-1",
  "data": {
    "cacheSize": 500,
    "repositories": [
      {
        "name": "Iglu Server",
        "priority": 0,
        "vendorPrefixes": [
          "com.snowplowanalytics"
        ],
        "connection": {
          "http": {
            "uri": "http://localhost:8081/api"
          }
        }
      }
    ]
  }
}`

	expectedIgluConfigTwo :=
		`{
  "schema": "iglu:com.snowplowanalytics.iglu/resolver-config/jsonschema/1-0-1",
  "data": {
    "cacheSize": 500,
    "repositories": [
      {
        "name": "Iglu Server",
        "priority": 0,
        "vendorPrefixes": [
          "com.snowplowanalytics"
        ],
        "connection": {
          "http": {
            "uri": "http://localhost:8081/api",
            "apikey": "iglu_apikey"
          }
        }
      }
    ]
  }
}`

	dir, err := ioutil.TempDir("", "testDir")
	assert.Nil(err)

	defer os.RemoveAll(dir)

	tmpfn := filepath.Join(dir, "tmpfile")

	localIglu := LocalIglu{
		ConfigPath: tmpfn,
		IgluApikey: "iglu_apikey",
		Psql:       PsqlInfos{},
	}

	err = ioutil.WriteFile(tmpfn, []byte(igluConfigOne), 0666)
	assert.Nil(err)

	err = localIglu.addApiKeyToConfig()
	assert.Nil(err)

	afterInsert, err := ioutil.ReadFile(tmpfn)
	assert.Nil(err)

	assert.True(string(afterInsert) == expectedIgluConfigOne)

	err = ioutil.WriteFile(tmpfn, []byte(igluConfigTwo), 0666)
	assert.Nil(err)

	err = localIglu.addApiKeyToConfig()
	assert.Nil(err)

	afterInsert, err = ioutil.ReadFile(tmpfn)
	assert.Nil(err)

	assert.True(string(afterInsert) == expectedIgluConfigTwo)
}
