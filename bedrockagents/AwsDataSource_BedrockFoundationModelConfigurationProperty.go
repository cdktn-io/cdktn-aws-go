package bedrockagents


// Experimental.
type AwsDataSource_BedrockFoundationModelConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#model_arn AwsDataSource#model_arn}.
	// Experimental.
	ModelArn *string `field:"required" json:"modelArn" yaml:"modelArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#parsing_modality AwsDataSource#parsing_modality}.
	// Experimental.
	ParsingModality *string `field:"optional" json:"parsingModality" yaml:"parsingModality"`
	// parsing_prompt block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#parsing_prompt AwsDataSource#parsing_prompt}
	// Experimental.
	ParsingPrompt interface{} `field:"optional" json:"parsingPrompt" yaml:"parsingPrompt"`
}

