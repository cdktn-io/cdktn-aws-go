package bedrock


// Experimental.
type AwsGuardrail_WordPolicyConfigProperty struct {
	// managed_word_lists_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#managed_word_lists_config AwsGuardrail#managed_word_lists_config}
	// Experimental.
	ManagedWordListsConfig interface{} `field:"optional" json:"managedWordListsConfig" yaml:"managedWordListsConfig"`
	// words_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bedrock_guardrail#words_config AwsGuardrail#words_config}
	// Experimental.
	WordsConfig interface{} `field:"optional" json:"wordsConfig" yaml:"wordsConfig"`
}

