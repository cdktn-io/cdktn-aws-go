package awssesmailmanager


// Experimental.
type TfRuleSet_InvokeLambdaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#function_arn TfRuleSet#function_arn}.
	// Experimental.
	FunctionArn *string `field:"required" json:"functionArn" yaml:"functionArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#invocation_type TfRuleSet#invocation_type}.
	// Experimental.
	InvocationType *string `field:"required" json:"invocationType" yaml:"invocationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#role_arn TfRuleSet#role_arn}.
	// Experimental.
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#action_failure_policy TfRuleSet#action_failure_policy}.
	// Experimental.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_rule_set#retry_time_minutes TfRuleSet#retry_time_minutes}.
	// Experimental.
	RetryTimeMinutes *float64 `field:"optional" json:"retryTimeMinutes" yaml:"retryTimeMinutes"`
}

