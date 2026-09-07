package s3


// Experimental.
type AwsBucketInventory_DestinationProperty struct {
	// bucket block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_inventory#bucket AwsBucketInventory#bucket}
	// Experimental.
	Bucket *AwsBucketInventory_BucketProperty `field:"required" json:"bucket" yaml:"bucket"`
}

