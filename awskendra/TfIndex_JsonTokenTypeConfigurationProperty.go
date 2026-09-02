package awskendra


// Experimental.
type TfIndex_JsonTokenTypeConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#group_attribute_field TfIndex#group_attribute_field}.
	// Experimental.
	GroupAttributeField *string `field:"required" json:"groupAttributeField" yaml:"groupAttributeField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_index#user_name_attribute_field TfIndex#user_name_attribute_field}.
	// Experimental.
	UserNameAttributeField *string `field:"required" json:"userNameAttributeField" yaml:"userNameAttributeField"`
}

