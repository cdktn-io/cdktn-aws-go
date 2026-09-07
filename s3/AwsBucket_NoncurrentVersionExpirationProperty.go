package s3


// Experimental.
type AwsBucket_NoncurrentVersionExpirationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#days AwsBucket#days}.
	// Experimental.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

