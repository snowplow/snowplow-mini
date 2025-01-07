/*
 * Copyright (c) 2016-present Snowplow Analytics Ltd. All rights reserved.
 *
 * This software is made available by Snowplow Analytics, Ltd.,
 * under the terms of the Snowplow Limited Use License Agreement, Version 1.1
 * located at https://docs.snowplow.io/limited-use-license-1.1
 * BY INSTALLING, DOWNLOADING, ACCESSING, USING OR DISTRIBUTING ANY PORTION
 * OF THE SOFTWARE, YOU AGREE TO THE TERMS OF SUCH LICENSE AGREEMENT.
 */

/// <reference path="../typings/node/node.d.ts" />
/// <reference path="../typings/react/react.d.ts" />
/// <reference path="../typings/react/react-dom.d.ts" />
/// <reference path="./Interfaces.d.ts"/>

declare var Router;
import React = require("react");
import ReactDOM = require("react-dom");
import { Elasticsearch } from "./components/Elasticsearch";
import { IgluServer } from "./components/IgluServer";
import { JSTracker } from "./components/JSTracker";
import { Overview } from "./components/Overview";
import { ControlPlane } from "./components/ControlPlane";
var Tabs = require('react-simpletabs');

/**
 * Entry point for the Application
 */
export class SnowplowMiniApp extends React.Component<{}, IAppState> {

  /**
   * Constructs the SnowplowMini React Component
   * @param props The starting properties
   */
  constructor(props) {
    super(props);
    this.state = { index: 1 };
  }

  /**
   * When component attached to DOM will set up the
   * link router.
   */
  public componentDidMount() {
    var routes = { 
      "/overview": this.setState.bind(this, { index: 1 }),
      "/example-events": this.setState.bind(this, { index: 2 }),
      "/elasticsearch": this.setState.bind(this, { index: 3 }),
      "/iglu-server": this.setState.bind(this, { index: 4 }),
      "/control-plane": this.setState.bind(this, { index: 5 })  
    };
    var router = Router(routes).configure({
      notfound: function() {
        document.location.href = "#/overview";
      }
    });
    router.init('/overview');
  }

  /**
   * Will render the content in the DOM
   */
  public render() {
    return (
      <div>
        <div className="main-content">
          <Tabs 
            tabActive={this.state.index}
            onBeforeChange={this.handleBefore.bind(this)}>
            <Tabs.Panel key={1} title="Overview">
              <Overview />
            </Tabs.Panel>
            <Tabs.Panel key={2} title="Example events">
              <JSTracker />
            </Tabs.Panel>
            <Tabs.Panel key={3} title="Opensearch">
              <Elasticsearch />
            </Tabs.Panel>
            <Tabs.Panel key={4} title="Iglu Server">
              <IgluServer />
            </Tabs.Panel>
            <Tabs.Panel key={5} title="Control Plane">
              <ControlPlane />
            </Tabs.Panel>
          </Tabs>
        </div>
      </div>
    );
  }

  /**
   * Will set the state to the correct index on click.
   */
  private handleBefore(selectedIndex, $selectedPanel, $selectedTabMenu) {
    if (selectedIndex == 1) {
      document.location.href = "#/overview";
    } else if (selectedIndex == 2) {
      document.location.href = "#/example-events";
    } else if (selectedIndex == 3) {
      document.location.href = "#/elasticsearch";
    } else if (selectedIndex == 4) {
      document.location.href = "#/iglu-server";
    } else if (selectedIndex == 5) {
      document.location.href = "#/control-plane";
    }
  }
}

/**
 * Renders the application component.
 */
function render() {
  ReactDOM.render(
    <SnowplowMiniApp />,
    document.getElementsByClassName('snowplowminiapp')[0]
  );
}

render();
