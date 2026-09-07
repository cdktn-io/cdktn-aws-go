package bedrockagentcore


// Experimental.
type AwsEvaluator_RatingScaleProperty struct {
	// categorical block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#categorical AwsEvaluator#categorical}
	// Experimental.
	Categorical interface{} `field:"optional" json:"categorical" yaml:"categorical"`
	// numerical block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#numerical AwsEvaluator#numerical}
	// Experimental.
	Numerical interface{} `field:"optional" json:"numerical" yaml:"numerical"`
}

