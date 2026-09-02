package awsresourcegroups


// Experimental.
type TfGroup_ResourceQueryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resourcegroups_group#query TfGroup#query}.
	// Experimental.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resourcegroups_group#type TfGroup#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

