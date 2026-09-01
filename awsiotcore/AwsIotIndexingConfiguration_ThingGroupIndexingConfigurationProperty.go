package awsiotcore


// Experimental.
type AwsIotIndexingConfiguration_ThingGroupIndexingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_indexing_configuration#thing_group_indexing_mode AwsIotIndexingConfiguration#thing_group_indexing_mode}.
	// Experimental.
	ThingGroupIndexingMode *string `field:"required" json:"thingGroupIndexingMode" yaml:"thingGroupIndexingMode"`
	// custom_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_indexing_configuration#custom_field AwsIotIndexingConfiguration#custom_field}
	// Experimental.
	CustomField interface{} `field:"optional" json:"customField" yaml:"customField"`
	// managed_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_indexing_configuration#managed_field AwsIotIndexingConfiguration#managed_field}
	// Experimental.
	ManagedField interface{} `field:"optional" json:"managedField" yaml:"managedField"`
}

