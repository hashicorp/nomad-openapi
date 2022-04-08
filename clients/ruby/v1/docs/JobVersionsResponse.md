# NomadClient::JobVersionsResponse

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **diffs** | [**Array&lt;JobDiff&gt;**](JobDiff.md) |  | [optional] |
| **known_leader** | **Boolean** |  | [optional] |
| **last_contact** | **Integer** |  | [optional] |
| **last_index** | **Integer** |  | [optional] |
| **next_token** | **String** |  | [optional] |
| **request_time** | **Integer** |  | [optional] |
| **versions** | [**Array&lt;Job&gt;**](Job.md) |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::JobVersionsResponse.new(
  diffs: null,
  known_leader: null,
  last_contact: null,
  last_index: null,
  next_token: null,
  request_time: null,
  versions: null
)
```

