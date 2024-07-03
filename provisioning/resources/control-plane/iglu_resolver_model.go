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

type ConnectionConf struct {
	Http map[string]string `json:"http"`
}

type RepoConf struct {
	Name           string         `json:"name"`
	Priority       int            `json:"priority"`
	VendorPrefixes []string       `json:"vendorPrefixes"`
	Conn           ConnectionConf `json:"connection"`
}

type DataConf struct {
	CacheSize int        `json:"cacheSize"`
	CacheTtl  int        `json:"cacheTtl"`
	Repos     []RepoConf `json:"repositories"`
}

type IgluConf struct {
	Schema string   `json:"schema"`
	Data   DataConf `json:"data"`
}
