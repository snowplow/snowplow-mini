/**
 * Copyright (c) 2021-present Snowplow Analytics Ltd. All rights reserved.
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
	"strings"
)

func getTelemetryDisable(configPath string) (string, error) {
	lines, err := fileToLines(configPath)
	if err != nil {
		return "", err
	}

	disable := ""
	for _, line := range lines {
		if strings.Contains(line, "disable") {
			split := strings.Split(line, "=")
			disable = strings.TrimSpace(split[1])
			break
		}
	}

	return disable, nil
}

func setTelemetryDisable(configPath string, disable string) error {
	lines, err := fileToLines(configPath)
	if err != nil {
		return err
	}

	fileContent := ""
	for _, line := range lines {
		if strings.Contains(line, "disable") {
			line = "      disable = " + disable
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(configPath, []byte(fileContent), 0644)
}
