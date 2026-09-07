package bedrockagentcore


// Experimental.
type AwsMemoryStrategy_ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#type AwsMemoryStrategy#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// consolidation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#consolidation AwsMemoryStrategy#consolidation}
	// Experimental.
	Consolidation interface{} `field:"optional" json:"consolidation" yaml:"consolidation"`
	// extraction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#extraction AwsMemoryStrategy#extraction}
	// Experimental.
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
	// reflection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#reflection AwsMemoryStrategy#reflection}
	// Experimental.
	Reflection interface{} `field:"optional" json:"reflection" yaml:"reflection"`
	// self_managed_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#self_managed_configuration AwsMemoryStrategy#self_managed_configuration}
	// Experimental.
	SelfManagedConfiguration interface{} `field:"optional" json:"selfManagedConfiguration" yaml:"selfManagedConfiguration"`
}

