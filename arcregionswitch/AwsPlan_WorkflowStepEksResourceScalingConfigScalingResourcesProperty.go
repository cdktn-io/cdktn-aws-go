package arcregionswitch


// Experimental.
type AwsPlan_WorkflowStepEksResourceScalingConfigScalingResourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#namespace AwsPlan#namespace}.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// resources block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#resources AwsPlan#resources}
	// Experimental.
	Resources interface{} `field:"optional" json:"resources" yaml:"resources"`
}

