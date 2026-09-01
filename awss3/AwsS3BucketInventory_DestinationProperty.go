package awss3


// Experimental.
type AwsS3BucketInventory_DestinationProperty struct {
	// bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_inventory#bucket AwsS3BucketInventory#bucket}
	// Experimental.
	Bucket *AwsS3BucketInventory_BucketProperty `field:"required" json:"bucket" yaml:"bucket"`
}

