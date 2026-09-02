package awss3


// Experimental.
type TfObjectCopy_OverrideProviderProperty struct {
	// default_tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3_object_copy#default_tags TfObjectCopy#default_tags}
	// Experimental.
	DefaultTags *TfObjectCopy_DefaultTagsProperty `field:"optional" json:"defaultTags" yaml:"defaultTags"`
}

