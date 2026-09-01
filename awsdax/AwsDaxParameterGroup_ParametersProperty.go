package awsdax


// Experimental.
type AwsDaxParameterGroup_ParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dax_parameter_group#name AwsDaxParameterGroup#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dax_parameter_group#value AwsDaxParameterGroup#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

