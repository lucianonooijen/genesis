# Genesis

## Project directory structure

* `/app`: mobile application
  * `/packages`: Typescript packages used throughout the project
      * `/api`: the api sdk connecting the app to the server
      * `/entities`: entities for Typescript code
* `/server`: backend application (api exposed)
* `/docs`: documentation for the application as a whole
* `/bin`: scripts for project-wide actions
* `/.gitlab`: GitLab related configurations


## Requirements

* NodeJS v16 LTS, with Yarn installed globally
* Android Studio for running the Android app
* XCode for running the iOS app
* Yarn 1.13.0+ (`npm i -g yarn@latest`)
* Golang 1.17+

## Development

Instructions for running the back-end can be found in `/server/README.md`.

For the app you can `cd app` and then after running `yarn` you can get a list of all available commands by running `yarn run .

## License

The project is licensed under GPL-3.0, except for the mobile application, which is Apache-2.0 licensed
