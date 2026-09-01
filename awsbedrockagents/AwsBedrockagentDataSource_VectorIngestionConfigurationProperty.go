package awsbedrockagents


// Experimental.
type AwsBedrockagentDataSource_VectorIngestionConfigurationProperty struct {
	// chunking_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#chunking_configuration AwsBedrockagentDataSource#chunking_configuration}
	// Experimental.
	ChunkingConfiguration interface{} `field:"optional" json:"chunkingConfiguration" yaml:"chunkingConfiguration"`
	// custom_transformation_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#custom_transformation_configuration AwsBedrockagentDataSource#custom_transformation_configuration}
	// Experimental.
	CustomTransformationConfiguration interface{} `field:"optional" json:"customTransformationConfiguration" yaml:"customTransformationConfiguration"`
	// parsing_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#parsing_configuration AwsBedrockagentDataSource#parsing_configuration}
	// Experimental.
	ParsingConfiguration interface{} `field:"optional" json:"parsingConfiguration" yaml:"parsingConfiguration"`
}

