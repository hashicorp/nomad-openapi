import nomad_client
import nomad_client.apis
from pprint import pprint


nc_config = nomad_client.Configuration(
    host="http://127.0.0.1:4646/v1"
)

nc = nomad_client.ApiClient(configuration=nc_config)
jobsApi = nomad_client.apis.JobsApi(nc)

try:
    jobs = jobsApi.get_jobs()    
    pprint(jobs)
    
    job = jobsApi.get_job(job_name="example")    
    pprint(job)
except nomad_client.ApiException as e:
    print("Exception when calling JobsApi.get_jobs: %s\n" % e)

