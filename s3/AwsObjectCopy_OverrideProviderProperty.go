package s3


// Experimental.
type AwsObjectCopy_OverrideProviderProperty struct {
	// default_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#default_tags AwsObjectCopy#default_tags}
	// Experimental.
	DefaultTags *AwsObjectCopy_DefaultTagsProperty `field:"optional" json:"defaultTags" yaml:"defaultTags"`
}

