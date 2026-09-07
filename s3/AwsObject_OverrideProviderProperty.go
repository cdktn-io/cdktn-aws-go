package s3


// Experimental.
type AwsObject_OverrideProviderProperty struct {
	// default_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object#default_tags AwsObject#default_tags}
	// Experimental.
	DefaultTags *AwsObject_DefaultTagsProperty `field:"optional" json:"defaultTags" yaml:"defaultTags"`
}

