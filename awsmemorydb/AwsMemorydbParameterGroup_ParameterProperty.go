package awsmemorydb


// Experimental.
type AwsMemorydbParameterGroup_ParameterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/memorydb_parameter_group#name AwsMemorydbParameterGroup#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/memorydb_parameter_group#value AwsMemorydbParameterGroup#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

