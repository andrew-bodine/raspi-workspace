# raspi-release-service

This service implements the raspi release API specification. Raspi release
agents consume this service to implement the continuous release strategy.

## Requirements
- Support fetching latest releases.
- Support watch connections and pushing new releases to connected agents.
- Ability for the caller to specify selection criteria for watch connections.
