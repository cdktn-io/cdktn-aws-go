package awsarcregionswitch


// Experimental.
type TfPlan_WorkflowStepParallelConfigStepEksResourceScalingConfigScalingResourcesResourcesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#name TfPlan#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#namespace TfPlan#namespace}.
	// Experimental.
	Namespace *string `field:"required" json:"namespace" yaml:"namespace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#resource_name TfPlan#resource_name}.
	// Experimental.
	ResourceName *string `field:"required" json:"resourceName" yaml:"resourceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/arcregionswitch_plan#hpa_name TfPlan#hpa_name}.
	// Experimental.
	HpaName *string `field:"optional" json:"hpaName" yaml:"hpaName"`
}

