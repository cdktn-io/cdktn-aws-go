package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreEvaluator_LlmAsAJudgeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#instructions AwsBedrockagentcoreEvaluator#instructions}.
	// Experimental.
	Instructions *string `field:"required" json:"instructions" yaml:"instructions"`
	// model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#model_config AwsBedrockagentcoreEvaluator#model_config}
	// Experimental.
	ModelConfig interface{} `field:"optional" json:"modelConfig" yaml:"modelConfig"`
	// rating_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#rating_scale AwsBedrockagentcoreEvaluator#rating_scale}
	// Experimental.
	RatingScale interface{} `field:"optional" json:"ratingScale" yaml:"ratingScale"`
}

