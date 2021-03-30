# VirtusLab Recruitment Task
> Simple command-line interface with server functionality.

## Table of contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)
* [Commands](#commands)
* [Features](#features)

## General info
Project was created to fulfill VirtusLab internship recruitment task requirements. It is a simple command-line interface (CLI) with server functionality. The server serves __only__ files that are passed as an argument to the program. More on that in [Commands](#commands) section.

## Technologies
* Go 1.16.2

## Setup
Make sure that in the directory with your `virtus-lab-http-server.exe` there is a directory `static` with files, that you want to serve. To run the server open command line, go to the location of your `virtus-lab-http-server.exe` and type in `virtus-lab-http-server.exe run --file <file_you_want_to_serve>`.

## Commands

### version
`virtus-lab-http-server.exe version`

Shows current version of project.
```
1.0.0
```

### help
`virtus-lab-http-server.exe help`

Shows help for the project.
```
NAME:
   Simple CLI - CLI project for VirtusLab

USAGE:
   virtus-lab-http-server.exe [global options] command [command options] [arguments...]

VERSION:
   1.0.0

COMMANDS:
   version, v  Check app version
   run, r      Run file server
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version
```

### run --file <file>
`virtus-lab-http-server.exe run --file index.html`

Starts `localhost` server on port `8080`
```
Starting server at port 8080
```

## Usage
After starting a server open any browser and go to http://localhost:8080/.
To stop server from running use `CTRL+C` in the command line.

## Features
* Running local server serving file with a name provided by user.