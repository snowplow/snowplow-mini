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
)

func changeDomainName(configPath string, domainName string) error {
	lines, err := fileToLines(configPath)
	if err != nil {
		return err
	}

	// write domain name to first line of Caddy config
	lines[0] = domainName + " *:80 {"
	// placeholder email address, not important
	lines[1] = "  tls example@example.com"

	fileContent := ""
	for _, line := range lines {
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(configPath, []byte(fileContent), 0644)
}
