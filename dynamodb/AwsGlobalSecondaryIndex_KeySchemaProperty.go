package dynamodb


// Experimental.
type AwsGlobalSecondaryIndex_KeySchemaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_global_secondary_index#attribute_name AwsGlobalSecondaryIndex#attribute_name}.
	// Experimental.
	AttributeName *string `field:"required" json:"attributeName" yaml:"attributeName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_global_secondary_index#attribute_type AwsGlobalSecondaryIndex#attribute_type}.
	// Experimental.
	AttributeType *string `field:"required" json:"attributeType" yaml:"attributeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dynamodb_global_secondary_index#key_type AwsGlobalSecondaryIndex#key_type}.
	// Experimental.
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
}

