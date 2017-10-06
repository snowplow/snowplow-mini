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
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestChangeDomainName(t *testing.T) {
	assert := assert.New(t)

	caddyConfigHeadBefore :=
		`*:80 {
  tls off
  basicauth username_test password_test {
    /home
    /kibana
    /elasticsearch
    /control-plane
    /_plugin
  }
`
	expectedCaddyConfigHeadAfter :=
		`example.com *:80 {
  tls example@example.com
  basicauth username_test password_test {
    /home
    /kibana
    /elasticsearch
    /control-plane
    /_plugin
  }
`
	dir, err := ioutil.TempDir("", "testDir")
	assert.Nil(err)

	defer os.RemoveAll(dir)

	tmpfn := filepath.Join(dir, "tmpfile")

	err = ioutil.WriteFile(tmpfn, []byte(caddyConfigHeadBefore), 0666)
	assert.Nil(err)

	err = changeDomainName(
		tmpfn,
		"example.com",
	)
	assert.Nil(err)

	caddyConfigAfter, err := ioutil.ReadFile(tmpfn)
	assert.Nil(err)

	assert.True(expectedCaddyConfigHeadAfter == string(caddyConfigAfter))
}
