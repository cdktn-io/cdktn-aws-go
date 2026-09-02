package awsdlm


// Experimental.
type TfLifecyclePolicy_CreateRuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#cron_expression TfLifecyclePolicy#cron_expression}.
	// Experimental.
	CronExpression *string `field:"optional" json:"cronExpression" yaml:"cronExpression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#interval TfLifecyclePolicy#interval}.
	// Experimental.
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#interval_unit TfLifecyclePolicy#interval_unit}.
	// Experimental.
	IntervalUnit *string `field:"optional" json:"intervalUnit" yaml:"intervalUnit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#location TfLifecyclePolicy#location}.
	// Experimental.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// scripts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#scripts TfLifecyclePolicy#scripts}
	// Experimental.
	Scripts *TfLifecyclePolicy_ScriptsProperty `field:"optional" json:"scripts" yaml:"scripts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dlm_lifecycle_policy#times TfLifecyclePolicy#times}.
	// Experimental.
	Times *[]*string `field:"optional" json:"times" yaml:"times"`
}

