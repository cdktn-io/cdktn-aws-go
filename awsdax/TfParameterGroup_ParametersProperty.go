package awsdax


// Experimental.
type TfParameterGroup_ParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dax_parameter_group#name TfParameterGroup#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dax_parameter_group#value TfParameterGroup#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

