package batch


// Experimental.
type AwsComputeEnvironment_UpdatePolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#job_execution_timeout_minutes AwsComputeEnvironment#job_execution_timeout_minutes}.
	// Experimental.
	JobExecutionTimeoutMinutes *float64 `field:"optional" json:"jobExecutionTimeoutMinutes" yaml:"jobExecutionTimeoutMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_compute_environment#terminate_jobs_on_update AwsComputeEnvironment#terminate_jobs_on_update}.
	// Experimental.
	TerminateJobsOnUpdate interface{} `field:"optional" json:"terminateJobsOnUpdate" yaml:"terminateJobsOnUpdate"`
}

