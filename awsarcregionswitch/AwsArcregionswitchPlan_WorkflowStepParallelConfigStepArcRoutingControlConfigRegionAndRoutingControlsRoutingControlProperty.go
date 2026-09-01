package awsarcregionswitch


// Experimental.
type AwsArcregionswitchPlan_WorkflowStepParallelConfigStepArcRoutingControlConfigRegionAndRoutingControlsRoutingControlProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#routing_control_arn AwsArcregionswitchPlan#routing_control_arn}.
	// Experimental.
	RoutingControlArn *string `field:"required" json:"routingControlArn" yaml:"routingControlArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#state AwsArcregionswitchPlan#state}.
	// Experimental.
	State *string `field:"required" json:"state" yaml:"state"`
}

