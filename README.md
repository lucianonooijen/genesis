# Genesis

_Bytecode Digital Agency's Project Scaffolding, with the name taken from the Ancient Greek 'γένεσις', meaning origin, generation, source, beginning. Ἐν ἀρχῇ ἦν ὁ ρεπος._

## Project directory structure

* `/app`: mobile application
* `/packagves`: Typescript packages used throughout the project
  * `/api`: the api sdk connectiong the app to the server
* `/server`: backend application (api exposed)
* `/docs`: documentation for the application as a whole
* `/bin`: scripts for project-wide actions
* `/.gitlab`: GitLab related configurations

## Stuff that should be replaced in Scaffold step

### Variables in code

* Genesis
* genesis
* @genesis/app
* @genesis/api
* com.genesis
* GenesisTests
* git.bytecode.nl/bytecode/genesis
* git@git.bytecode.nl:bytecode/genesis.git
* https://bytecode.nl
* https://placekitten.com/400/400

### FolderNames
* app/android/app/src/debug/java/com/genesis
* app/android/app/src/main/java/com/genesis
* app/ios/Genesis

## Files

* `/app/ios/Genesis.xcodeproj`
* `/app/ios/Genesis.xcodeproj/xcshareddata/xcschemes/Genesis.xcscheme`
* `/server/static/logo.png`

---

_Project documentation below_

# Genesis

## Requirements

* NodeJS v16 LTS
* Yarn 1.13.0+ (`npm i -g yarn@latest`)
* Golang 1.17+

## Development

Instructions for running the back-end can be found in `/server/README.md`.

The mobile application uses [Yarn Workspaces](https://classic.yarnpkg.com/lang/en/docs/workspaces/) to organize Typescipt code and to separate diffrent packages.

Because Yarn Workspaces is used, you will have to run all `yarn` commands from the root of the repository.

Run `yarn install && yarn run` to see which commands are available.

## License

The project is licensed under GPL-3.0, except for the mobile application, which is Apache-2.0 licensed
