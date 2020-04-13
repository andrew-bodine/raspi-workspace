# raspi-release

Responsible for workspace application release management on the Raspberry Pi
where it is located.

## Definition of a Release

A `Release` can be defined with the following criteria:
- Named according to this pattern: `TODO`
- Contains a systemd unit definition for the release target.
- Contains a workspace release target binary.
- Contains automated documentation for the release target.
- Optionally defines post deploy scripting.
- Compressed tarball.

## CI

When new changes are proposed in a PR:
1. Travis CI job initiates to run all workspace checks.
1. Upon successful Travis CI run the PR is able to merge.

## CD

One-time steps per environment:
1. Raspi release agent is installed on a Raspberry Pi.
1. Raspi release agent is configured with desired release configs.
1. Raspi release agent connects to raspi release proxy.

Upon a merge to master:
1. A Release is created for each enabled workspace target.
1. Releases are uploaded to workspace repo releases index.
1. Raspi release proxy detects a one or more new releases available.
1. Raspi release proxy pushes metadata for new releases to connected agents.
1. Agent pulls latest releases and apply them to local filesystem.
1. Agent guarantees systemd units are healthy per release targets.
1. Agent runs release post scripting if provided.
