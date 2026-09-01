package awsiotcore


// Experimental.
type AwsIotThingGroup_PropertiesProperty struct {
	// attribute_payload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_thing_group#attribute_payload AwsIotThingGroup#attribute_payload}
	// Experimental.
	AttributePayload *AwsIotThingGroup_AttributePayloadProperty `field:"optional" json:"attributePayload" yaml:"attributePayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_thing_group#description AwsIotThingGroup#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

