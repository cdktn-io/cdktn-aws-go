package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreOnlineEvaluationConfig_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#key AwsBedrockagentcoreOnlineEvaluationConfig#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#operator AwsBedrockagentcoreOnlineEvaluationConfig#operator}.
	// Experimental.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#value AwsBedrockagentcoreOnlineEvaluationConfig#value}
	// Experimental.
	Value interface{} `field:"optional" json:"value" yaml:"value"`
}

