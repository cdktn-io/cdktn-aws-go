package awsrecyclebin


// Experimental.
type AwsRbinRule_ExcludeResourceTagsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rbin_rule#resource_tag_key AwsRbinRule#resource_tag_key}.
	// Experimental.
	ResourceTagKey *string `field:"required" json:"resourceTagKey" yaml:"resourceTagKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rbin_rule#resource_tag_value AwsRbinRule#resource_tag_value}.
	// Experimental.
	ResourceTagValue *string `field:"optional" json:"resourceTagValue" yaml:"resourceTagValue"`
}

