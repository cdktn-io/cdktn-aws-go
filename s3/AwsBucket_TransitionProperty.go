package s3


// Experimental.
type AwsBucket_TransitionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#storage_class AwsBucket#storage_class}.
	// Experimental.
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#date AwsBucket#date}.
	// Experimental.
	Date *string `field:"optional" json:"date" yaml:"date"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#days AwsBucket#days}.
	// Experimental.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
}

