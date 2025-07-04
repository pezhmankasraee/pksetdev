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

This is the flag definitions:

---

## Usage

```sh
pksetdev [OPTIONS]
```

### Options

| Option                            | Description                       |
|-----------------------------------|-----------------------------------|
| `-i`, `--install`               | Download the applications and install 
| `-p PATH`, `--path PATH`               | Path to the YAML config file. This file should define device settings and parameters. `Example: ./config/config.yaml` configuration file.    |
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

## Git Branching & Release Strategy

1. **main**: Production-ready code. Only updated via pull requests from release or hotfix branches.
2. **develop**: Integration branch for features. All feature branches are merged here.
3. **release-vM.m.p**: Created from `develop` when preparing a new release.
4. **hotfix/x.y.z**: For urgent fixes to production. Branched from `main`, then merged into both `main` and `develop`.

**Release Steps:**
1. Create a branch from `develop` to `release-vM.m.p`
2. Push the branch `release-vM.m.p`
3. The push will trigger the `create-tag` workflow/actin, and it will create the tag with version `vM.m.p` which points to the head of of `release-vM.m.p` branch
4. Create PR from `release-vM.m.p` to `main`, creating PR will trigger `draft-release` workflow/action and it creates draft release
5. Do NOT merge the release branch directly to `main` instead do a fast-forward-merge locally.
```bash
$ git checkout main
$ git merge --ff-only release vM.m.p
$ git push origin merge
```
This will close the PR which is already open cleanly.
Do not forget to delete the release branch after pushing the main.

6. Review and publish draft release

### Back merge (This should be discussed further)
you do not always need to back-merge from main (or master) to develop after every release, but you should do so whenever there have been changes or fixes made directly to main that are not yet in develop. This is a core part of the Git Flow model and is considered best practice in many team workflows

When should you back-merge from main to develop?

*   After a release:
    If you made any final changes, fixes, or version bumps on the release branch before merging into main, those changes should also be merged back into develop to keep both branches in sync

*   After a hotfix:
    Hotfixes are usually branched from main to fix urgent production issues. Once merged into main, the hotfix should also be merged into develop to ensure ongoing development includes the fix


### Additional Information

For more details, documentation, and advanced usage, please visit the [project repository](https://github.com/pezhmankasraee/pksetdev).

---