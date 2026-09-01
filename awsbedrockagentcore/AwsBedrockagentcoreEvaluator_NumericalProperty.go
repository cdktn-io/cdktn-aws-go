package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreEvaluator_NumericalProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#definition AwsBedrockagentcoreEvaluator#definition}.
	// Experimental.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#label AwsBedrockagentcoreEvaluator#label}.
	// Experimental.
	Label *string `field:"required" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#value AwsBedrockagentcoreEvaluator#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

