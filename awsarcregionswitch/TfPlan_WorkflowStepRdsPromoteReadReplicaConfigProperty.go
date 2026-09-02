package awsarcregionswitch


// Experimental.
type TfPlan_WorkflowStepRdsPromoteReadReplicaConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#db_instance_arn_map TfPlan#db_instance_arn_map}.
	// Experimental.
	DbInstanceArnMap *map[string]*string `field:"required" json:"dbInstanceArnMap" yaml:"dbInstanceArnMap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#cross_account_role TfPlan#cross_account_role}.
	// Experimental.
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#external_id TfPlan#external_id}.
	// Experimental.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#timeout_minutes TfPlan#timeout_minutes}.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

