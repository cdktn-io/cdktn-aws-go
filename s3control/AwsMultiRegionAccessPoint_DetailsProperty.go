package s3control


// Experimental.
type AwsMultiRegionAccessPoint_DetailsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point#name AwsMultiRegionAccessPoint#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point#region AwsMultiRegionAccessPoint#region}
	// Experimental.
	Region interface{} `field:"required" json:"region" yaml:"region"`
	// public_access_block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point#public_access_block AwsMultiRegionAccessPoint#public_access_block}
	// Experimental.
	PublicAccessBlock *AwsMultiRegionAccessPoint_PublicAccessBlockProperty `field:"optional" json:"publicAccessBlock" yaml:"publicAccessBlock"`
}

