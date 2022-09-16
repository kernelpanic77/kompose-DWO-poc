# Compose File Support for Devworkspace Operator POC

[![Build Status](https://travis-ci.org/joemccann/dillinger.svg?branch=master)](https://travis-ci.org/joemccann/dillinger)

This repository is a proof of concept for integrating the compose file support in the devworkspace operator repostiory by resuing the pkgs of kubernetes kompose.

- Currently, the POC only works with compose files which contain deployments and services.
- kompose doesn't provide support for compose files with the build command.
- kubeconfig file needs to be set specifically for Openshift, unlike minikube.

## Work in progress

- Addition of unit testing for testing different compose files.
- Allowing testing on different clusters.
- Addition of support for different kubernetes components like volumes and RBACs for advanced compose file

## Installation

This repo is written in pure golang, so make sure you have golang setup (go 1.16+) on your system before setting up the development environment.

Install the dependencies and devDependencies and start the server.

```sh
go build main.go
```

To run...

```sh
./main
```
