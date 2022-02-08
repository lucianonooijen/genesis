# CI/CD configuration

Instructions to set up the CI/CD for a new GitLab repository

## Required GitLab CI variables

Make sure these keys are set for the project in the CI/CD -> Variables section.

| Key                   | Description                                                    |
|-----------------------|----------------------------------------------------------------|
| `DO_REGISTRY_API_KEY` | DigitalOcean Docker Registry API key for pushing Docker images |
| `DOCKER_AUTH`         | Docker Authentication key to avoid getting rate-limited        |
| `EXPO_TOKEN`          | Access token for Expo/EAS to trigger app builds                |
| `GITLAB_TOKEN`        | Project Access Token used by the Semantic Release step         |
