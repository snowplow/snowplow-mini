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

func restartSPService(service string) (error, int) {
	var initMap = map[string]string{
		"collector":     config.Inits.Collector,
		"enrich":        config.Inits.Enrich,
		"esLoaderGood":  config.Inits.EsLoaderGood,
		"esLoaderBad":   config.Inits.EsLoaderBad,
		"iglu":          config.Inits.Iglu,
		"kibana":        config.Inits.Kibana,
		"elasticsearch": config.Inits.Elasticsearch,
	}
	if service == "caddy" {
		cmd := exec.Command("/bin/systemctl", "reload", "caddy")
		err := cmd.Run()
		if err != nil {
			return err, 500
		}
		return nil, 200
	} else {
		if serviceName, ok := initMap[service]; ok {
			restartCommandArgs := []string{"restart", serviceName}
			cmd := exec.Command("/usr/local/bin/docker-compose", restartCommandArgs...)
			cmd.Dir = "/home/ubuntu/snowplow"
			err := cmd.Run()
			if err != nil {
				return err, 500
			}
			return nil, 200
		}
		return errors.New("Recognized service names: collector, enrich, esLoaderGood, esLoaderBad, iglu, kibana, elasticsearch. Received: " + service), 400
	}
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
