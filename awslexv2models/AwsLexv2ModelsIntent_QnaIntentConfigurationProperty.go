package awslexv2models


// Experimental.
type AwsLexv2ModelsIntent_QnaIntentConfigurationProperty struct {
	// bedrock_model_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#bedrock_model_configuration AwsLexv2ModelsIntent#bedrock_model_configuration}
	// Experimental.
	BedrockModelConfiguration interface{} `field:"optional" json:"bedrockModelConfiguration" yaml:"bedrockModelConfiguration"`
	// data_source_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lexv2models_intent#data_source_configuration AwsLexv2ModelsIntent#data_source_configuration}
	// Experimental.
	DataSourceConfiguration interface{} `field:"optional" json:"dataSourceConfiguration" yaml:"dataSourceConfiguration"`
}

