# pksetdev
This application is going to install necessary tools and set up the necessary related environment variables.

pksetdev is a CLI tool to manage device configuration using YAML files. It helps automate device setup and ensures consistency across environments.

## Badges
![golang version](https://img.shields.io/github/go-mod/go-version/pezhmankasraee/pksetdev)
[![GitHub release](https://img.shields.io/github/v/release/pezhmankasraee/pksetdev)](https://github.com/pezhmankasraee/pksetdev/releases)
[![Go Reference](https://pkg.go.dev/badge/github.com/pezhmankasraee/pksetdev.svg)](https://pkg.go.dev/github.com/pezhmankasraee/pksetdev)
[![License](https://img.shields.io/github/license/pezhmankasraee/pksetdev)](https://github.com/pezhmankasraee/pksetdev/blob/master/LICENSE)

## How to build the application
To install the applicatio, you only need to execute `makefile`.

```bash
$ make
```

## Flags

Certainly! Here’s a clear and user-friendly way to elaborate and present the help section for **pksetdev** in your `README.md`, making it informative and easy to follow:

---

## Usage

```sh
pksetdev [OPTIONS]
```

### Options

| Option                            | Description                       |
|-----------------------------------|-----------------------------------|
| `-p`, `--path PATH`               | Path to the YAML config file. This file should define device settings and parameters. `Example: ./config/config.yaml` configuration file.    |
| `-h`, `--help`                    | Display help information about the command and its options. |
| `-v`, `--Version`                 | Show the current version of the application. |

### Examples

Run **pksetdev** with a specific configuration file:

```sh
./pksetdev --path ./path/to/config.yaml
```

Or using the short option:

```sh
./pksetdev -p ./path/to/config.yaml
```

### Additional Information

For more details, documentation, and advanced usage, please visit the [project repository](https://github.com/pezhmankasraee/pksetdev).

---