# golang example web app

## group: `/v1/`

`/config`: shows content of config.yml file

`/env`: shows environment variables available for pod

`/headers`: shows headers set for request

`/hostname`: show hostname

`/health`: show app health

## global

`:9781/metrics`: show app's prometheus metrics

`/stop`: gracefully stop the app

## ports

`8080`: app port

`9781`: metrics middleware port
