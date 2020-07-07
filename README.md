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
$ p2jsongo parse [properties file or directory to parse] [destination filename (optional)] [flags]
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

```
parse description:
Parse java's properties file format to json. You can parse single file or entire directory. When parsing single file you have an option to add second argument: output filename, when parsing entire directory this argument is skipped and all files will be named like the source files

Usage:
  p2jsongo parse [properties file or directory to parse] [destination filename (optional)] [flags]

Flags:
  -h, --help   help for parse

Global Flags:
  -f, --flat   flat parse
```

## License

This software is released under the MIT License, see [LICENSE.](https://github.com/LasEmil/p2jsongo/blob/master/LICENSE)
