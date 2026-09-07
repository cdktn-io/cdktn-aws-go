package arcregionswitch


// Experimental.
type AwsPlan_WorkflowStepEksResourceScalingConfigScalingResourcesResourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#name AwsPlan#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#namespace AwsPlan#namespace}.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#resource_name AwsPlan#resource_name}.
	// Experimental.
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#hpa_name AwsPlan#hpa_name}.
	// Experimental.
	HpaName *string `field:"optional" json:"hpaName" yaml:"hpaName"`
}

