package awsarcregionswitch


// Experimental.
type AwsArcregionswitchPlan_ConditionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#associated_alarm_name AwsArcregionswitchPlan#associated_alarm_name}.
	// Experimental.
	AssociatedAlarmName *string `field:"required" json:"associatedAlarmName" yaml:"associatedAlarmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#condition AwsArcregionswitchPlan#condition}.
	// Experimental.
	Condition *string `field:"required" json:"condition" yaml:"condition"`
}

