package awss3


// Experimental.
type AwsS3ObjectCopy_OverrideProviderProperty struct {
	// default_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#default_tags AwsS3ObjectCopy#default_tags}
	// Experimental.
	DefaultTags *AwsS3ObjectCopy_DefaultTagsProperty `field:"optional" json:"defaultTags" yaml:"defaultTags"`
}

