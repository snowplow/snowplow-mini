# Snowplow Mini - UI

## Quickstart Guide

```bash
$guest cd /vagrant/ui
$guest npm install
$guest tsc -p js --outDir dist/ && browserify dist/SnowplowMiniApp.js -o dist/bundle.js
```

Then open `index.html` in a browser of your choice.
