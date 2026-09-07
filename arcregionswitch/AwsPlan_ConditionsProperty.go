package arcregionswitch


// Experimental.
type AwsPlan_ConditionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#associated_alarm_name AwsPlan#associated_alarm_name}.
	// Experimental.
	AssociatedAlarmName *string `field:"required" json:"associatedAlarmName" yaml:"associatedAlarmName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#condition AwsPlan#condition}.
	// Experimental.
	Condition *string `field:"required" json:"condition" yaml:"condition"`
}

