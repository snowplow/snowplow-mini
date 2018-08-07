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
