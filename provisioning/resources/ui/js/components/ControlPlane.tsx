/*
 * Copyright (c) 2016-2017 Snowplow Analytics Ltd. All rights reserved.
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
import RestartServicesSection from "./ControlPlaneComponents/RestartServices";
import UploadEnrichmentsForm from "./ControlPlaneComponents/UploadEnrichments";
import AddExternalIgluServerForm from "./ControlPlaneComponents/AddExternalIgluServer";
import AddLocalIgluApikeyForm from "./ControlPlaneComponents/AddLocalIgluApikey";
import ChangeUsernamePasswordForm from "./ControlPlaneComponents/ChangeUsernamePassword";
import AddDomainNameForm from "./ControlPlaneComponents/AddDomainName";

export class ControlPlane extends React.Component<{}, {}> {

  public render() {
    return (
      <div className="tab-content">
        <p>The buttons below can be used to interact with the internal systems of Snowplow Mini:</p>
        <RestartServicesSection />
        <UploadEnrichmentsForm />
        <AddExternalIgluServerForm />
        <AddLocalIgluApikeyForm />
        <ChangeUsernamePasswordForm />
        <AddDomainNameForm />
      </div>
    );
  }
}
