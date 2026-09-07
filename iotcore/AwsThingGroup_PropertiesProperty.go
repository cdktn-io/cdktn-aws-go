package iotcore


// Experimental.
type AwsThingGroup_PropertiesProperty struct {
	// attribute_payload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_thing_group#attribute_payload AwsThingGroup#attribute_payload}
	// Experimental.
	AttributePayload *AwsThingGroup_AttributePayloadProperty `field:"optional" json:"attributePayload" yaml:"attributePayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_thing_group#description AwsThingGroup#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

