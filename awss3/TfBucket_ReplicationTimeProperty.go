package awss3


// Experimental.
type TfBucket_ReplicationTimeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#minutes TfBucket#minutes}.
	// Experimental.
	Minutes *float64 `field:"optional" json:"minutes" yaml:"minutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#status TfBucket#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

