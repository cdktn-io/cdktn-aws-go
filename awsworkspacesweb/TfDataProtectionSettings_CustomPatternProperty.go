package awsworkspacesweb


// Experimental.
type TfDataProtectionSettings_CustomPatternProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#pattern_name TfDataProtectionSettings#pattern_name}.
	// Experimental.
	PatternName *string `field:"required" json:"patternName" yaml:"patternName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#pattern_regex TfDataProtectionSettings#pattern_regex}.
	// Experimental.
	PatternRegex *string `field:"required" json:"patternRegex" yaml:"patternRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#keyword_regex TfDataProtectionSettings#keyword_regex}.
	// Experimental.
	KeywordRegex *string `field:"optional" json:"keywordRegex" yaml:"keywordRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#pattern_description TfDataProtectionSettings#pattern_description}.
	// Experimental.
	PatternDescription *string `field:"optional" json:"patternDescription" yaml:"patternDescription"`
}

