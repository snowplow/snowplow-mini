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
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddExternalIgluServer(t *testing.T) {
	assert := assert.New(t)

	igluConfigBefore :=
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
}
`

	expectedIgluConfigAfter :=
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
      },
      {
        "name": "Iglu Server",
        "priority": 0,
        "vendorPrefixes": [
          "vendor"
        ],
        "connection": {
          "http": {
            "uri": "iglu_uri",
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
	err = ioutil.WriteFile(tmpfn, []byte(igluConfigBefore), 0666)
	assert.Nil(err)

	externalIgluServer := ExternalIgluServer{
		ConfigPath: tmpfn,
		IgluInfo: IgluInfo{
			Name:         "Iglu Server",
			Priority:     0,
			VendorPrefix: "vendor",
			Uri:          "iglu_uri",
			Apikey:       "iglu_apikey",
		},
	}

	err = externalIgluServer.addExternalIgluServer()
	assert.Nil(err)

	afterInsert, err := ioutil.ReadFile(tmpfn)
	assert.Nil(err)

	assert.True(string(afterInsert) == expectedIgluConfigAfter)
}

func TestAddExternalIgluServerWithoutApiKey(t *testing.T) {
	assert := assert.New(t)

	igluConfigBefore :=
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
}
`

	expectedIgluConfigAfter :=
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
      },
      {
        "name": "Iglu Server",
        "priority": 0,
        "vendorPrefixes": [
          "vendor"
        ],
        "connection": {
          "http": {
            "uri": "iglu_uri"
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
	err = ioutil.WriteFile(tmpfn, []byte(igluConfigBefore), 0666)
	assert.Nil(err)

	externalIgluServer := ExternalIgluServer{
		ConfigPath: tmpfn,
		IgluInfo: IgluInfo{
			Name:         "Iglu Server",
			Priority:     0,
			VendorPrefix: "vendor",
			Uri:          "iglu_uri",
			Apikey:       "",
		},
	}

	err = externalIgluServer.addExternalIgluServer()
	assert.Nil(err)

	afterInsert, err := ioutil.ReadFile(tmpfn)
	assert.Nil(err)

	assert.True(string(afterInsert) == expectedIgluConfigAfter)
}
