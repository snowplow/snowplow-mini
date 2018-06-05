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

var kibana_guide = "https://github.com/snowplow/snowplow-mini/wiki/Quickstart-guide#view-in-kibana";
var head_plugin_repo = "https://github.com/mobz/elasticsearch-head";
var elasticsearch_dsl = "https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl.html";

export class Elasticsearch extends React.Component<{}, {}> {

  public render() {
    return (
      <div className="tab-content">
        <p>
          The local Elasticsearch cluster contains your enriched and failed events in two distinct indices.  
          For ease of management Snowplow Mini comes pre-installed with both Kibana and the Head Plugin.
        </p>
        <p>
          <b>Kibana</b> can be used to view, query and discover the data sent into Snowplow Mini.  You can also build visualizations and dashboards from the information available.
          <br></br>For help with setting up Kibana please consult <a href={kibana_guide}>the guide</a>.
        </p>
        <p>
          <b>The Head Plugin</b> for Elasticsearch allows you to view the state of your cluster and to query the data directly through the <a href={elasticsearch_dsl}>Elasticsearch Query DSL</a>.
          <br></br>For more information please consult the <a href={head_plugin_repo}>elasticsearch-head</a> repository.
        </p>
        <h3>Quicklinks:</h3>
        <ul>
          <li>Link to <a href={location.protocol +'//' + window.location.host + '/kibana/'}>Kibana</a></li>
        </ul>
      </div>
    );
  }
}
