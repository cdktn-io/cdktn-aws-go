package awsarcregionswitch


// Experimental.
type TfPlan_WorkflowStepParallelConfigStepRoute53HealthCheckConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#hosted_zone_id TfPlan#hosted_zone_id}.
	// Experimental.
	HostedZoneId *string `field:"required" json:"hostedZoneId" yaml:"hostedZoneId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#record_name TfPlan#record_name}.
	// Experimental.
	RecordName *string `field:"required" json:"recordName" yaml:"recordName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#cross_account_role TfPlan#cross_account_role}.
	// Experimental.
	CrossAccountRole *string `field:"optional" json:"crossAccountRole" yaml:"crossAccountRole"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#external_id TfPlan#external_id}.
	// Experimental.
	ExternalId *string `field:"optional" json:"externalId" yaml:"externalId"`
	// record_set block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#record_set TfPlan#record_set}
	// Experimental.
	RecordSet interface{} `field:"optional" json:"recordSet" yaml:"recordSet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#timeout_minutes TfPlan#timeout_minutes}.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
}

