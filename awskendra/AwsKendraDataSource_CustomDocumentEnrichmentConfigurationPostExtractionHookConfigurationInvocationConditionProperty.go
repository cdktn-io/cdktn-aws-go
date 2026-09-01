package awskendra


// Experimental.
type AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#condition_document_attribute_key AwsKendraDataSource#condition_document_attribute_key}.
	// Experimental.
	ConditionDocumentAttributeKey *string `field:"required" json:"conditionDocumentAttributeKey" yaml:"conditionDocumentAttributeKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#operator AwsKendraDataSource#operator}.
	// Experimental.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// condition_on_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#condition_on_value AwsKendraDataSource#condition_on_value}
	// Experimental.
	ConditionOnValue *AwsKendraDataSource_CustomDocumentEnrichmentConfigurationPostExtractionHookConfigurationInvocationConditionConditionOnValueProperty `field:"optional" json:"conditionOnValue" yaml:"conditionOnValue"`
}

