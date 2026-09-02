package awsbedrock


// Experimental.
type TfGuardrail_SensitiveInformationPolicyConfigProperty struct {
	// pii_entities_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#pii_entities_config TfGuardrail#pii_entities_config}
	// Experimental.
	PiiEntitiesConfig interface{} `field:"optional" json:"piiEntitiesConfig" yaml:"piiEntitiesConfig"`
	// regexes_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#regexes_config TfGuardrail#regexes_config}
	// Experimental.
	RegexesConfig interface{} `field:"optional" json:"regexesConfig" yaml:"regexesConfig"`
}

