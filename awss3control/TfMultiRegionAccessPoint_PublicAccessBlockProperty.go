package awss3control


// Experimental.
type TfMultiRegionAccessPoint_PublicAccessBlockProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point#block_public_acls TfMultiRegionAccessPoint#block_public_acls}.
	// Experimental.
	BlockPublicAcls interface{} `field:"optional" json:"blockPublicAcls" yaml:"blockPublicAcls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point#block_public_policy TfMultiRegionAccessPoint#block_public_policy}.
	// Experimental.
	BlockPublicPolicy interface{} `field:"optional" json:"blockPublicPolicy" yaml:"blockPublicPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point#ignore_public_acls TfMultiRegionAccessPoint#ignore_public_acls}.
	// Experimental.
	IgnorePublicAcls interface{} `field:"optional" json:"ignorePublicAcls" yaml:"ignorePublicAcls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point#restrict_public_buckets TfMultiRegionAccessPoint#restrict_public_buckets}.
	// Experimental.
	RestrictPublicBuckets interface{} `field:"optional" json:"restrictPublicBuckets" yaml:"restrictPublicBuckets"`
}

