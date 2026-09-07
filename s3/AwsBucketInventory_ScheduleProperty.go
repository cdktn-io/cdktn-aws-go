package s3


// Experimental.
type AwsBucketInventory_ScheduleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket_inventory#frequency AwsBucketInventory#frequency}.
	// Experimental.
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
}

