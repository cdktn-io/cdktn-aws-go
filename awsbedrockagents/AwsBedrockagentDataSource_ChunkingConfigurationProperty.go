package awsbedrockagents


// Experimental.
type AwsBedrockagentDataSource_ChunkingConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#chunking_strategy AwsBedrockagentDataSource#chunking_strategy}.
	// Experimental.
	ChunkingStrategy *string `field:"required" json:"chunkingStrategy" yaml:"chunkingStrategy"`
	// fixed_size_chunking_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#fixed_size_chunking_configuration AwsBedrockagentDataSource#fixed_size_chunking_configuration}
	// Experimental.
	FixedSizeChunkingConfiguration interface{} `field:"optional" json:"fixedSizeChunkingConfiguration" yaml:"fixedSizeChunkingConfiguration"`
	// hierarchical_chunking_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#hierarchical_chunking_configuration AwsBedrockagentDataSource#hierarchical_chunking_configuration}
	// Experimental.
	HierarchicalChunkingConfiguration interface{} `field:"optional" json:"hierarchicalChunkingConfiguration" yaml:"hierarchicalChunkingConfiguration"`
	// semantic_chunking_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_data_source#semantic_chunking_configuration AwsBedrockagentDataSource#semantic_chunking_configuration}
	// Experimental.
	SemanticChunkingConfiguration interface{} `field:"optional" json:"semanticChunkingConfiguration" yaml:"semanticChunkingConfiguration"`
}

