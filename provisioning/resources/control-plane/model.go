/**
 * Copyright (c) 2016-present Snowplow Analytics Ltd. All rights reserved.
 *
 * This software is made available by Snowplow Analytics, Ltd.,
 * under the terms of the Snowplow Limited Use License Agreement, Version 1.1
 * located at https://docs.snowplow.io/limited-use-license-1.1
 * BY INSTALLING, DOWNLOADING, ACCESSING, USING OR DISTRIBUTING ANY PORTION
 * OF THE SOFTWARE, YOU AGREE TO THE TERMS OF SUCH LICENSE AGREEMENT.
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
	Enrich        string `toml:"enrich"`
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
