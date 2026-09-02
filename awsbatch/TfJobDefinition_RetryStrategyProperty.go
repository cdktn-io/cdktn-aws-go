package awsbatch


// Experimental.
type TfJobDefinition_RetryStrategyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#attempts TfJobDefinition#attempts}.
	// Experimental.
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
	// evaluate_on_exit block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/batch_job_definition#evaluate_on_exit TfJobDefinition#evaluate_on_exit}
	// Experimental.
	EvaluateOnExit interface{} `field:"optional" json:"evaluateOnExit" yaml:"evaluateOnExit"`
}

