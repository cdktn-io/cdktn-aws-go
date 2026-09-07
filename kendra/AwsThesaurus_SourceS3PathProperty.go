package kendra


// Experimental.
type AwsThesaurus_SourceS3PathProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_thesaurus#bucket AwsThesaurus#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_thesaurus#key AwsThesaurus#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
}

