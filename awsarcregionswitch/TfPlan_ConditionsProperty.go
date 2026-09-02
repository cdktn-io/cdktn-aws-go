package awsarcregionswitch


// Experimental.
type TfPlan_ConditionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#associated_alarm_name TfPlan#associated_alarm_name}.
	// Experimental.
	AssociatedAlarmName *string `field:"required" json:"associatedAlarmName" yaml:"associatedAlarmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#condition TfPlan#condition}.
	// Experimental.
	Condition *string `field:"required" json:"condition" yaml:"condition"`
}

