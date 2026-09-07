package dlm


// Experimental.
type AwsLifecyclePolicy_CreateRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#cron_expression AwsLifecyclePolicy#cron_expression}.
	// Experimental.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#interval AwsLifecyclePolicy#interval}.
	// Experimental.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#interval_unit AwsLifecyclePolicy#interval_unit}.
	// Experimental.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#location AwsLifecyclePolicy#location}.
	// Experimental.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// scripts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#scripts AwsLifecyclePolicy#scripts}
	// Experimental.
	Scripts *AwsLifecyclePolicy_ScriptsProperty `field:"optional" json:"scripts" yaml:"scripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#times AwsLifecyclePolicy#times}.
	// Experimental.
	Times *[]*string `field:"optional" json:"times" yaml:"times"`
}

