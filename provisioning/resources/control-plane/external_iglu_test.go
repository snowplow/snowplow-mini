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
