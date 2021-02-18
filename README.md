# Snowplow-Mini

[![Join the chat at https://gitter.im/snowplow/snowplow-mini](https://badges.gitter.im/snowplow/snowplow-mini.svg)](https://gitter.im/snowplow/snowplow-mini?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Build Status][build-image]][build-wf]
[![Release][release-image]][releases]
[![License][license-image]][license]

An easily-deployable, single instance version of Snowplow that serves three use cases:

1. Gives a Snowplow consumer (e.g. an analyst / data team / marketing team) a way to quickly understand what Snowplow "does" i.e. what you put in at one end and take out of the other
2. Gives developers new to Snowplow an easy way to start with Snowplow and understand how the different pieces fit together
3. Gives people running Snowplow a quick way to debug tracker updates (because they can)

## Features

* [x] Data is tracked and processed in real time
* [x] Added Iglu Server to allow for custom schemas to be uploaded
* [x] Data is validated during processing
  - This is done using both our standard Iglu schemas and any custom ones that you have loaded into the Iglu Server
* [x] Data is loaded into Elasticsearch
  - Can be queried directly or through a Kibana dashboard
  - Good and bad events are in distinct indexes
* [x] Create UI to indicate what is happening with each of the different subsystems (collector, enrich etc.), so as to provide developers a very indepth way of understanding how the different Snowplow subsystems work with one another

## Topology

Snowplow-Mini runs several distinct applications on the same box which are all linked by NSQ topics.  In a production deployment each instance could be an Autoscaling Group and each NSQ topic would be a distinct Kinesis Stream.

* Scala Stream Collector:
  - Starts server listening on `http://< sp mini public ip>/` which events can be sent to.
  - Sends "good" events to the `RawEvents` NSQ topic
  - Sends "bad" events to the `BadEvents` NSQ topic
* Stream Enrich:
  - Reads events in from the `RawEvents` NSQ topic
  - Sends events which passed the enrichment process to the `EnrichedEvents` NSQ topic
  - Sends events which failed the enrichment process to the `BadEvents` NSQ topic
* Elasticsearch Sink Good:
  - Reads events from the `EnrichedEvents` NSQ topic
  - Sends those events to the `good` Elasticsearch index
  - On failure to insert, writes errors to `BadElasticsearchEvents` NSQ topic
* Elasticsearch Sink Bad:
  - Reads events from the `BadEvents` NSQ topic
  - Sends those events to the `bad` Elasticsearch index
  - On failure to insert, writes errors to `BadElasticsearchEvents` NSQ topic

These events can then be viewed in Kibana at `http://< sp mini public ip>/kibana`.

![](https://raw.githubusercontent.com/snowplow/snowplow-mini/master/utils/topology/snowplow-mini-topology.jpg)

## Roadmap

* [ ] Support loading data into Redshift. To give analysts / data teams a good idea to understand what Snowplow "does".

## Documentation

Documentation is available at [our docs website][mini-docs].

## Vagrant

Some advice on how to handle certain errors if you're trying to build this locally with Vagrant.

### `The box 'ubuntu/xenial64' could not be found or could not be accessed in the remote catalog.`

Your Vagrant version is probably outdated. Use Vagrant 2.0.0+.

### `npm install` results in `enoent ENOENT: no such file or directory, open '/package.json'`

This is caused by trying to use NFS. Comment the relevant lines in `Vagrantfile`.

Most likely this will happen on `TASK [sp_mini_5_build_ui : Install npm packages based on package.json.]` but see also: https://discourse.snowplowanalytics.com/t/snowplow-mini-local-vagrant/2930.

## Copyright and license

Snowplow Mini is copyright 2016-2021 Snowplow Analytics Ltd.

Licensed under the **[Apache License, Version 2.0][license]** (the "License");
you may not use this software except in compliance with the License.

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

[mini-docs]: https://docs.snowplowanalytics.com/docs/open-source-components-and-applications/snowplow-mini/

[build-image]: https://github.com/snowplow/snowplow-mini/workflows/build/badge.svg
[build-wf]: https://github.com/snowplow/snowplow-mini/actions?query=workflow%3Apublish

[release-image]: https://img.shields.io/github/v/release/snowplow/snowplow-mini?sort=semver&style=flat
[releases]: https://github.com/snowplow/snowplow-mini/releases

[license-image]: https://img.shields.io/badge/license-Apache--2-blue.svg?style=flat
[license]: https://www.apache.org/licenses/LICENSE-2.0
