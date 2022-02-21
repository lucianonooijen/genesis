# Genesis

_Bytecode Digital Agency's Project Scaffolding, with the name taken from the Ancient Greek 'γένεσις', meaning origin, generation, source, beginning. Ἐν ἀρχῇ ἦν ὁ ρεπος._

## Usage

* This only works on MacOS
* Make sure Golang 1.17+ is installed
* Make a fork of the Genesis repo and clone to your computer
* Open the Genesis repo and `cd` into `init` (where you are now)
* Run `go run .` and answer the questions
* Replace `/server/static/logo.png` manually
* Commit your changes to the Git repository and push
* There might be some variables that need to be changed by hand later, for example in `app/ios/genesis.xcodeproj/project.pbxproj`: `DEVELOPMENT_TEAM = JS5JAXY9DB;`
* Follow instructions in `/docs` to set up the CI/CD and push notifications
* Change the contents of `android/sentry.properties` and `ios/sentry.properties` to connect to Sentry

## Stuff that should be replaced in Scaffold step

### Variables in code

* Genesis
* genesis
* @genesis/app
* @genesis/api
* nl.bytecode.genesis
* GenesisTests
* git.bytecode.nl/bytecode/genesis/server
* https://bytecode.nl
* https://placekitten.com/400/400
* development@genesis.bytecode.dev
* registry.digitalocean.com/dawny/genesis-server

### FolderNames
* app/android/app/src/debug/java/com/genesis
* app/android/app/src/main/java/com/genesis
* app/ios/Genesis (and other directories in app/ios)

## Files

* `/app/ios/Genesis.xcodeproj`
* `/app/ios/Genesis.xcodeproj/xcshareddata/xcschemes/Genesis.xcscheme`
* `/server/static/logo.png`
