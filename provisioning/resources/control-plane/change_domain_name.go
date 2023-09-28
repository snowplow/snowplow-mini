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
