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
