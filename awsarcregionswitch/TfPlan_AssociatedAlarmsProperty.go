package awsarcregionswitch


// Experimental.
type TfPlan_AssociatedAlarmsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#alarm_type TfPlan#alarm_type}.
	// Experimental.
	AlarmType *string `field:"required" json:"alarmType" yaml:"alarmType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#map_block_key TfPlan#map_block_key}.
	// Experimental.
	MapBlockKey *string `field:"required" json:"mapBlockKey" yaml:"mapBlockKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#resource_identifier TfPlan#resource_identifier}.
	// Experimental.
	ResourceIdentifier *string `field:"required" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#cross_account_role TfPlan#cross_account_role}.
	// Experimental.
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#external_id TfPlan#external_id}.
	// Experimental.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
}

