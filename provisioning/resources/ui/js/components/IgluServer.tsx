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
