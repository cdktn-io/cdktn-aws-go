package awsresourcegroups


// Experimental.
type TfGroup_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resourcegroups_group#create TfGroup#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resourcegroups_group#update TfGroup#update}.
	// Experimental.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

