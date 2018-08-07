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

/// <reference path="../../../typings/node/node.d.ts" />
/// <reference path="../../../typings/react/react.d.ts" />
/// <reference path="../../../typings/react/react-dom.d.ts" />
/// <reference path="../.././Interfaces.d.ts"/>

import React = require('react');
import ReactDOM = require("react-dom");
import AlertContainer from 'react-alert';
import alertOptions from './AlertOptions'
import axios from 'axios';

var alertContainer = new AlertContainer();

export default React.createClass({
  getInitialState () {
    return {
      iglu_repo_name: '',
      iglu_repo_priority: '',
      iglu_repo_uri: '',
      iglu_repo_apikey: '',
      vendor_prefix:'',
      disabled: false
    };
  },

  handleChange(evt) {
    if (evt.target.name == 'iglu_repo_uri'){
      this.setState({
        iglu_repo_uri: evt.target.value
      });
    }
    else if (evt.target.name == 'iglu_repo_apikey'){
      this.setState({
        iglu_repo_apikey: evt.target.value
      });
    }
    else if (evt.target.name == 'vendor_prefix'){
      this.setState({
        vendor_prefix: evt.target.value
      });
    }
    else if (evt.target.name == 'iglu_repo_name'){
      this.setState({
        iglu_repo_name: evt.target.value
      });
    }
    else if (evt.target.name == 'iglu_repo_priority'){
      this.setState({
        iglu_repo_priority: evt.target.value
      });
    }
  },

  sendFormData()  {
    var alertShow = alertContainer.show
    var _this = this

    var igluRepoUri = this.state.iglu_repo_uri
    var igluRepoApikey = this.state.iglu_repo_apikey
    var vendorPrefix = this.state.vendor_prefix
    var igluRepoName = this.state.iglu_repo_name
    var igluRepoPriority = this.state.iglu_repo_priority

    function setInitState() {
      _this.setState({
        iglu_repo_uri: "",
        iglu_repo_apikey: "",
        vendor_prefix:"",
        iglu_repo_name: '',
        iglu_repo_priority: '',
        disabled: false
      });
    }

    _this.setState({
      disabled: true
    });
    var params = new URLSearchParams();
    params.append('uri', _this.state.iglu_repo_uri)
    params.append('apikey', _this.state.iglu_repo_apikey)
    params.append('vendor_prefix', _this.state.vendor_prefix)
    params.append('name', _this.state.iglu_repo_name)
    params.append('priority', _this.state.iglu_repo_priority)

    axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded'
    axios.post('/control-plane/external-iglu', params, {})
    .then(function (response) {
      setInitState()
      alertShow('Uploaded successfully', {
        time: 2000,
        type: 'success'
      });
    })
    .catch(function (error) {
      setInitState()
      alertShow('Error: ' + error.response.data, {
        time: 2000,
        type: 'error'
      });
    });
  },

  handleSubmit(event) {
    var alertShow = alertContainer.show
    alertShow('Please wait...', {
      time: 2000,
      type: 'info'
    });
    event.preventDefault();
    this.sendFormData();
  },

  render() {
    return (
      <div className="tab-content">
        <h4>Add external Iglu repository: </h4>
        <form action="" onSubmit={this.handleSubmit}>
          <div className="form-group">
            <label htmlFor="iglu_repo_name">Name of the Iglu repository: </label>
            <input className="form-control" name="iglu_repo_name" ref="iglu_repo_name" required type="text" onChange={this.handleChange} value={this.state.iglu_repo_name} />
          </div>
          <div className="form-group">
            <label htmlFor="iglu_repo_priority">Priority of the Iglu repository: </label>
            <input className="form-control" name="iglu_repo_priority" ref="iglu_repo_priority" required type="number" onChange={this.handleChange} value={this.state.iglu_repo_priority} />
          </div>
          <div className="form-group">
            <label htmlFor="vendor_prefix">Vendor prefix: </label>
            <input className="form-control" name="vendor_prefix" ref="vendor_prefix" required type="text" onChange={this.handleChange} value={this.state.vendor_prefix} />
          </div>
          <div className="form-group">
            <label htmlFor="iglu_repo_uri">Iglu repository URI: </label>
            <input className="form-control" name="iglu_repo_uri" ref="iglu_repo_uri" required type="text" onChange={this.handleChange} value={this.state.iglu_repo_uri} />
          </div>
          <div className="form-group">
            <label htmlFor="iglu_repo_apikey">Optional Iglu repository api key: </label>
            <input className="form-control" name="iglu_repo_apikey" ref="iglurepo_apikey" type="text" onChange={this.handleChange} value={this.state.iglu_repo_apikey}/>
          </div>
          <div className="form-group">
            <button className="btn btn-primary" type="submit" disabled={this.state.disabled}>Add external Iglu repository</button>
          </div>
        </form>
        <AlertContainer ref={a => alertContainer = a} {...alertOptions} />
      </div>
    );
  }
});
