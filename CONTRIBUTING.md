# Contributing to homer-service-discovery

We use GitHub to host code, build code, create releases, track issues and feature requests, and accept pull requests.

## MIT Software License

In short, when you submit code changes, your submissions are understood to be under the same [MIT License](https://github.com/calvinbui/homer-service-discovery/blob/master/LICENSE) that covers the project. Feel free to contact the maintainers if that's a concern.

## Report Bugs or Request Features using GitHub Issues

We use GitHub issues to raise any bug or request new features. GitHub will present templates for both.

## Contribution Process

Pull requests are the best way to propose changes to the codebase (we use [GitHub Flow](https://guides.github.com/introduction/flow/index.html)). We actively welcome your pull requests:

1. Fork this repo and create a branch from master.
2. If you've added code that should be tested, add tests.
3. If you've changed the config, update the README.md file.
4. Ensure the test suite passes.
5. Make sure your code lints.
6. Issue that pull request!

## Release Process

When there are code changes from myself or from a pull request, a tag will be created following semantic versioning. This will trigger the CI to generate all the artifacts for commit.

## Dependency Management

This project uses [Go modules](https://golang.org/cmd/go/#hdr-Modules__module_versions__and_more) to manage dependencies on external packages

To add or update a new dependency, use the `go get` command:

```bash
# Pick the latest tagged release.
go get example.com/some/module/pkg

# Pick a specific version.
go get example.com/some/module/pkg@vX.Y.Z
```
You have to commit the changes to `go.mod` and `go.sum` before submitting the pull request.
