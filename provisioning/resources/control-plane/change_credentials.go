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
	"io/ioutil"
	"strings"
)

func changeCredentials(configPath string, username string, password string) error {
	lines, err := fileToLines(configPath)
	if err != nil {
		return err
	}

	fileContent := ""
	for _, line := range lines {
		if strings.Contains(line, "basicauth") {
			line = "  basicauth " + username + " " + password + " {"
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(configPath, []byte(fileContent), 0644)
}
