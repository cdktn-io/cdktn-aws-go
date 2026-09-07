package kendra


// Experimental.
type AwsFaq_S3PathProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_faq#bucket AwsFaq#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_faq#key AwsFaq#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}

