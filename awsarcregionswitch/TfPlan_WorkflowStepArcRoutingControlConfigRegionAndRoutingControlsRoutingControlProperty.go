package awsarcregionswitch


// Experimental.
type TfPlan_WorkflowStepArcRoutingControlConfigRegionAndRoutingControlsRoutingControlProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#routing_control_arn TfPlan#routing_control_arn}.
	// Experimental.
	RoutingControlArn *string `field:"required" json:"routingControlArn" yaml:"routingControlArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#state TfPlan#state}.
	// Experimental.
	State *string `field:"required" json:"state" yaml:"state"`
}

