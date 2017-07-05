/**
 * Copyright (c) 2014-2017 Snowplow Analytics Ltd.
 * All rights reserved.
 *
 * This program is licensed to you under the Apache License Version 2.0,
 * and you may not use this file except in compliance with the Apache
 * License Version 2.0.
 * You may obtain a copy of the Apache License Version 2.0 at
 * http://www.apache.org/licenses/LICENSE-2.0.
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the Apache License Version 2.0 is distributed
 * on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied.
 *
 * See the Apache License Version 2.0 for the specific language
 * governing permissions and limitations there under.
 */

package main

import (
    "github.com/emicklei/go-restful"
    "io"
    "net/http"
    "log"
    "os/exec"
    "flag"
)

//global variable for script's path
var scriptsPath string

func main() {
    scriptsPathFlag := flag.String("scriptsPath", "", "path for control-plane-api scripts")
    flag.Parse()
    scriptsPath = *scriptsPathFlag

    ws := new(restful.WebService)
    ws.Route(ws.PUT("/restart-services").To(restartSPServices))
    restful.Add(ws)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func restartSPServices(req *restful.Request, resp *restful.Response) {
    cmd := exec.Command("/bin/sh", scriptsPath + "/" +  "restart_SP_services.sh")
    err := cmd.Run()
    if err != nil {
        io.WriteString(resp, "ERR")
    } else { 
    	io.WriteString(resp, "OK")
    }
}
