package awskendra


// Experimental.
type TfFaq_S3PathProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_faq#bucket TfFaq#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_faq#key TfFaq#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}

