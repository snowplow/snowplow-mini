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
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddApiKeyToConfig(t *testing.T) {
	assert := assert.New(t)

	igluConfigOne :=
		`{
  "schema": "iglu:com.snowplowanalytics.iglu/resolver-config/jsonschema/1-0-1",
  "data": {
    "cacheSize": 500,
    "cacheTtl": 1,
    "repositories": [
      {
        "name": "Iglu Server",
        "priority": 0,
        "vendorPrefixes": [
          "com.snowplowanalytics"
        ],
        "connection": {
          "http": {
            "uri": "http://iglu-server:8081/api",
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
    "cacheTtl": 1,
    "repositories": [
      {
        "name": "Iglu Server",
        "priority": 0,
        "vendorPrefixes": [
          "com.snowplowanalytics"
        ],
        "connection": {
          "http": {
            "uri": "http://iglu-server:8081/api",
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
    "cacheTtl": 1,
    "repositories": [
      {
        "name": "Iglu Server",
        "priority": 0,
        "vendorPrefixes": [
          "com.snowplowanalytics"
        ],
        "connection": {
          "http": {
            "uri": "http://iglu-server:8081/api"
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
    "cacheTtl": 1,
    "repositories": [
      {
        "name": "Iglu Server",
        "priority": 0,
        "vendorPrefixes": [
          "com.snowplowanalytics"
        ],
        "connection": {
          "http": {
            "uri": "http://iglu-server:8081/api",
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
	assert.JSONEq(string(afterInsert), expectedIgluConfigOne)

	err = ioutil.WriteFile(tmpfn, []byte(igluConfigTwo), 0666)
	assert.Nil(err)

	err = localIglu.addApiKeyToConfig()
	assert.Nil(err)

	afterInsert, err = ioutil.ReadFile(tmpfn)
	assert.Nil(err)

	assert.JSONEq(string(afterInsert), expectedIgluConfigTwo)
}
