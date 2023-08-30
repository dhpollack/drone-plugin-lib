# drone-plugin-lib

[![Build Status](https://img.shields.io/drone/build/dhpollack/drone-plugin-lib?logo=drone&server=https%3A%2F%2Fdrone.thegeeklab.de)](https://drone.thegeeklab.de/dhpollack/drone-plugin-lib)
[![Go Report Card](https://goreportcard.com/badge/github.com/dhpollack/drone-plugin-lib)](https://goreportcard.com/report/github.com/dhpollack/drone-plugin-lib)
[![GitHub contributors](https://img.shields.io/github/contributors/dhpollack/drone-plugin-lib)](https://github.com/dhpollack/drone-plugin-lib/graphs/contributors)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/dhpollack/drone-plugin-lib)
[![License: Apache-2.0](https://img.shields.io/github/license/dhpollack/drone-plugin-lib)](https://github.com/dhpollack/drone-plugin-lib/blob/main/LICENSE)

> **DISCONTINUED:** As I don't use Drone CI anymore, this project is unmaintained. If you are interested in a free and open source CI system check out [Woodpecker CI](https://woodpecker-ci.org/).

Helper library to reduce the boilerplate code for writing Drone CI plugins.

## Usage

### Download the package

```Shell
go get -d github.com/dhpollack/drone-plugin-lib/v2/errors
go get -d github.com/dhpollack/drone-plugin-lib/v2/urfave
go get -d github.com/dhpollack/drone-plugin-lib/v2/drone
```

### Import the package

```Go
import "github.com/dhpollack/drone-plugin-lib/v2/errors"
import "github.com/dhpollack/drone-plugin-lib/v2/urfave"
import "github.com/dhpollack/drone-plugin-lib/v2/drone"
```

## Contributors

Special thanks to all [contributors](https://github.com/dhpollack/drone-plugin-lib/graphs/contributors). If you would like to contribute, please see the [instructions](https://github.com/dhpollack/drone-plugin-lib/blob/main/CONTRIBUTING.md).

## License

This project is licensed under the Apache-2.0 License - see the [LICENSE](https://github.com/dhpollack/drone-plugin-lib/blob/main/LICENSE) file for details.
