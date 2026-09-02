package awsssoidentitystore


// Experimental.
type DataTfGroup_UniqueAttributeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_group#attribute_path DataTfGroup#attribute_path}.
	// Experimental.
	AttributePath *string `field:"required" json:"attributePath" yaml:"attributePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/identitystore_group#attribute_value DataTfGroup#attribute_value}.
	// Experimental.
	AttributeValue *string `field:"required" json:"attributeValue" yaml:"attributeValue"`
}

