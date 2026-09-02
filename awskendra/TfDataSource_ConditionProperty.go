package awskendra


// Experimental.
type TfDataSource_ConditionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#condition_document_attribute_key TfDataSource#condition_document_attribute_key}.
	// Experimental.
	ConditionDocumentAttributeKey *string `field:"required" json:"conditionDocumentAttributeKey" yaml:"conditionDocumentAttributeKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#operator TfDataSource#operator}.
	// Experimental.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// condition_on_value block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#condition_on_value TfDataSource#condition_on_value}
	// Experimental.
	ConditionOnValue *TfDataSource_CustomDocumentEnrichmentConfigurationInlineConfigurationsConditionConditionOnValueProperty `field:"optional" json:"conditionOnValue" yaml:"conditionOnValue"`
}

