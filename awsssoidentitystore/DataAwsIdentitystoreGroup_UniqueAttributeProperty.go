package awsssoidentitystore


// Experimental.
type DataAwsIdentitystoreGroup_UniqueAttributeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_group#attribute_path DataAwsIdentitystoreGroup#attribute_path}.
	// Experimental.
	AttributePath *string `field:"required" json:"attributePath" yaml:"attributePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_group#attribute_value DataAwsIdentitystoreGroup#attribute_value}.
	// Experimental.
	AttributeValue *string `field:"required" json:"attributeValue" yaml:"attributeValue"`
}

