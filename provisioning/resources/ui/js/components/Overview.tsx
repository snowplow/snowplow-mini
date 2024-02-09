/*
 * Copyright (c) 2016-present Snowplow Analytics Ltd. All rights reserved.
 *
 * This software is made available by Snowplow Analytics, Ltd.,
 * under the terms of the Snowplow Limited Use License Agreement, Version 1.0
 * located at https://docs.snowplow.io/limited-use-license-1.0
 * BY INSTALLING, DOWNLOADING, ACCESSING, USING OR DISTRIBUTING ANY PORTION
 * OF THE SOFTWARE, YOU AGREE TO THE TERMS OF SUCH LICENSE AGREEMENT.
 */

/// <reference path="../../typings/node/node.d.ts" />
/// <reference path="../../typings/react/react.d.ts" />
/// <reference path="../../typings/react/react-dom.d.ts" />
/// <reference path=".././Interfaces.d.ts"/>

import React = require("react");
import ReactDOM = require("react-dom");

export class Overview extends React.Component<{}, {}> {

  public render() {

    var collector: string = location.protocol + '//' + window.location.host;
    var kibana: string = location.protocol + '//' + window.location.host + '/kibana/';
    var elasticsearch: string = location.protocol + '//' + window.location.host + '/elasticsearch';
    var cAdvisor: string = location.protocol + '//' + window.location.host + '/metrics';

    return (
      <div className="tab-content">
        <p>
          Snowplow Mini is, in essence, the Snowplow real time pipeline contained within a single box.
          In place of Kinesis Streams we are using NSQ and instead of distributing all of the applications across Autoscaling Groups they are deployed onto a single instance.
        </p>
        <h3>1. Sending events</h3>
        <p>You can send events into Snowplow Mini automatically from the <a href="#/example-events">Example events</a> page.  Simply go to that page and click the sample event buttons.</p>
        <p>Alternatively, you can setup any of the Snowplow trackers to send data to this endpoint: {collector}</p>
        <h3>2. Viewing the events</h3>
        <p>You can view the events that have been sent to Opensearch in the <a href={kibana}>Opensearch Dashboards</a> </p>
        <p>You can also submit queries directly to the <a href={elasticsearch}>Opensearch endpoint</a>.</p>
        <h3>3. Understanding how Snowplow Mini works</h3>
        <h3>Quicklinks: </h3>
        <ul>
          <li><a href={'https://docs.snowplowanalytics.com/docs/understanding-your-pipeline/what-is-snowplow-mini'}>What is Snowplow Mini?</a></li>
          <li><a href={'https://docs.snowplowanalytics.com/docs/pipeline-components-and-applications/snowplow-mini/usage-guide/'}>Usage guide</a></li>
          <li><a href={'https://docs.snowplowanalytics.com/docs/pipeline-components-and-applications/snowplow-mini/control-plane-api/'}>Control Plane API</a></li>
          <li>Link to <a href={'https://github.com/snowplow/snowplow-mini'}>Snowplow Mini</a> repository</li>
          <li>Collector endpoint <a href={collector}>{collector}</a></li>
          <li>Metrics endpoint <a href={cAdvisor}>{cAdvisor}</a></li>
        </ul>
        <h3>The software stack installed: </h3>
        <ul>
        <li><b>Snowplow Mini 0.19.1</b></li>
          <li>Snowplow Stream Collector NSQ 3.0.0-rc9</li>
          <li>Snowplow Stream Enrich NSQ 3.8.3</li>
          <li>Snowplow Elasticsearch Loader 2.1.0</li>
          <li>Snowplow Iglu Server 0.10.0</li>
          <li>Postgres 15.1</li>
          <li>NSQ v1.2.1</li>
          <li>Opensearch 2.4.0</li>
          <li>Opensearch Dashboards 2.4.0</li>
          <li>cAdvisor 0.43.0</li>
        </ul>
        <h3>Stack topology: </h3>
        <div>
          <img src="/home/assets/img/snowplow-mini-topology.jpg" />
        </div>
      </div>
    );
  }
}
