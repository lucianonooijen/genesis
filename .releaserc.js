module.exports = {
    branches: ["master", "cicd"],
    plugins: [
        ["@semantic-release/commit-analyzer", {
            preset: "conventionalcommits",
            releaseRules: [
                {breaking: true, release: "major"},
                {revert: true, release: "patch"},
                {type: "feat", release: "minor"},
                {type: "fix", release: "patch"},
                {type: "perf", release: "patch"},
                {type: "build", release: "patch"},
                {type: "ci", release: "patch"},
                {type: "docs", release: "patch"}
            ]
        }],
        "@semantic-release/release-notes-generator",
        "@semantic-release/changelog",
        // "@semantic-release/gitlab",
        // ["@semantic-release/git", {
        //     message: "release: ${nextRelease.version}\n\n${nextRelease.notes}",
        //     assets: ["CHANGELOG.md"]
        // }]
    ]
}
