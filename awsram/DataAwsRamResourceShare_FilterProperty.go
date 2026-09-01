package awsram


// Experimental.
type DataAwsRamResourceShare_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ram_resource_share#name DataAwsRamResourceShare#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ram_resource_share#values DataAwsRamResourceShare#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

