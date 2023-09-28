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
			cmd := exec.Command("docker-compose", restartCommandArgs...)
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
	cmd := exec.Command("docker-compose", restartCommandArgs...)
	cmd.Dir = "/home/ubuntu/snowplow"
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
