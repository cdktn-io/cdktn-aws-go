package awsarcregionswitch


// Experimental.
type AwsArcregionswitchPlan_TriggersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#action AwsArcregionswitchPlan#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#min_delay_minutes_between_executions AwsArcregionswitchPlan#min_delay_minutes_between_executions}.
	// Experimental.
	MinDelayMinutesBetweenExecutions *float64 `field:"required" json:"minDelayMinutesBetweenExecutions" yaml:"minDelayMinutesBetweenExecutions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#target_region AwsArcregionswitchPlan#target_region}.
	// Experimental.
	TargetRegion *string `field:"required" json:"targetRegion" yaml:"targetRegion"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#conditions AwsArcregionswitchPlan#conditions}
	// Experimental.
	Conditions interface{} `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#description AwsArcregionswitchPlan#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

