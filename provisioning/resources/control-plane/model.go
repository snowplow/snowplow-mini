/**
 * Copyright (c) 2016-present Snowplow Analytics Ltd. All rights reserved.
 *
 * This program is licensed to you under the Snowplow Community License Version 1.0,
 * and you may not use this file except in compliance with the Snowplow Community License Version 1.0.
 * You may obtain a copy of the Snowplow Community License Version 1.0 at https://docs.snowplow.io/community-license-1.0
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
