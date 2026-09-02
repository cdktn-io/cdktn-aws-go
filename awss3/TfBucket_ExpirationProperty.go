package awss3


// Experimental.
type TfBucket_ExpirationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#date TfBucket#date}.
	// Experimental.
	Date *string `field:"optional" json:"date" yaml:"date"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#days TfBucket#days}.
	// Experimental.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#expired_object_delete_marker TfBucket#expired_object_delete_marker}.
	// Experimental.
	ExpiredObjectDeleteMarker interface{} `field:"optional" json:"expiredObjectDeleteMarker" yaml:"expiredObjectDeleteMarker"`
}

