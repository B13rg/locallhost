# locallhost

A simple http server that echoes client http request information.
Also provides links to localhost (`127.0.0.1`, `[::1]`) for easy navigation to locally hosted projects.

<details>
<summary><b>Table of Contents</b></summary>
<p>

- [locallhost](#locallhost)
  - [Key Features](#key-features)
  - [Technical Overview](#technical-overview)
  - [Installation](#installation)
  - [Getting Started](#getting-started)
    - [Configuration](#configuration)
    - [Deployment](#deployment)
  - [Documentation and Additional Resources](#documentation-and-additional-resources)
    - [Links](#links)
  - [Development](#development)
    - [Contributing](#contributing)
  - [License](#license)

</p>
</details>

## Key Features

* Show request IP address and port
* `/ip` endpoint provides raw request IP address
* Show entire HTTP request header
* `/json` endpoint provides request info formatted as JSON
* ipv4+ipv6 localhost links
  * Quick links for port 80
  * Customizable port number

## Technical Overview

Uses Golang's `net/http` and `` packages to display request info.

* 

## Installation

```
go install https://github.com/B13rg/locallhost.git@main
```

## Getting Started

### Configuration

```sh
❯ ./locallhost -h                                                                                                                         ─╯
Start a server on a configured port that returns info about http requests.

Usage:
  locallhost [flags]
  locallhost [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  version     Get version

Flags:
      --color               enable colorized output (default true)
      --debug               log additional information about what the tool is doing. Overrides --loglevel
  -h, --help                help for locallhost
  -L, --loglevel string     set zerolog log level (default "info")
  -p, --port int            Set http port for server to listen on (default 8080)
      --profiledir string   directory to write pprof profile data to

Use "locallhost [command] --help" for more information about a command.
```

### Deployment

#### Natively

```
go build ./locallhost
./locallhost
```

#### Docker

```
docker-compose -f ./docker/compose.yml build
docker-compose -f ./docker/compose.yml up -d
```

## Documentation and Additional Resources

See [./docs]() folder.
Most if not all of the documentation is generated from the golang code.

### Links

* Inspiration:
  * https://locallhost.com

## Development

Uses Golang `1.24.2`.

### Contributing

Open an issue or PR.

## License

[MIT](LICENSE)
