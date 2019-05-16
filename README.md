Kyber-API
=======

[![license](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![GitHub issues](https://img.shields.io/github/issues/matthewedwards/kyber-api.svg)](https://github.com/matthewedwards/kyber-api/issues)
![Contributions welcome](https://img.shields.io/badge/contributions-welcome-green.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/matthewedwards/kyber-api)](https://goreportcard.com/report/github.com/matthewedwards/kyber-api)


Kyber-API is the API server for the kyber news dashboard.

## Table of Contents
- [Installation](#installation)
    - [Prerequisites](#prerequisites)
    - [Getting the Code](#getting-the-code)
- [Versioning](#versioning)

## Installation

### Prerequisites
Before you get the code you will need to install the following dependencies

- A MongoDB database which can be obtained from the [Mongo Website](https://www.mongodb.com/)

### Getting the Code

First checkout the code using git
```
git clone git@github.com:matthewedwards/kyber-api/platform.git src\github.com\matthewedwards/kyber-api
```

Then cd into the directory and build the project
``` 
make build
``` 

Finally start the API 
``` 
make run
``` 

## Versioning

We use [SemVer](http://semver.org/) for versioning. 