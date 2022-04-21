# ZAManager

ZAManager is controller of [ZASentinel](https://github.com/ztalab/ZASentinel)

ZAManager includes the following components：

1. Access model, including Resource, Relay, Server, generate certificate.
2. Oauth2 (developing).
2. System (developing).

## Building

```
$ git clone git@github.com:ztalab/ZAManager.git
$ cd ZAManager
$ make
```

You can set GOOS and GOARCH environment variables to allow Go to cross-compile alternative platforms.

The resulting binaries will be in the bin folder:

```
$ tree bin
bin
├── zaca
```

## Configuration reference

**configuration file:**

The configuration file is in the project root directory：`config.yml` ,The file format is standard yaml format, which can be used as a reference.

ca config see [ZACA](https://github.com/ztalab/ZACA)


## Service Installation

Start command：`ZAManager -c config.yaml`，Default listening port 80
