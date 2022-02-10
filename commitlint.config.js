module.exports = {
  extends: ["@commitlint/config-conventional"],
  rules: {
    'type-enum': [
      2,
      'always',
      ['build',
        'ci',
        'docs',
        'feat',
        'fix',
        'perf',
        'refactor',
        'release',
        'revert',
        'style',
        'test']
    ],
    'scope-enum': [
      2,
      'always',
      ['app', 'docs', 'init', 'server']
    ]
  },
  ignores: [(commit) => commit.includes("Merge")],
}