package arcregionswitch


// Experimental.
type AwsPlan_WorkflowStepArcRoutingControlConfigRegionAndRoutingControlsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#region AwsPlan#region}.
	// Experimental.
	Region *string `field:"required" json:"region" yaml:"region"`
	// routing_control block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#routing_control AwsPlan#routing_control}
	// Experimental.
	RoutingControl interface{} `field:"optional" json:"routingControl" yaml:"routingControl"`
}

