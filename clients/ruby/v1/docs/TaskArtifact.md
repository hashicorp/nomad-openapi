# NomadClient::TaskArtifact

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **getter_headers** | **Hash&lt;String, String&gt;** |  | [optional] |
| **getter_mode** | **String** |  | [optional] |
| **getter_options** | **Hash&lt;String, String&gt;** |  | [optional] |
| **getter_source** | **String** |  | [optional] |
| **relative_dest** | **String** |  | [optional] |

## Example

```ruby
require 'nomad_client'

instance = NomadClient::TaskArtifact.new(
  getter_headers: null,
  getter_mode: null,
  getter_options: null,
  getter_source: null,
  relative_dest: null
)
```

