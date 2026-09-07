package glue


// Experimental.
type AwsClassifier_GrokClassifierProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#classification AwsClassifier#classification}.
	// Experimental.
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#grok_pattern AwsClassifier#grok_pattern}.
	// Experimental.
	GrokPattern *string `field:"required" json:"grokPattern" yaml:"grokPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#custom_patterns AwsClassifier#custom_patterns}.
	// Experimental.
	CustomPatterns *string `field:"optional" json:"customPatterns" yaml:"customPatterns"`
}

