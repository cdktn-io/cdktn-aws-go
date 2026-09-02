package awsbedrockagents


// Experimental.
type TfDataSource_ParsingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#parsing_strategy TfDataSource#parsing_strategy}.
	// Experimental.
	ParsingStrategy *string `field:"required" json:"parsingStrategy" yaml:"parsingStrategy"`
	// bedrock_data_automation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#bedrock_data_automation_configuration TfDataSource#bedrock_data_automation_configuration}
	// Experimental.
	BedrockDataAutomationConfiguration interface{} `field:"optional" json:"bedrockDataAutomationConfiguration" yaml:"bedrockDataAutomationConfiguration"`
	// bedrock_foundation_model_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#bedrock_foundation_model_configuration TfDataSource#bedrock_foundation_model_configuration}
	// Experimental.
	BedrockFoundationModelConfiguration interface{} `field:"optional" json:"bedrockFoundationModelConfiguration" yaml:"bedrockFoundationModelConfiguration"`
}

