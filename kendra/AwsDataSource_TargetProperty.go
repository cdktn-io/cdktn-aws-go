package kendra


// Experimental.
type AwsDataSource_TargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#target_document_attribute_key AwsDataSource#target_document_attribute_key}.
	// Experimental.
	TargetDocumentAttributeKey *string `field:"optional" json:"targetDocumentAttributeKey" yaml:"targetDocumentAttributeKey"`
	// target_document_attribute_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#target_document_attribute_value AwsDataSource#target_document_attribute_value}
	// Experimental.
	TargetDocumentAttributeValue *AwsDataSource_TargetDocumentAttributeValueProperty `field:"optional" json:"targetDocumentAttributeValue" yaml:"targetDocumentAttributeValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#target_document_attribute_value_deletion AwsDataSource#target_document_attribute_value_deletion}.
	// Experimental.
	TargetDocumentAttributeValueDeletion interface{} `field:"optional" json:"targetDocumentAttributeValueDeletion" yaml:"targetDocumentAttributeValueDeletion"`
}

