package awsglue


// Experimental.
type TfClassifier_GrokClassifierProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#classification TfClassifier#classification}.
	// Experimental.
	Classification *string `field:"required" json:"classification" yaml:"classification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#grok_pattern TfClassifier#grok_pattern}.
	// Experimental.
	GrokPattern *string `field:"required" json:"grokPattern" yaml:"grokPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_classifier#custom_patterns TfClassifier#custom_patterns}.
	// Experimental.
	CustomPatterns *string `field:"optional" json:"customPatterns" yaml:"customPatterns"`
}

