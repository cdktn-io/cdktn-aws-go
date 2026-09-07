package kendra


// Experimental.
type AwsDataSource_CustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#condition_document_attribute_key AwsDataSource#condition_document_attribute_key}.
	// Experimental.
	ConditionDocumentAttributeKey *string `field:"required" json:"conditionDocumentAttributeKey" yaml:"conditionDocumentAttributeKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#operator AwsDataSource#operator}.
	// Experimental.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// condition_on_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#condition_on_value AwsDataSource#condition_on_value}
	// Experimental.
	ConditionOnValue *AwsDataSource_CustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueProperty `field:"optional" json:"conditionOnValue" yaml:"conditionOnValue"`
}

