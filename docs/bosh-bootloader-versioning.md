### What is this?

This is an outline of what I (Nick M) believe we need to correctly version releases with our bbl pipeline. This has been tried many times, but we always find some problem. I suspect that one factor is people jumping in and trying to implement it without a clear idea of the necessary design. I will attempt to lay that design out clearly here.

### The problems we have run into

- making ourselves remember to bump the version *before* the main pipeline starts
- the bump-deployments pipeline publishing using the latest release's major and minor versions, rather than the major and minor versions of the release that it was built on top of.

### The design

Relevant jobs:
1. **main** pipeline: cut github draft release
2. **bump-deployments** pipeline: publish github release
3. **main** pipeline: manually bump major version
4. **main** pipeline: manually bump minor version
5. **bump-deployments** pipeline: whichever job pulls in the latest github release

Relevant resources:
1. main pipeline semver
2. full semver of release that bump-deployments pipeline is building on top of

Requirements:
1. job `1` should cut a github draft release with version in resource `1` without any `passed` constraints (so that we can bump the version right before cutting the release rather than needing to do it before the whole pipeline runs)
2. job `5` should get the full semver of the release it pulls in, bump the patch version, and `put` that to resource `2`
3. resource `2` should be passed through the bump-deployments pipeline with `passed` constraints (so that job `2` doesn't use the wrong semver in the case where the pipeline kicks off again with a new release)
4. job `2` should publish a github release with the version in resource `2` (with `passed` constraint)
5. job `3` should bump major version of resource `1` (resetting minor and patch to `0`)
6. job `4` should bump minor version of resource `1` (leaving major the same and resetting patch to `0`)

Note: Although I talked about resetting the patch version of resource `1` (in requirements `5` and `6`), it should always be set to `0`. Patch versions are only bumped by the bump-deployments pipeline, which uses resource `2`.
