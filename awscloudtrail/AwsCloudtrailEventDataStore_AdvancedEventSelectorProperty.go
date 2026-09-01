package awscloudtrail


// Experimental.
type AwsCloudtrailEventDataStore_AdvancedEventSelectorProperty struct {
	// field_selector block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#field_selector AwsCloudtrailEventDataStore#field_selector}
	// Experimental.
	FieldSelector interface{} `field:"optional" json:"fieldSelector" yaml:"fieldSelector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudtrail_event_data_store#name AwsCloudtrailEventDataStore#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

