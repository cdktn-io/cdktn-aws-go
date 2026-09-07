package batch


// Experimental.
type AwsJobQueue_ComputeEnvironmentOrderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#compute_environment AwsJobQueue#compute_environment}.
	// Experimental.
	ComputeEnvironment *string `field:"required" json:"computeEnvironment" yaml:"computeEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#order AwsJobQueue#order}.
	// Experimental.
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

