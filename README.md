# learn-cicd-starter (Notely)

[![ci](https://github.com/gainax2k1/learn-cicd-starter/actions/workflows/ci.yml/badge.svg)](https://github.com/gainax2k1/learn-cicd-starter/actions/workflows/ci.yml)



This repo contains the starter code for the "Notely" application for the "Learn CICD" course on [Boot.dev](https://boot.dev).

## Local Development

Make sure you're on Go version 1.22+.

Create a `.env` file in the root of the project with the following contents:

```bash
PORT="8080"
```

Run the server:

```bash
go build -o notely && ./notely
```

*This starts the server in non-database mode.* It will serve a simple webpage at `http://localhost:8080`.

You do *not* need to set up a database or any interactivity on the webpage yet. Instructions for that will come later in the course!

gainax2k1's version of Boot.dev's Notely app