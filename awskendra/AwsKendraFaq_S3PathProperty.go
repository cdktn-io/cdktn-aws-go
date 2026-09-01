package awskendra


// Experimental.
type AwsKendraFaq_S3PathProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_faq#bucket AwsKendraFaq#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_faq#key AwsKendraFaq#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}

