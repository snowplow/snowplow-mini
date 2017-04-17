#!/bin/sh

/etc/init.d/postgresql start && \
/etc/init.d/elasticsearch start && \
/etc/init.d/kibana4_init start && \
/etc/init.d/snowplow_stream_collector_0.6.0 start && \
/etc/init.d/snowplow_stream_enrich_0.7.0 start && \
/etc/init.d/snowplow_elasticsearch_sink_good_0.5.0 start && \
/etc/init.d/snowplow_elasticsearch_sink_bad_0.5.0 start && \
/etc/init.d/iglu_server_0.2.0 start && \
/etc/init.d/nginx start

while true;do sleep 5;done
