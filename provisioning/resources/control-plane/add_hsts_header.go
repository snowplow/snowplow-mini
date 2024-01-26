/**
 * Copyright (c) 2024-present Snowplow Analytics Ltd. All rights reserved.
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

func addHstsHeader(configPath string) error {
	currentConfig, err := ioutil.ReadFile(configPath)

	if err != nil {
		return err
	}
	toReplacePattern :=
		`
      handle @isHttps {
        import handleProtectedPaths
      }
`
	replaceWithHsts :=
		`
      handle @isHttps {
        import handleProtectedPaths
        header Strict-Transport-Security "max-age=31536000; includeSubDomains"
      }
`
	newCaddyConfig := strings.Replace(string(currentConfig), toReplacePattern, replaceWithHsts, 1)
	return ioutil.WriteFile(configPath, []byte(newCaddyConfig), 0644)
}
