package awsneptune


// Experimental.
type AwsNeptuneClusterParameterGroup_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_parameter_group#name AwsNeptuneClusterParameterGroup#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_parameter_group#value AwsNeptuneClusterParameterGroup#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster_parameter_group#apply_method AwsNeptuneClusterParameterGroup#apply_method}.
	// Experimental.
	ApplyMethod *string `field:"optional" json:"applyMethod" yaml:"applyMethod"`
}

