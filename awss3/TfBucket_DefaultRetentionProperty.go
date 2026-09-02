package awss3


// Experimental.
type TfBucket_DefaultRetentionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#mode TfBucket#mode}.
	// Experimental.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#days TfBucket#days}.
	// Experimental.
	Days *float64 `field:"optional" json:"days" yaml:"days"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_bucket#years TfBucket#years}.
	// Experimental.
	Years *float64 `field:"optional" json:"years" yaml:"years"`
}

