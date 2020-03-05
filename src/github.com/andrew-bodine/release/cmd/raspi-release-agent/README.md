# raspi-release-agent

Responsible for workspace application release management on the Raspberry Pi
where it is located.

## Release and Release Packages

A `Release` can be defined with the following criteria:
- Named according to this pattern: `TODO`
- Contains one or more release packages, a.k.a workspace release target contexts.
- Optionally defines post deploy scripting.
- Compressed tarball.

A `Release Package` can be defined with the following criteria:
- Contains a systemd unit definition for the release target.
- Contains a workspace release target binary.
- Contains automated documentation for the release target.

## Integration

1. Changes merge to master branch of this workspace.
1. Travis CI job initiates to run all workspace checks.
1. Release is created that contains each enabled workspace target.
1. Release is uploaded to workspace repo releases index.

## Delivery

1. Raspi release agent is installed on a Raspberry Pi.
1. Raspi release agent is configured with desired release configs.
1. Raspi release agent detects a new releases available.
1. Agent pulls latest release and applies it to local filesystem.
1. Agent guarantees systemd units are healthy per release targets.
1. Agent runs release post scripting if provided.
