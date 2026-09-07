package bedrockagentcore


// Experimental.
type AwsOnlineEvaluationConfig_RuleProperty struct {
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#filter AwsOnlineEvaluationConfig#filter}
	// Experimental.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// sampling_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#sampling_config AwsOnlineEvaluationConfig#sampling_config}
	// Experimental.
	SamplingConfig interface{} `field:"optional" json:"samplingConfig" yaml:"samplingConfig"`
	// session_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#session_config AwsOnlineEvaluationConfig#session_config}
	// Experimental.
	SessionConfig interface{} `field:"optional" json:"sessionConfig" yaml:"sessionConfig"`
}

