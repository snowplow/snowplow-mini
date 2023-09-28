/**
 * Copyright (c) 2016-present Snowplow Analytics Ltd. All rights reserved.
 *
 * This program is licensed to you under the Snowplow Community License Version 1.0,
 * and you may not use this file except in compliance with the Snowplow Community License Version 1.0.
 * You may obtain a copy of the Snowplow Community License Version 1.0 at https://docs.snowplow.io/community-license-1.0
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
