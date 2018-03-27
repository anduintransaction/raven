## Raven - A simple fake mail server

### Features

 - SMTP Server
 - Mailgun-compatible API (for send message only)
 - Dashboard to view captured emails
 
### Install

Download latest release from [release page](https://github.com/anduintransaction/raven/releases)

### How to run

 - Start a Postgres server
 - Edit config file (see `sample.conf.yml`)
 - Run `./raven run config.yml --ui-data=./frontend`
 - Mailgun API running at http://localhost:8001
 - SMTP Server running at http://localhost:8025
 - Dashboard running at http://localhost:8000
 
 ### Docker
 
 - `docker run -e POSTGRES_HOST=local -e POSTGRES_USER=your-username -e POSTGRES_PASSWORD=your-password -e POSTGRES_DB=your-db anduin/raven:1.0.0`
