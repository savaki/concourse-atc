# concourse-atc

savaki/concourse-atc is a containerized version of github.com/concourse/atc

## Quick Start - docker cli

The quickest way to launch concourse atc is via docker containers.  Run the following command start a postgres db along with a connected atc server.

```
docker run -d --name atcdb -e POSTGRES_USER=atc -e POSTGRES_PASSWORD=password -p 5432:5432 postgres
docker run -d --name atc -e ATC_DEV=true --link atcdb:db -p 8080:8080 savaki/concourse-atc
```

## Quick Start - docker-compose

If you're using docker compose, the quickest way to get started is clone this repository and run:

```
docker-compose up
```

## Environment Variables

The savaki/concourse-atc container can be configured via environment parameters passed in.  The following lists the supported environment variables:

| Env | CLI Equivalent | Description |
| :--- | :--- | :--- |
| ATC_DEV | -dev | dev mode; lax security |
| ATC_CALLBACK_URL | -callbacksURL | URL used for callbacks to reach the ATCD (excluding basic auth) |
| ATC_CHECK_INTERVAL | -checkInterval | interval on which to poll for new versions of resources |
| ATC_CLI_DOWNLOADS_DIR | -cliDownloadsDir | directory containing CLI binaries to serve |
| ATC_USERNAME | -httpUsername | basic auth username for the server |
| ATC_PASSWORD | -httpPassword | basic auth password for the server |
| ATC_SQL_DATASOURCE | -sqlDataSource | database/sql data source configuration string |
| ATC_SQL_DRIVER | -sqlDriver | database/sql driver name |
| ATC_PUBLIC | -public | path to directory containing public resources (javascript, css, etc.) | 
| ATC_TEMPLATES | -templates | path to directory containing the html templates |
| PORT | -webListenPort | port for the web server to listen on |


