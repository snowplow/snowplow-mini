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
	"errors"
	"os/exec"
)

func restartService(service string) error {
	var initMap = map[string]string{
		"streamCollector": config.Inits.StreamCollector,
		"streamEnrich":    config.Inits.StreamEnrich,
		"esLoaderGood":    config.Inits.EsLoaderGood,
		"esLoaderBad":     config.Inits.EsLoaderBad,
		"iglu":            config.Inits.Iglu,
		"caddy":           config.Inits.Caddy,
	}

	if val, ok := initMap[service]; ok {
		if service == "caddy" {
			restartCommand := []string{"service", val, "restart"}
			cmd := exec.Command("/bin/bash", restartCommand...)
			err := cmd.Run()
			if err != nil {
				return err
			}
			return nil
		} else {
			restartCommandArgs := []string{"-f", "/home/ubuntu/snowplow/docker-compose.yml", 
											"restart", val}
			cmd := exec.Command("/usr/local/bin/docker-compose", restartCommandArgs...)
			cmd.Dir = "/home/ubuntu/snowplow"
			err := cmd.Run()
			if err != nil {
				return err
			}
			return nil
		}
	}
	return errors.New("unrecognized service: " + service)
}

func restartSPServices() error {
	restartCommandArgs := []string{"-f", "/home/ubuntu/snowplow/docker-compose.yml", "restart"}
	cmd := exec.Command("/usr/local/bin/docker-compose", restartCommandArgs...)
	cmd.Dir = "/home/ubuntu/snowplow"
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
