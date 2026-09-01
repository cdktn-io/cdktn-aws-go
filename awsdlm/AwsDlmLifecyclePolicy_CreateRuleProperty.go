package awsdlm


// Experimental.
type AwsDlmLifecyclePolicy_CreateRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#cron_expression AwsDlmLifecyclePolicy#cron_expression}.
	// Experimental.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#interval AwsDlmLifecyclePolicy#interval}.
	// Experimental.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#interval_unit AwsDlmLifecyclePolicy#interval_unit}.
	// Experimental.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#location AwsDlmLifecyclePolicy#location}.
	// Experimental.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// scripts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#scripts AwsDlmLifecyclePolicy#scripts}
	// Experimental.
	Scripts *AwsDlmLifecyclePolicy_ScriptsProperty `field:"optional" json:"scripts" yaml:"scripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#times AwsDlmLifecyclePolicy#times}.
	// Experimental.
	Times *[]*string `field:"optional" json:"times" yaml:"times"`
}

