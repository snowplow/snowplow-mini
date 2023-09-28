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

export class JSTracker extends React.Component<{}, {}> {

  public render() {
    return (
      <div className="tab-content">
        <p>Press the buttons below to trigger individual tracking events:<br></br><br></br>
          <button type="button" onClick={this.trackPageView.bind(this) }>Track this page view</button>
          <br></br>
          <button type="button" onClick={this.playMix.bind(this)}>Play a mix</button>
          <br></br>
          <button type="button" onClick={this.addProduct.bind(this)}>Add a product</button>
          <br></br>
          <button type="button" onClick={this.addEcommerceTransaction.bind(this)}>Add an ecommerce transaction</button>
        </p>
        <p>All of these events are sent using the Snowplow Javascript Tracker and will be sent directly to the Snowplow Mini collector.</p>
      </div>
    );
  }

  // --- Example Events

  private trackPageView(): void {
    alert("Tracking this page view")
    window['snowplow']('trackPageView', 'Example events');
  }

  private playMix(): void {
    alert("Playing a mix");
    window['snowplow']('trackStructEvent', 'Mixes', 'Play', 'MRC/fabric-0503-mix', '', '0.0');
  }

  private addProduct(): void {
    alert("Adding a product to basket");
    window['snowplow']('trackStructEvent', 'Checkout', 'Add', 'ASO01043', 'blue:xxl', '2.0');
  }

  private addEcommerceTransaction(): void {
    alert('Adding an ecommerce transaction');
    var orderId = 'order-123';
    window['snowplow']('addTrans',orderId,'','8000','','','','','','JPY');
    window['snowplow']('addItem',orderId,'1001','Blue t-shirt','','2000','2','JPY');
    window['snowplow']('addItem',orderId,'1002','Red shoes','','4000','1','JPY');
    window['snowplow']('trackTrans');
  }
}
