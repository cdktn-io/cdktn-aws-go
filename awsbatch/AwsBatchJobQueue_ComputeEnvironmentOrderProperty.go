package awsbatch


// Experimental.
type AwsBatchJobQueue_ComputeEnvironmentOrderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#compute_environment AwsBatchJobQueue#compute_environment}.
	// Experimental.
	ComputeEnvironment *string `field:"required" json:"computeEnvironment" yaml:"computeEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#order AwsBatchJobQueue#order}.
	// Experimental.
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

