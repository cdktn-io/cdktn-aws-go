package awsbedrockagentcore


// Experimental.
type TfMemoryStrategy_MetadataSchemaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#key TfMemoryStrategy#key}.
	// Experimental.
	Key *string `field:"required" json:"key" yaml:"key"`
	// extraction_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#extraction_config TfMemoryStrategy#extraction_config}
	// Experimental.
	ExtractionConfig interface{} `field:"optional" json:"extractionConfig" yaml:"extractionConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#extraction_type TfMemoryStrategy#extraction_type}.
	// Experimental.
	ExtractionType *string `field:"optional" json:"extractionType" yaml:"extractionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#type TfMemoryStrategy#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

