package awsbedrockagentcore


// Experimental.
type TfEvaluator_LlmAsAJudgeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#instructions TfEvaluator#instructions}.
	// Experimental.
	Instructions *string `field:"required" json:"instructions" yaml:"instructions"`
	// model_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#model_config TfEvaluator#model_config}
	// Experimental.
	ModelConfig interface{} `field:"optional" json:"modelConfig" yaml:"modelConfig"`
	// rating_scale block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#rating_scale TfEvaluator#rating_scale}
	// Experimental.
	RatingScale interface{} `field:"optional" json:"ratingScale" yaml:"ratingScale"`
}

