
# flake8: noqa

# Import all APIs into this package.
# If you have many APIs here with many many models used in each API this may
# raise a `RecursionError`.
# In order to avoid this, import only the API that you directly need like:
#
#   from nomad_client.api.acl_api import ACLApi
#
# or import this package, but before doing it, use:
#
#   import sys
#   sys.setrecursionlimit(n)

# Import APIs into API package:
from nomad_client.api.acl_api import ACLApi
from nomad_client.api.allocations_api import AllocationsApi
from nomad_client.api.deployments_api import DeploymentsApi
from nomad_client.api.enterprise_api import EnterpriseApi
from nomad_client.api.evaluations_api import EvaluationsApi
from nomad_client.api.jobs_api import JobsApi
from nomad_client.api.metrics_api import MetricsApi
from nomad_client.api.namespaces_api import NamespacesApi
from nomad_client.api.nodes_api import NodesApi
from nomad_client.api.operator_api import OperatorApi
from nomad_client.api.plugins_api import PluginsApi
from nomad_client.api.regions_api import RegionsApi
from nomad_client.api.scaling_api import ScalingApi
from nomad_client.api.search_api import SearchApi
from nomad_client.api.status_api import StatusApi
from nomad_client.api.system_api import SystemApi
from nomad_client.api.volumes_api import VolumesApi
