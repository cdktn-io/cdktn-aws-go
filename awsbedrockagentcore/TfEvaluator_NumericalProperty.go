package awsbedrockagentcore


// Experimental.
type TfEvaluator_NumericalProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#definition TfEvaluator#definition}.
	// Experimental.
	Definition *string `field:"required" json:"definition" yaml:"definition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#label TfEvaluator#label}.
	// Experimental.
	Label *string `field:"required" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#value TfEvaluator#value}.
	// Experimental.
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

