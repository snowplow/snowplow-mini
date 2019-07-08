package main

import (
    "github.com/caddyserver/caddy/caddy/caddymain"
)

func main() {
    // optional: disable telemetry
    caddymain.EnableTelemetry = false
    caddymain.Run()
}
