package awsiotcore


// Experimental.
type AwsIotIndexingConfiguration_ThingIndexingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_indexing_configuration#thing_indexing_mode AwsIotIndexingConfiguration#thing_indexing_mode}.
	// Experimental.
	ThingIndexingMode *string `field:"required" json:"thingIndexingMode" yaml:"thingIndexingMode"`
	// custom_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_indexing_configuration#custom_field AwsIotIndexingConfiguration#custom_field}
	// Experimental.
	CustomField interface{} `field:"optional" json:"customField" yaml:"customField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_indexing_configuration#device_defender_indexing_mode AwsIotIndexingConfiguration#device_defender_indexing_mode}.
	// Experimental.
	DeviceDefenderIndexingMode *string `field:"optional" json:"deviceDefenderIndexingMode" yaml:"deviceDefenderIndexingMode"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_indexing_configuration#filter AwsIotIndexingConfiguration#filter}
	// Experimental.
	Filter *AwsIotIndexingConfiguration_FilterProperty `field:"optional" json:"filter" yaml:"filter"`
	// managed_field block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_indexing_configuration#managed_field AwsIotIndexingConfiguration#managed_field}
	// Experimental.
	ManagedField interface{} `field:"optional" json:"managedField" yaml:"managedField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_indexing_configuration#named_shadow_indexing_mode AwsIotIndexingConfiguration#named_shadow_indexing_mode}.
	// Experimental.
	NamedShadowIndexingMode *string `field:"optional" json:"namedShadowIndexingMode" yaml:"namedShadowIndexingMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_indexing_configuration#thing_connectivity_indexing_mode AwsIotIndexingConfiguration#thing_connectivity_indexing_mode}.
	// Experimental.
	ThingConnectivityIndexingMode *string `field:"optional" json:"thingConnectivityIndexingMode" yaml:"thingConnectivityIndexingMode"`
}

