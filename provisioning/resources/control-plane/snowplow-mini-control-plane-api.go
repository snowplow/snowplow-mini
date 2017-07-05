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
    ws.Route(ws.PUT("/restartspservices").To(restartSPServices))
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
