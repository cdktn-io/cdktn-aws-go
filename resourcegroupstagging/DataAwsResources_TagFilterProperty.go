package resourcegroupstagging


// Experimental.
type DataAwsResources_TagFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/resourcegroupstaggingapi_resources#key DataAwsResources#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/resourcegroupstaggingapi_resources#values DataAwsResources#values}.
	// Experimental.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

