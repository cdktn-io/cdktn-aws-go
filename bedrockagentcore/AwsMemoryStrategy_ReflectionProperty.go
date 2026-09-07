package bedrockagentcore


// Experimental.
type AwsMemoryStrategy_ReflectionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#append_to_prompt AwsMemoryStrategy#append_to_prompt}.
	// Experimental.
	AppendToPrompt *string `field:"required" json:"appendToPrompt" yaml:"appendToPrompt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#model_id AwsMemoryStrategy#model_id}.
	// Experimental.
	ModelId *string `field:"required" json:"modelId" yaml:"modelId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#namespace_templates AwsMemoryStrategy#namespace_templates}.
	// Experimental.
	NamespaceTemplates *[]*string `field:"required" json:"namespaceTemplates" yaml:"namespaceTemplates"`
}

