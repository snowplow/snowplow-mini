/*
 * Copyright (c) 2016-present Snowplow Analytics Ltd. All rights reserved.
 *
 * This program is licensed to you under the Snowplow Community License Version 1.0,
 * and you may not use this file except in compliance with the Snowplow Community License Version 1.0.
 * You may obtain a copy of the Snowplow Community License Version 1.0 at https://docs.snowplow.io/community-license-1.0
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
