package awsresourcegroups


// Experimental.
type AwsResourcegroupsGroup_ParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resourcegroups_group#name AwsResourcegroupsGroup#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/resourcegroups_group#values AwsResourcegroupsGroup#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

