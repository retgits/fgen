---
title: flogo
weight: 5020
pre: "<i class=\"fa fa-terminal\" aria-hidden=\"true\"></i> "
---

The **flogo** CLI tool gives you the ability to build flows and microservices. With this tool you can, among other things, create your applications, build applications and install new extensions. It is also great to use with Continuous Integration and Continuous Deployment tools like Jenkins and Travis-CI. Below is a complete list of all all commands supported, including samples on how to use them.

{{% notice info %}}
Please make sure that you have installed the **flogo** tools as described in [Getting Started > Flogo CLI](../../getting-started/getting-started-cli/)
{{% /notice %}}

## Top-level commands
The Flogo cli has six top-level commands, which you can see by running `flogo` without any additional parameters (or with `-h`, `--help`) Available Commands:
  app         Lists all app actions available to Project Flogo
  contrib     Lists all contrib actions available to Project Flogo
  device      Lists all device actions available to Project Flogo
  help        Help about any command
  plugin      Plugin management for the Project Flogo CLI
  version     Display the current version of the CLI

_The `help` and `version` commands do not have subcommands and display help about the command and the version of the cli respectively_

## App
The app command supports the app capabilities of the Project Flogo CLI.
The commands available are:

build.................... Build a Flogo application
create................... Create a Flogo application
ensure................... Ensure gets a project into a complete, reproducible, and likely compilable state
install.................. Install an app contribution
list..................... List installed contributions
prepare.................. Prepare the flogo application
uninstall................ Uninstall an app contribution

### build
Build a Flogo application

Usage:
  flogo app build [flags]

Flags:
      --docker string   build docker
      --e               embed config
      --gen             only generation
  -h, --help            help for build
      --nogen           no generation
      --o               optimize build
      --shim string     trigger shim

### create
Create a Flogo application

Usage:
  flogo app create [flags]

Flags:
      --file string     The flogo.json to create project from
      --flv string       The flogo dependency constraints as comma separated value (for example github.com/TIBCOSoftware/flogo-lib@0.0.0,github.com/TIBCOSoftware/flogo-contrib@0.0.0)
  -h, --help            help for create
      --name string     The name of the app (required)
      --vendor string   Copy sources from an existing vendor directory

### ensure
Ensure gets a project into a complete, reproducible, and likely compilable state

Usage:
  flogo app ensure [flags]

Flags:
      --add string    add new dependencies, or populate Gopkg.toml with constraints for existing dependencies (default: false)
  -h, --help          help for ensure
      --no-vendor      update Gopkg.lock (if needed), but do not update vendor/ (default: false)
      --update        update the named dependencies (or all, if none are named) in Gopkg.lock to the latest allowed by Gopkg.toml (default: false)
      --vendor-only   populate vendor/ from Gopkg.lock without updating it first (default: false)
      --verbose       enable verbose logging (default: false)

### install
Install an app contribution

Usage:
  flogo app install [flags]

Flags:
  -h, --help             help for install
  -n, --name string      The name of the contribution (required)
  -p, --palette          Install palette file
  -v, --version string   Specify the version of the contribution (optional)

### list
List installed contributions

Usage:
  flogo app list [flags]

Flags:
  -h, --help          help for list
  -j, --json          Generate output as json
  -t, --type string   The type of contribution you want to list ("actions"|"triggers"|"activities") (required)

### prepare
Prepare the flogo application

Usage:
  flogo app prepare [flags]

Flags:
  -e, --embed      Embed configuration into the application
  -h, --help       help for prepare
  -o, --optimize   Optimize the preparation

### uninstall
Uninstall an app contribution

Usage:
  flogo app uninstall [flags]

Flags:
  -h, --help          help for uninstall
  -n, --name string   The name of the contribution (required)

## Contrib
The contrib command supports the contribution capabilities of the Project Flogo CLI.
The commands available are:

create................... Generate a new contribution for Project Flogo
search................... Search for Project Flogo contributions

### create
Generate a new contribution for Project Flogo

Usage:
  flogo contrib create [flags]

Flags:
  -h, --help          help for create
  -n, --name string   The name you want to give your contribution (required)
  -t, --type string   The type of contribution you want to generate ("action"|"activity"|"flowmodel"|"trigger") (required)

### search
Search for Project Flogo contributions

Usage:
  flogo contrib search [flags]

Flags:
  -h, --help            help for search
  -s, --string string   The search string you want to use (optional)
  -t, --type string     The type you're looking for ("all"|"activity"|"trigger") (required)

## Device
The Device command supports the device capabilities of the Project Flogo CLI.
The commands available are:

build.................... Build a device application
create................... Create a device project
install.................. Install a device contribution
prepare.................. Prepare the device application
upload................... Upload the device application

### build
Build a device application

Usage:
  flogo device build [flags]

Flags:
  -h, --help     help for build
  -g, --no-gen   Only perform the build, without performing the generation of metadata

### create
Create a device project

Usage:
  flogo device create [flags]

Flags:
  -f, --file string   Specify the device.json to create device project from (optional)
  -h, --help          help for create
  -n, --name string   The name of the device project (required)

### install
Install a device contribution

Usage:
  flogo device install [flags]

Flags:
  -h, --help             help for install
  -n, --name string      The name of the contribution (required)
  -v, --version string   Specify the version of the contribution (optional)

### prepare
Prepare the device application

Usage:
  flogo device prepare [flags]

Flags:
  -e, --embed      Embed configuration into the application
  -h, --help       help for prepare
  -o, --optimize   Optimize the preparation

### upload
Upload the device application

Usage:
  flogo device upload [flags]

Flags:
  -h, --help   help for upload

## Plugin
The Plugin command supports the plugin management capabilities of the Project Flogo CLI.
The commands available are:

install.................. Install and add a new plugin to the Project Flogo CLI
list..................... Lists all installed plugins
uninstall................ Uninstall and remove a plugin from the Project Flogo CLI

### install
Install and add a new plugin to the Project Flogo CLI

Usage:
  flogo plugin install [flags]

Flags:
  -h, --help          help for install
  -r, --repo string   The repository you want to install the plugin from (required)

### list
Lists all installed plugins

Usage:
  flogo plugin list [flags]

Flags:
  -h, --help   help for list

### uninstall
Uninstall and remove a plugin from the Project Flogo CLI

Usage:
  flogo plugin uninstall [flags]

Flags:
  -h, --help          help for uninstall
  -r, --repo string   The repository of the plugin you want to uninstall (required)