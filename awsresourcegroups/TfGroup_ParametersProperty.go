package awsresourcegroups


// Experimental.
type TfGroup_ParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resourcegroups_group#name TfGroup#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resourcegroups_group#values TfGroup#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

