# Snowplow Mini - UI

## Quickstart Guide

```bash
$guest cd /vagrant/ui
$guest npm install
$guest tsc -p js --outDir dist/ && browserify dist/SnowplowMiniApp.js -o dist/bundle.js && uglifyjs dist/bundle.js > dist/snowplow-mini.js
```

Then open `index.html` in a browser of your choice.
