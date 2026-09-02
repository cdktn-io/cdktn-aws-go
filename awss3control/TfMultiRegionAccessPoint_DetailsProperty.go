package awss3control


// Experimental.
type TfMultiRegionAccessPoint_DetailsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point#name TfMultiRegionAccessPoint#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// region block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point#region TfMultiRegionAccessPoint#region}
	// Experimental.
	Region interface{} `field:"required" json:"region" yaml:"region"`
	// public_access_block block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point#public_access_block TfMultiRegionAccessPoint#public_access_block}
	// Experimental.
	PublicAccessBlock *TfMultiRegionAccessPoint_PublicAccessBlockProperty `field:"optional" json:"publicAccessBlock" yaml:"publicAccessBlock"`
}

