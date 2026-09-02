package awsarcregionswitch


// Experimental.
type TfPlan_WorkflowStepParallelConfigStepArcRoutingControlConfigRegionAndRoutingControlsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#region TfPlan#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// routing_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#routing_control TfPlan#routing_control}
	// Experimental.
	RoutingControl interface{} `field:"optional" json:"routingControl" yaml:"routingControl"`
}

