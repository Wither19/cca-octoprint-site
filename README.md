# CCA Octoprint Website

This is a custom-made frontend made by a compiled React + TypeScript Vite app, and served by Go that talks to Octoprint's API.

## Prerequisites

- Latest LTS version of NodeJS
- Latest release of Go
- Some printers using OctoPrint

## Installation

`make node` will install all frontend dependencies with NPM as a rule (If you want to use some other package manager you can enter the frontend directory and set it up manually).

`make build` first compiles the frontend, then embeds it all into the Go binary. It will then run the binary and is hosted at [localhost:8080](localhost:8080) by default.