<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

# bins

Assert everyone on your team (or CI server) has required binaries installed for the project. Makes onboarding with a new project a more positive experience. 

Instead of creating a README that lists all of the binaries that your project uses that all contributors must install, all you need to do is tell the new contributors to install `bins` and run `bin install`. 

# Getting started 

### 1. Install 
// TODO 

### 2. Configure `bins`

* Add a file `.bins.yml` to the root of your project. 
* In the file, configure all the binaries that your project requires. 

```yml
bins:
  - binary: sourcery # the name of the CLI to install 
    version: 
      requirement: "^2.0.0" # saying sourcery minimum version required is 2.0
    # Installers define all of the installers that are able to install this binary. 
    # Such as npm, gem, homebrew, snap, etc. 
    # Read more about installers below in the documentation. 
    installers: [brew] # use homebrew to install sourcery

  - binary: lefthook 
    version:
      requirement: "^0.7"
      command: "lefthook version" # command to run to get the version of lefthook installed 
    installers: [brew]
    postInstall: # Run commands after installing is successful. Install git hooks, for example. 
      - command: "lefthook install"

  - binary: fastlane
    version:
      requirement: "^2.0.0"
      command: fastlane --version
      # If you need to set environment variables when running `fastlane --version`, set them below: 
      commandEnvVars: [FASTLANE_SKIP_UPDATE_CHECK=true] 
    installers: [gem, brew] # use ruby gem or homebrew to install 

# Where you define the behavior of programs used to install binaries (npm, apt-get, gem, homebrew, etc)
# If you want to use the installers: bin, gem 
# then just add `brew` or `gem` in the `installers: []` list in your config file like you see above. 
# bins comes bundled with these installers so you don't need to define any custom installers. 
# If you want to use a separate installer besides these, you need to define them yourself 
# such as you see below: 
installer:
  # id that gets referenced in config file: `installers: [id-here]`
  - id: brew
    # the name of the CLI that you call after it's installed 
    binary: brew 
    # If user needs to *also* install the installer, this command will be run to install homebrew for them. 
    installCommand: "/bin/bash -c \"$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)\""
    # Template string that defines how homebrew installs new programs. 
    # Template gets access to all variables that are defined in `bins` above. 
    # So, {{.Binary}} is `sourcery`. 
    # {{.Version.Requirement}} is `^2.0.0`
    installTemplate: "brew install {{.Binary}}"
    # Like `installTemplate` but instead the template string for how to upgrade an existing binary. 
    updateTemplate: "brew upgrade {{.Binary}}"
```

# Features 

### Interactive install process of missing binaries 

`bins` has a feature that checks what binaries need installed or upgraded on your local machine. Then, it walks you through installing or upgrading them all for you in an interactive way. A great experience for new users of your project! 

All you need to do is run: 

```
bins install
```

> Tip: Run `bins install` in your git hooks before running a binary tool. Great way to make sure that others do not have a negative experience using your git hooks or are using an old version of the CLI tool the hooks uses. 

### Install required bins on CI server 

Make your CI server setup easier by installing all required binaries with just 1 line of code on your CI server:

```
bins install 
```

`bins` looks for the `CI` environment variable. If it set (most CI server providers have this set already), `bins` will install binaries in a non-interactive mode. 

### Assert required bins installed with minimum version met 

If all you want to do is check if your machine has all of the required binaries installed and meet the minimum version:

```
bins assert
```

`bins` will return exit code 0 if all binaries are installed and the version of the binary is satisfied. 

## Development 

bins is a Go lang program. To start developing for bins is as simple as (1) cloning the repo and (2) running `go run main.go --debug`.

## Contribute

bins is open for pull requests. Check out the [list of issues](https://github.com/levibostian/bins/issues) for tasks planned to be worked on. Check them out if you wish to contribute in that way.

**Want to add features to bins?** Before you decide to take a bunch of time and add functionality to the CLI, please, [create an issue](https://github.com/levibostian/bins/issues/new) stating what you wish to add. This might save you some time in case your purpose does not fit well in the use cases of bins.

## Contributors 

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key))

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/levibostian"><img src="https://avatars1.githubusercontent.com/u/2041082?v=4" width="100px;" alt=""/><br /><sub><b>Levi Bostian</b></sub></a><br /><a href="https://github.com/levibostian/bins/commits?author=levibostian" title="Code">💻</a> <a href="https://github.com/levibostian/bins/commits?author=levibostian" title="Documentation">📖</a> <a href="#maintenance-levibostian" title="Maintenance">🚧</a></td>
  </tr>
</table>

<!-- markdownlint-enable -->
<!-- prettier-ignore-end -->
<!-- ALL-CONTRIBUTORS-LIST:END -->
