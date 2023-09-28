/**
 * Copyright (c) 2016-present Snowplow Analytics Ltd. All rights reserved.
 *
 * This program is licensed to you under the Snowplow Community License Version 1.0,
 * and you may not use this file except in compliance with the Snowplow Community License Version 1.0.
 * You may obtain a copy of the Snowplow Community License Version 1.0 at https://docs.snowplow.io/community-license-1.0
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
	Repos     []RepoConf `json:"repositories"`
}

type IgluConf struct {
	Schema string   `json:"schema"`
	Data   DataConf `json:"data"`
}
