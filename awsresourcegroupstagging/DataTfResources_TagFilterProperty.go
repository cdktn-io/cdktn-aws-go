package awsresourcegroupstagging


// Experimental.
type DataTfResources_TagFilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/resourcegroupstaggingapi_resources#key DataTfResources#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/resourcegroupstaggingapi_resources#values DataTfResources#values}.
	// Experimental.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

