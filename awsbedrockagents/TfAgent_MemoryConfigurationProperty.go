package awsbedrockagents


// Experimental.
type TfAgent_MemoryConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#enabled_memory_types TfAgent#enabled_memory_types}.
	// Experimental.
	EnabledMemoryTypes *[]*string `field:"optional" json:"enabledMemoryTypes" yaml:"enabledMemoryTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#session_summary_configuration TfAgent#session_summary_configuration}.
	// Experimental.
	SessionSummaryConfiguration interface{} `field:"optional" json:"sessionSummaryConfiguration" yaml:"sessionSummaryConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrockagent_agent#storage_days TfAgent#storage_days}.
	// Experimental.
	StorageDays *float64 `field:"optional" json:"storageDays" yaml:"storageDays"`
}

