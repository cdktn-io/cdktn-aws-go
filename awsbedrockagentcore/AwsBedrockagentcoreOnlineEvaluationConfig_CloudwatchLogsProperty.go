package awsbedrockagentcore


// Experimental.
type AwsBedrockagentcoreOnlineEvaluationConfig_CloudwatchLogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#log_group_names AwsBedrockagentcoreOnlineEvaluationConfig#log_group_names}.
	// Experimental.
	LogGroupNames *[]*string `field:"required" json:"logGroupNames" yaml:"logGroupNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_online_evaluation_config#service_names AwsBedrockagentcoreOnlineEvaluationConfig#service_names}.
	// Experimental.
	ServiceNames *[]*string `field:"required" json:"serviceNames" yaml:"serviceNames"`
}

