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
