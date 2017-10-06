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

/// <reference path="../../../typings/node/node.d.ts" />
/// <reference path="../../../typings/react/react.d.ts" />
/// <reference path="../../../typings/react/react-dom.d.ts" />
/// <reference path="../.././Interfaces.d.ts"/>

import React = require('react');
import ReactDOM = require("react-dom");
import axios from 'axios';

export default React.createClass({
  getInitialState () {
    return {
      new_username: '',
      new_password: '',
      disabled: false
    };
  },

  handleChange(evt) {
    if (evt.target.name == 'new_username'){
      this.setState({
        new_username: evt.target.value
      });
    }
    else if (evt.target.name == 'new_password'){
      this.setState({
        new_password: evt.target.value
      });
    }
  },

  sendFormData()  {
    var _this = this

    // there is no need to make 'disabled' false after
    // because connection will be lost after request is sent
    // and page must be loaded again
    _this.setState({
      disabled: true
    });

    var params = new URLSearchParams();
    params.append('new_username', this.state.new_username)
    params.append('new_password', this.state.new_password)

    var _this = this
    axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'
    axios.post('/control-plane/credentials', params, {})
    .then(function (response) {
      // there is no need to this part because status will be
      // 400 in everytime and this will be handled by catch section
    })
    .catch(function (error) {
    });
  },

  handleSubmit(event) {
    event.preventDefault();
    this.sendFormData();
  },

  render() {
    return  (
      <div className="tab-content">
        <h4>Change username and password for basic http authentication: </h4>
        <form action="" onSubmit={this.handleSubmit}>
          <div className="form-group">
            <label htmlFor="new_username">Username: </label>
            <input className="form-control" name="new_username" ref="new_username" required type="text" onChange={this.handleChange} value={this.state.new_username} />
          </div>
          <div className="form-group">
            <label htmlFor="new_password">Password: </label>
            <input className="form-control" name="new_password" ref="new_password" required type="password" onChange={this.handleChange} value={this.state.new_password} />
          </div>
          <div className="form-group">
            <button className="btn btn-primary" type="submit" disabled={this.state.disabled}>Submit</button>
          </div>
        </form>
      </div>
    );
  }
});
