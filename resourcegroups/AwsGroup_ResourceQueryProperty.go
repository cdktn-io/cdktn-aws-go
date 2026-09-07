package resourcegroups


// Experimental.
type AwsGroup_ResourceQueryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resourcegroups_group#query AwsGroup#query}.
	// Experimental.
	Query *string `field:"required" json:"query" yaml:"query"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resourcegroups_group#type AwsGroup#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

