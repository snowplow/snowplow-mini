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
	"encoding/base64"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/trustelem/zxcvbn"
	"golang.org/x/crypto/bcrypt"
)

func changeCredentials(configPath string, username string, password string) (error, int) {

	hashedPassword, err := checkPassword(password)
	if err != nil {
		return err, 400
	}

	lines, err := fileToLines(configPath)
	if err != nil {
		return err, 500
	}
	fileContent := ""
	hit := false
	for _, line := range lines {
		if hit {
			line = "        " + username + " " + base64Encode(hashedPassword)
			hit = false
		}
		if strings.Contains(line, "basicauth") {
			hit = true
		}
		fileContent += line
		fileContent += "\n"
	}
	return ioutil.WriteFile(configPath, []byte(fileContent), 0644), 500
}

func checkPassword(password string) (string, error) {

	if len(password) < 8 {
		return "", errors.New("weak password: length can not be shorter than 8 characters")
	}

	minEntropyMatch := zxcvbn.PasswordStrength(password, nil)

	// see https://github.com/nbutton23/zxcvbn-go#use
	if minEntropyMatch.Score < 4 {
		return "", errors.New("weak password: strength score is " + strconv.Itoa(minEntropyMatch.Score) + " but must be 4 at least")
	}

	hashedPassword, err := bcryptHash(password)
	if err != nil {
		return "", err
	}

	return hashedPassword, nil
}

func bcryptHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func base64Encode(in string) string {
	return base64.StdEncoding.EncodeToString([]byte(in))
}
