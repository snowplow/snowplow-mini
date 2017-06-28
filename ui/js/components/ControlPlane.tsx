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
import axios from 'axios';

export class ControlPlane extends React.Component<{}, {}> {

  public render() {
    return (
      <div className="tab-content">
        <p>Press the buttons for controlling the Snowplow-Mini without ssh:<br></br><br></br>
          <button type="button" onClick={this.restartCache.bind(this) }>Reset Cache</button>
          <br></br>
        </p>
      </div>
    );
  }

  private restartCache(): void {
    alert("Restarting cache...")

    axios.put('/controlplane/restartspservices', {}, {})
    .then(function (response) {
        console.log(response.data)
    })
    .catch(function (error) {
        console.log(error);
     });
  }
}


