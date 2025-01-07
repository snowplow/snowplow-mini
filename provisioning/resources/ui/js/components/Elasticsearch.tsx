/*
 * Copyright (c) 2016-present Snowplow Analytics Ltd. All rights reserved.
 *
 * This software is made available by Snowplow Analytics, Ltd.,
 * under the terms of the Snowplow Limited Use License Agreement, Version 1.1
 * located at https://docs.snowplow.io/limited-use-license-1.1
 * BY INSTALLING, DOWNLOADING, ACCESSING, USING OR DISTRIBUTING ANY PORTION
 * OF THE SOFTWARE, YOU AGREE TO THE TERMS OF SUCH LICENSE AGREEMENT.
 */


/// <reference path="../../typings/node/node.d.ts" />
/// <reference path="../../typings/react/react.d.ts" />
/// <reference path="../../typings/react/react-dom.d.ts" />
/// <reference path=".././Interfaces.d.ts"/>

import React = require("react");
import ReactDOM = require("react-dom");

export class Elasticsearch extends React.Component<{}, {}> {

  public render() {
    return (
      <div className="tab-content">
        <p>
          The local Opensearch cluster contains your enriched and failed events in two distinct indices.
          For ease of management Snowplow Mini comes pre-installed with Opensearch Dashboards.
        </p>
        <p>
          <b>Opensearch Dashboards</b> can be used to view, query and discover the data sent into Snowplow Mini.
          You can also build visualizations and dashboards from the information available.
        </p>
        <h3>Quicklinks:</h3>
        <ul>
          <li>Link to <a href={location.protocol +'//' + window.location.host + '/kibana/'}>Opensearch Dashboards</a></li>
        </ul>
      </div>
    );
  }
}
