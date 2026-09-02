package awscodepipeline


// Experimental.
type TfWebhook_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_webhook#json_path TfWebhook#json_path}.
	// Experimental.
	JsonPath *string `field:"required" json:"jsonPath" yaml:"jsonPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codepipeline_webhook#match_equals TfWebhook#match_equals}.
	// Experimental.
	MatchEquals *string `field:"required" json:"matchEquals" yaml:"matchEquals"`
}

