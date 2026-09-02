package awss3control


// Experimental.
type TfMultiRegionAccessPoint_RegionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point#bucket TfMultiRegionAccessPoint#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3control_multi_region_access_point#bucket_account_id TfMultiRegionAccessPoint#bucket_account_id}.
	// Experimental.
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
}

