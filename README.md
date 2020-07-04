# p2jsongo

Simple, incomplete tool for transforming `.properties` files into json. Not suitable for production.

## Installation

p2jsongo available in homebrew and linuxbrew.

```bash
$ brew tap LasEmil/p2jsongo
$ brew install p2jsongo
```

## Usage

Basic usage:

```bash
$ p2jsongo parse [properties file to parse] [destination file] [flags]
```

#### -f, --flat

If this flag is set the resulting json file will have only one level

### Commands and options

```
Usage:
  p2jsongo [command]

Available Commands:
  help        Help about any command
  parse       Parse properties file to json

Flags:
  -f, --flat   flat parse
  -h, --help   help for p2json
```

## License

This software is released under the MIT License, see [LICENSE.](https://github.com/LasEmil/p2jsongo/blob/master/LICENSE)
