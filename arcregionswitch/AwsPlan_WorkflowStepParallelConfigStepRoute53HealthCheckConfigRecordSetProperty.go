package arcregionswitch


// Experimental.
type AwsPlan_WorkflowStepParallelConfigStepRoute53HealthCheckConfigRecordSetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#record_set_identifier AwsPlan#record_set_identifier}.
	// Experimental.
	RecordSetIdentifier *string `field:"required" json:"recordSetIdentifier" yaml:"recordSetIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#region AwsPlan#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
}

