package awsconfig


// Experimental.
type AwsConfigConfigRule_CustomPolicyDetailsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#policy_runtime AwsConfigConfigRule#policy_runtime}.
	// Experimental.
	PolicyRuntime *string `field:"required" json:"policyRuntime" yaml:"policyRuntime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#policy_text AwsConfigConfigRule#policy_text}.
	// Experimental.
	PolicyText *string `field:"required" json:"policyText" yaml:"policyText"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/config_config_rule#enable_debug_log_delivery AwsConfigConfigRule#enable_debug_log_delivery}.
	// Experimental.
	EnableDebugLogDelivery interface{} `field:"optional" json:"enableDebugLogDelivery" yaml:"enableDebugLogDelivery"`
}

