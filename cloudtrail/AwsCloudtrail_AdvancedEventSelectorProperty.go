package cloudtrail


// Experimental.
type AwsCloudtrail_AdvancedEventSelectorProperty struct {
	// field_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail#field_selector AwsCloudtrail#field_selector}
	// Experimental.
	FieldSelector interface{} `field:"required" json:"fieldSelector" yaml:"fieldSelector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail#name AwsCloudtrail#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

