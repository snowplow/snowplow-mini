FROM java:7
MAINTAINER Snowplow Support <support@snowplowanalytics.com>

EXPOSE 8080
EXPOSE 5601
EXPOSE 9200

ADD resources/kibana/kibana4_init /etc/init.d/kibana4_init
ADD resources/configs /home/ubuntu/snowplow/configs
ADD resources/elasticsearch /home/ubuntu/snowplow/elasticsearch
ADD scripts /home/ubuntu/snowplow/scripts

RUN /home/ubuntu/snowplow/scripts/1_setup_docker.sh
RUN rm -rf /home/ubuntu/snowplow/staging

CMD /home/ubuntu/snowplow/scripts/2_run_docker.sh
