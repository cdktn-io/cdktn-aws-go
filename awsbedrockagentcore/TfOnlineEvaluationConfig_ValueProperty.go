package awsbedrockagentcore


// Experimental.
type TfOnlineEvaluationConfig_ValueProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#boolean_value TfOnlineEvaluationConfig#boolean_value}.
	// Experimental.
	BooleanValue interface{} `field:"optional" json:"booleanValue" yaml:"booleanValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#double_value TfOnlineEvaluationConfig#double_value}.
	// Experimental.
	DoubleValue *float64 `field:"optional" json:"doubleValue" yaml:"doubleValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#string_value TfOnlineEvaluationConfig#string_value}.
	// Experimental.
	StringValue *string `field:"optional" json:"stringValue" yaml:"stringValue"`
}

