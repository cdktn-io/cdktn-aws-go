package awsbatch


// Experimental.
type AwsBatchJobQueue_JobStateTimeLimitActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#action AwsBatchJobQueue#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#max_time_seconds AwsBatchJobQueue#max_time_seconds}.
	// Experimental.
	MaxTimeSeconds *float64 `field:"required" json:"maxTimeSeconds" yaml:"maxTimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#reason AwsBatchJobQueue#reason}.
	// Experimental.
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#state AwsBatchJobQueue#state}.
	// Experimental.
	State *string `field:"required" json:"state" yaml:"state"`
}

