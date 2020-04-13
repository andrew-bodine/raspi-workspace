# raspi-release-proxy

This service implements the raspi release API specification. Raspi release
agents consume this service to implement the continuous release strategy.

## Non-Functional Requirements
[ ] Must support a configurable number of releases per component to cache. Assuming
an average release size of 10M, if you want to support 1,000 components its
going to take 1G of memory to store 1 copy of each component release.
[ ] Nothing written to disk, if this service restart it has to rebuild its cache.
[ ] Must start at system start.

## Functional Requirements
[ ] Build release cache from Github upon startup.
[ ] Periodically and well under the Github API rate limit, the proxy should poll
for new releases from the raspi workspace.
[ ] Support fetching latest releases with filter by component support.
[ ] Support watch connections and pushing new releases to connected agents.
[ ] Ability for the caller to specify selection criteria for watch connections,
allowing them to designate specific components they are interested in, and
which versions they already have.
[ ] Keep track of release agent metrics, will require storage:
  + Connected agents
  + Component release application success/failure history
  + Component release application timings
    - when agent was notified
    - when agent reported release application start
    - when agent reported release application result
[ ] Once telemetry is in place, when an agent notifies proxy of a failed release
that release should be whitelisted and not sent to that agent.
[ ] Logs should be shipped somewhere useful and sensical.
[ ] Must support exponential backoff in the face of getting rate limited by Github API.
[ ] Support webhook callback so Github can notify proxy of new releases.
[ ] Support downloading release artifacts from Github API based on release meta from proxy.

## Roadmap
[ ] Mutual Authentication and Authorization
