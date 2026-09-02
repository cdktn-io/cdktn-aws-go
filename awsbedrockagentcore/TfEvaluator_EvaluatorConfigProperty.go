package awsbedrockagentcore


// Experimental.
type TfEvaluator_EvaluatorConfigProperty struct {
	// code_based block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#code_based TfEvaluator#code_based}
	// Experimental.
	CodeBased interface{} `field:"optional" json:"codeBased" yaml:"codeBased"`
	// llm_as_a_judge block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_evaluator#llm_as_a_judge TfEvaluator#llm_as_a_judge}
	// Experimental.
	LlmAsAJudge interface{} `field:"optional" json:"llmAsAJudge" yaml:"llmAsAJudge"`
}

