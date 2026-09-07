package kendra


// Experimental.
type AwsDataSource_CustomDocumentEnrichmentConfigurationProperty struct {
	// inline_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#inline_configurations AwsDataSource#inline_configurations}
	// Experimental.
	InlineConfigurations interface{} `field:"optional" json:"inlineConfigurations" yaml:"inlineConfigurations"`
	// post_extraction_hook_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#post_extraction_hook_configuration AwsDataSource#post_extraction_hook_configuration}
	// Experimental.
	PostExtractionHookConfiguration *AwsDataSource_PostExtractionHookConfigurationProperty `field:"optional" json:"postExtractionHookConfiguration" yaml:"postExtractionHookConfiguration"`
	// pre_extraction_hook_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#pre_extraction_hook_configuration AwsDataSource#pre_extraction_hook_configuration}
	// Experimental.
	PreExtractionHookConfiguration *AwsDataSource_PreExtractionHookConfigurationProperty `field:"optional" json:"preExtractionHookConfiguration" yaml:"preExtractionHookConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/kendra_data_source#role_arn AwsDataSource#role_arn}.
	// Experimental.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

