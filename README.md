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

## Git Branching & Release Strategy

1. **main**: Production-ready code. Only updated via pull requests from release or hotfix branches.
2. **develop**: Integration branch for features. All feature branches are merged here.
3. **release/x.y.z**: Created from `develop` when preparing a new release.
   - Final fixes, version bumps, and documentation updates are made here.
   - A tag (e.g., `v1.2.0`) is created and pushed to trigger CI/CD.
   - CI workflow creates a draft release.
   - Pull request from `release/x.y.z` to `main` for final review.
   - After review, publish the release and merge to `main`.
   - Merge back to `develop` if there were changes.
4. **hotfix/x.y.z**: For urgent fixes to production. Branched from `main`, then merged into both `main` and `develop`.

**Release Steps:**
1. Branch from `develop` to `release/x.y.z`
2. Finalize release (fixes, docs, version bump)
3. Create and push tag `vX.Y.Z`
4. CI creates draft release
5. PR from `release/x.y.z` to `main`
6. Review and publish draft release
7. Merge release to `main` (and back to `develop` if needed)

### Back merge
you do not always need to back-merge from main (or master) to develop after every release, but you should do so whenever there have been changes or fixes made directly to main that are not yet in develop. This is a core part of the Git Flow model and is considered best practice in many team workflows

.
When should you back-merge from main to develop?

*   After a release:
    If you made any final changes, fixes, or version bumps on the release branch before merging into main, those changes should also be merged back into develop to keep both branches in sync

*   After a hotfix:
    Hotfixes are usually branched from main to fix urgent production issues. Once merged into main, the hotfix should also be merged into develop to ensure ongoing development includes the fix


### Additional Information

For more details, documentation, and advanced usage, please visit the [project repository](https://github.com/pezhmankasraee/pksetdev).

---