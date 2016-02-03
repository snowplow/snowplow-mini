# Snowplow-Mini

An easily-deployable, single instance version of Snowplow that serves three use cases:

1. Gives a Snowplow consumer (e.g. an analyst / data team / marketing team) a way to quickly understand what Snowplow "does" i.e. what you put it at one end and take out of the other
2. Gives developers new to Snowplow an easy way to start with Snowplow and understand how the different pieces fit together
3. Gives people running Snowplow a quick way to debug tracker updates (because they can )

## v1

The initial version of Snowplow-mini has only a limited subset of functionality:

1. Data can be tracked in real time and loaded into Elasticsearch, where it can be queried (either directly or via Kibana)
2. Loading data into Redshift is not supported. (So this does not yet give analysts / data teams a good idea to understand what Snowplow "does")
3. No UI is provided to indicate what is happening with each of the different subsystems (collector, enrich etc.), so this does not provide developers a very good way of understanding how the different Snowplow subsystems work with one another
4. No validation is perfomed on the data, so this is not especially useful for Snowplow users who want to debug instrumentations of e.g. new trackers prior to pushing them live on Snowplow proper

## Documentation

1. [Quick start guide] [get-started-guide]

[get-started-guide]: https://github.com/snowplow/snowplow-mini/wiki/Quickstart-guide
