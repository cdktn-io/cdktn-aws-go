package arcregionswitch


// Experimental.
type AwsPlan_WorkflowProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#workflow_target_action AwsPlan#workflow_target_action}.
	// Experimental.
	WorkflowTargetAction *string `field:"required" json:"workflowTargetAction" yaml:"workflowTargetAction"`
	// step block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#step AwsPlan#step}
	// Experimental.
	Step interface{} `field:"optional" json:"step" yaml:"step"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#workflow_description AwsPlan#workflow_description}.
	// Experimental.
	WorkflowDescription *string `field:"optional" json:"workflowDescription" yaml:"workflowDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#workflow_target_region AwsPlan#workflow_target_region}.
	// Experimental.
	WorkflowTargetRegion *string `field:"optional" json:"workflowTargetRegion" yaml:"workflowTargetRegion"`
}

