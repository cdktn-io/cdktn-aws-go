package awsworkspacesweb


// Experimental.
type AwsWorkspaceswebDataProtectionSettings_CustomPatternProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#pattern_name AwsWorkspaceswebDataProtectionSettings#pattern_name}.
	// Experimental.
	PatternName *string `field:"required" json:"patternName" yaml:"patternName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#pattern_regex AwsWorkspaceswebDataProtectionSettings#pattern_regex}.
	// Experimental.
	PatternRegex *string `field:"required" json:"patternRegex" yaml:"patternRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#keyword_regex AwsWorkspaceswebDataProtectionSettings#keyword_regex}.
	// Experimental.
	KeywordRegex *string `field:"optional" json:"keywordRegex" yaml:"keywordRegex"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#pattern_description AwsWorkspaceswebDataProtectionSettings#pattern_description}.
	// Experimental.
	PatternDescription *string `field:"optional" json:"patternDescription" yaml:"patternDescription"`
}

