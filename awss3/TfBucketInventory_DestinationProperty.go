package awss3


// Experimental.
type TfBucketInventory_DestinationProperty struct {
	// bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_inventory#bucket TfBucketInventory#bucket}
	// Experimental.
	Bucket *TfBucketInventory_BucketProperty `field:"required" json:"bucket" yaml:"bucket"`
}

