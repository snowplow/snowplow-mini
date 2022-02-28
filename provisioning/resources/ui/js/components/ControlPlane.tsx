/*
 * Copyright (c) 2016-2018 Snowplow Analytics Ltd. All rights reserved.
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

export class ControlPlane extends React.Component<{}, {}> {

  public render() {
    return (
      <div className="tab-content">
        <p>The control-plan can be used to interact with the internal systems of Snowplow Mini:</p>
        <h3>Quicklinks:</h3>
        <ul>
          <li><a href='/swagger/'>Control Plane Swagger ui</a> page</li>
          <li><a href='https://docs.snowplowanalytics.com/docs/pipeline-components-and-applications/snowplow-mini/control-plane-api/'>Guide to using the Control Plane API</a></li>
        </ul>
      </div>
    );
  }
}
