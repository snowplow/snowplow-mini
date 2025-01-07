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
