package awsbedrockagentcore


// Experimental.
type TfOnlineEvaluationConfig_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#key TfOnlineEvaluationConfig#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#operator TfOnlineEvaluationConfig#operator}.
	// Experimental.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#value TfOnlineEvaluationConfig#value}
	// Experimental.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

