package awsbedrockagentcore


// Experimental.
type TfMemoryStrategy_ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#type TfMemoryStrategy#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// consolidation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#consolidation TfMemoryStrategy#consolidation}
	// Experimental.
	Consolidation interface{} `field:"optional" json:"consolidation" yaml:"consolidation"`
	// extraction block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#extraction TfMemoryStrategy#extraction}
	// Experimental.
	Extraction interface{} `field:"optional" json:"extraction" yaml:"extraction"`
	// reflection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#reflection TfMemoryStrategy#reflection}
	// Experimental.
	Reflection interface{} `field:"optional" json:"reflection" yaml:"reflection"`
	// self_managed_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagentcore_memory_strategy#self_managed_configuration TfMemoryStrategy#self_managed_configuration}
	// Experimental.
	SelfManagedConfiguration interface{} `field:"optional" json:"selfManagedConfiguration" yaml:"selfManagedConfiguration"`
}

