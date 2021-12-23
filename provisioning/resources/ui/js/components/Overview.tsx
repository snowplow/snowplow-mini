/*
 * Copyright (c) 2016 Snowplow Analytics Ltd. All rights reserved.
 *
 * This program is licensed to you under the Apache License Version 2.0,
 * and you may not use this file except in compliance with the Apache License Version 2.0.
 * You may obtain a copy of the Apache License Version 2.0 at http://www.apache.org/licenses/LICENSE-2.0.
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the Apache License Version 2.0 is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the Apache License Version 2.0 for the specific language governing permissions and limitations there under.
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
        <p>You can view the events that have been sent to Elasticsearch in the <a href={kibana}>Kibana Dashboard</a> </p>
        <p>You can also submit queries directly to the <a href={elasticsearch}>Elasticsearch endpoint</a>.</p>
        <h3>3. Understanding how Snowplow Mini works</h3>
        <h3>Quicklinks: </h3>
        <ul>
          <li><a href={'https://docs.snowplowanalytics.com/docs/understanding-your-pipeline/what-is-snowplow-mini'}>What is Snowplow Mini?</a></li>
          <li><a href={'https://docs.snowplowanalytics.com/docs/open-source-components-and-applications/snowplow-mini/snowplow-mini-0-11-0/usage-guide/'}>Usage guide</a></li>
          <li><a href={'https://docs.snowplowanalytics.com/docs/open-source-components-and-applications/snowplow-mini/snowplow-mini-0-11-0/control-plane-api/'}>Control Plane API</a></li>
          <li>Link to <a href={'https://github.com/snowplow/snowplow-mini'}>Snowplow Mini</a> repository</li>
          <li>Collector endpoint <a href={collector}>{collector}</a></li>
          <li>Metrics endpoint <a href={cAdvisor}>{cAdvisor}</a></li>
        </ul>
        <h3>The software stack installed: </h3>
        <ul>
        <li><b>Snowplow Mini 0.13.5</b></li>
          <li>Snowplow Stream Collector NSQ 2.4.5</li>
          <li>Snowplow Stream Enrich NSQ 2.0.5</li>
          <li>Snowplow Elasticsearch Loader 1.0.4</li>
          <li>Snowplow Iglu Server 0.6.1</li>
          <li>Postgres 9.5</li>
          <li>NSQ v1.2.1</li>
          <li>Elasticsearch-OSS 6.8.20</li>
          <li>Kibana-OSS 6.8.20</li>
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
