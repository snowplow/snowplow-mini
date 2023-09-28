/**
 * Copyright (c) 2016-present Snowplow Analytics Ltd. All rights reserved.
 *
 * This program is licensed to you under the Snowplow Community License Version 1.0,
 * and you may not use this file except in compliance with the Snowplow Community License Version 1.0.
 * You may obtain a copy of the Snowplow Community License Version 1.0 at https://docs.snowplow.io/community-license-1.0
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
