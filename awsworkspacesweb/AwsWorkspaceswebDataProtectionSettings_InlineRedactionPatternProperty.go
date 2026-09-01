package awsworkspacesweb


// Experimental.
type AwsWorkspaceswebDataProtectionSettings_InlineRedactionPatternProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#built_in_pattern_id AwsWorkspaceswebDataProtectionSettings#built_in_pattern_id}.
	// Experimental.
	BuiltInPatternId *string `field:"optional" json:"builtInPatternId" yaml:"builtInPatternId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#confidence_level AwsWorkspaceswebDataProtectionSettings#confidence_level}.
	// Experimental.
	ConfidenceLevel *float64 `field:"optional" json:"confidenceLevel" yaml:"confidenceLevel"`
	// custom_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#custom_pattern AwsWorkspaceswebDataProtectionSettings#custom_pattern}
	// Experimental.
	CustomPattern interface{} `field:"optional" json:"customPattern" yaml:"customPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#enforced_urls AwsWorkspaceswebDataProtectionSettings#enforced_urls}.
	// Experimental.
	EnforcedUrls *[]*string `field:"optional" json:"enforcedUrls" yaml:"enforcedUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#exempt_urls AwsWorkspaceswebDataProtectionSettings#exempt_urls}.
	// Experimental.
	ExemptUrls *[]*string `field:"optional" json:"exemptUrls" yaml:"exemptUrls"`
	// redaction_place_holder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#redaction_place_holder AwsWorkspaceswebDataProtectionSettings#redaction_place_holder}
	// Experimental.
	RedactionPlaceHolder interface{} `field:"optional" json:"redactionPlaceHolder" yaml:"redactionPlaceHolder"`
}

