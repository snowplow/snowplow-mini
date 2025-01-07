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

export class IgluServer extends React.Component<{}, {}> {

  public render() {
    return (
      <div className="tab-content">
        <p>
          The local Iglu Server contains all of your own custom JsonSchemas that have been uploaded to the server.
        </p>
        <h3>Quicklinks:</h3>
        <ul>
          <li>Link to <a href='/iglu-server/static/swagger-ui/index.html'>Iglu Server management</a> page</li>
        </ul>
      </div>
    );
  }
}
