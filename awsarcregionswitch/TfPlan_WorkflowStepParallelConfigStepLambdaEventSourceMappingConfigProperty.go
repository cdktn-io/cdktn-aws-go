package awsarcregionswitch


// Experimental.
type TfPlan_WorkflowStepParallelConfigStepLambdaEventSourceMappingConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#action TfPlan#action}.
	// Experimental.
	Action *string `field:"required" json:"action" yaml:"action"`
	// region_event_source_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#region_event_source_mapping TfPlan#region_event_source_mapping}
	// Experimental.
	RegionEventSourceMapping interface{} `field:"optional" json:"regionEventSourceMapping" yaml:"regionEventSourceMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#timeout_minutes TfPlan#timeout_minutes}.
	// Experimental.
	TimeoutMinutes *float64 `field:"optional" json:"timeoutMinutes" yaml:"timeoutMinutes"`
	// ungraceful block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#ungraceful TfPlan#ungraceful}
	// Experimental.
	Ungraceful interface{} `field:"optional" json:"ungraceful" yaml:"ungraceful"`
}

