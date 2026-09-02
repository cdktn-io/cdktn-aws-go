package awsglue


// Experimental.
type TfTrigger_EventBatchingConditionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#batch_size TfTrigger#batch_size}.
	// Experimental.
	BatchSize *float64 `field:"required" json:"batchSize" yaml:"batchSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_trigger#batch_window TfTrigger#batch_window}.
	// Experimental.
	BatchWindow *float64 `field:"optional" json:"batchWindow" yaml:"batchWindow"`
}

