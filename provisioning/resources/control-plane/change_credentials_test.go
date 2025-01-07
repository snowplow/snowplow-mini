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
	"encoding/base64"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestChangeCredentials(t *testing.T) {
	assert := assert.New(t)

	caddyConfigHeadBefore :=
		`basicauth @protectedPaths  {
        USERNAME_PLACEHOLDER JDJhJDA0JFRYSENkLi4vamh0cm1UcHhKWTZEaGVEWm1OMzk4SVZ0ZTVONVVLUzQ5Q3MvYjE0eUF4bEJL
    }
`

	dir, err := ioutil.TempDir("", "testDir")
	assert.Nil(err)

	defer os.RemoveAll(dir)

	tmpfn := filepath.Join(dir, "tmpfile")

	err = ioutil.WriteFile(tmpfn, []byte(caddyConfigHeadBefore), 0666)
	assert.Nil(err)

	testPassword := "5uwk1A,9kdj1!kdkA."

	err, _ = changeCredentials(
		tmpfn,
		"username_test",
		testPassword,
	)
	assert.Nil(err)

	caddyConfigLines, err := fileToLines(tmpfn)
	assert.Nil(err)

	testUsernameExist := false
	for _, line := range caddyConfigLines {
		if strings.Contains(line, "username_test") {
			testUsernameExist = true
			base64EncodedHashedPassword := strings.Fields(line)[1]
			hashedPassword, _ := base64.StdEncoding.DecodeString(base64EncodedHashedPassword)
			assert.Nil(bcrypt.CompareHashAndPassword(hashedPassword, []byte(testPassword)))
		}
	}
	assert.True(testUsernameExist)
}
