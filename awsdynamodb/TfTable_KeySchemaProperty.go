package awsdynamodb


// Experimental.
type TfTable_KeySchemaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#attribute_name TfTable#attribute_name}.
	// Experimental.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_table#key_type TfTable#key_type}.
	// Experimental.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
}

