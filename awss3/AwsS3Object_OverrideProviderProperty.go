package awss3


// Experimental.
type AwsS3Object_OverrideProviderProperty struct {
	// default_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object#default_tags AwsS3Object#default_tags}
	// Experimental.
	DefaultTags *AwsS3Object_DefaultTagsProperty `field:"optional" json:"defaultTags" yaml:"defaultTags"`
}

