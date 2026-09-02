package awskendra


// Experimental.
type TfDataSource_CustomDocumentEnrichmentConfigurationProperty struct {
	// inline_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#inline_configurations TfDataSource#inline_configurations}
	// Experimental.
	InlineConfigurations interface{} `field:"optional" json:"inlineConfigurations" yaml:"inlineConfigurations"`
	// post_extraction_hook_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#post_extraction_hook_configuration TfDataSource#post_extraction_hook_configuration}
	// Experimental.
	PostExtractionHookConfiguration *TfDataSource_PostExtractionHookConfigurationProperty `field:"optional" json:"postExtractionHookConfiguration" yaml:"postExtractionHookConfiguration"`
	// pre_extraction_hook_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#pre_extraction_hook_configuration TfDataSource#pre_extraction_hook_configuration}
	// Experimental.
	PreExtractionHookConfiguration *TfDataSource_PreExtractionHookConfigurationProperty `field:"optional" json:"preExtractionHookConfiguration" yaml:"preExtractionHookConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#role_arn TfDataSource#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

