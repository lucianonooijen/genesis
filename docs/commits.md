# Commit rules

Commit names are used to automatically generate changelogs. Therefore, commit titles must conform to a specific format.

The tool used for checking the commits is [commitlint](https://commitlint.js.org/).

The format used is:

```
type(optional scope): details
```

For example

```
feat(app): add jwt refresh if the token is valid for less than a week
```

## Types and scopes

The possible types are:

- build
- ci
- docs
- feat
- fix
- perf
- refactor
- release
- revert
- style
- test

The possible scopes are:

- app
- docs
- init
- server
