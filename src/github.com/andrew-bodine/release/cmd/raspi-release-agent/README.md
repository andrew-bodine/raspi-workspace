# raspi-release-agent

The raspi-release-agent is the local controller that is responsible for deploying
component releases from the raspi-release-proxy. The agent always seeks to keep
components as up-to-date as it can, and supports automatic rollback if anything
goes wrong during a deployment.

## Non-Functional Requirements
[ x ] The agent itself must be stateless.
[ ] The agent must support a runtime config which gives it a source-of-truth as to
what components it should be managing on the system via raspi release.
[ ] The agent must start at system start.
[ ] The agent must support reloading config at runtime.

## Functional Requirements
[ ] The agent must support downloading the specific release binaries from Github
API, and will need an appropriate API key provided at runtime to enable this.
[ ] The agent must discover what current releases it has applied on it's local
system at startup time, and use this information in it's initial release cache
sync with the raspi-release-proxy.
[ ] The agent must connect to a raspi-release-proxy service at startup, and watch
for new release notifications from the proxy service.
[ ] For any component only a single release can be getting applied to the system
at any time.
[ ] When discovering or being notified of a new release, the agent should ...
[ ] Must support exponential backoff in the face of getting rate limited by Github API.
[ ] The agent must be able to update itself when notified about a new
raspi-release-agent release.
[ ] If there is a new release while one for the same component is currently being
applied, the agent will finish applying the current in-flight release before
beginning the latest release deployment. Furthermore, if there are multiple new
releases for the same component while one is already currently being applied,
the agent will pick the latest one to apply next and disregard the others.
[ ] If an error is encountered while applying a component release, the agent will
roll the component back to the last known successful release, and the failed
release should be whitelisted in memory. If there are any other newer releases
that the agent has been informed of, the agent will attempt deployment of those
newest to oldest until it lands back on the current release.
[ ] The agent must send telemetry to the raspi-release-proxy:
  + When beginning to apply a release.
  + After applying a component release, regardless of success or failure.
