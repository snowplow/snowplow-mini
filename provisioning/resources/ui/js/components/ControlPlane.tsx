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
