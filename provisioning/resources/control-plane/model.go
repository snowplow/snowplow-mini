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

type directories struct {
	Enrichments string
	Config      string
}

type configFileNames struct {
	Caddy        string
	IgluResolver string `toml:"iglu_resolver"`
	IgluServer   string `toml:"iglu_server"`
	Collector    string
}

type initScripts struct {
	Collector     string `toml:"stream_collector"`
	Enrich        string `toml:"stream_enrich"`
	EsLoaderGood  string `toml:"es_loader_good"`
	EsLoaderBad   string `toml:"es_loader_bad"`
	Iglu          string
	Kibana        string
	Elasticsearch string
}

type psqlInfos struct {
	User     string
	Password string
	Database string
	Addr     string `toml:"address"`
}

type ControlPlaneConfig struct {
	VersionFilePath string          `toml:"version_file_path"`
	EC2MetaServivce string          `toml:"EC2_meta_service_url"`
	Dirs            directories     `toml:"directories"`
	ConfigNames     configFileNames `toml:"config_file_names"`
	Inits           initScripts     `toml:"init_scripts"`
	Psql            psqlInfos       `toml:"PSQL"`
}
