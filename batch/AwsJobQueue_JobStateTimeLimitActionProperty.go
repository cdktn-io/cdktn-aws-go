package batch


// Experimental.
type AwsJobQueue_JobStateTimeLimitActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#action AwsJobQueue#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#max_time_seconds AwsJobQueue#max_time_seconds}.
	// Experimental.
	MaxTimeSeconds *float64 `field:"required" json:"maxTimeSeconds" yaml:"maxTimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#reason AwsJobQueue#reason}.
	// Experimental.
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#state AwsJobQueue#state}.
	// Experimental.
	State *string `field:"required" json:"state" yaml:"state"`
}

