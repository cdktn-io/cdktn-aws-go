package awsiotcore


// Experimental.
type TfThingGroup_PropertiesProperty struct {
	// attribute_payload block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_thing_group#attribute_payload TfThingGroup#attribute_payload}
	// Experimental.
	AttributePayload *TfThingGroup_AttributePayloadProperty `field:"optional" json:"attributePayload" yaml:"attributePayload"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/iot_thing_group#description TfThingGroup#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
}

