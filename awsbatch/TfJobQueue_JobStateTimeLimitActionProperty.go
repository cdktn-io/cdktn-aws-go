package awsbatch


// Experimental.
type TfJobQueue_JobStateTimeLimitActionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#action TfJobQueue#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#max_time_seconds TfJobQueue#max_time_seconds}.
	// Experimental.
	MaxTimeSeconds *float64 `field:"required" json:"maxTimeSeconds" yaml:"maxTimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#reason TfJobQueue#reason}.
	// Experimental.
	Reason *string `field:"required" json:"reason" yaml:"reason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_queue#state TfJobQueue#state}.
	// Experimental.
	State *string `field:"required" json:"state" yaml:"state"`
}

