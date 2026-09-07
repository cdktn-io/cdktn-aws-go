package workspacesweb


// Experimental.
type AwsDataProtectionSettings_InlineRedactionConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#global_confidence_level AwsDataProtectionSettings#global_confidence_level}.
	// Experimental.
	GlobalConfidenceLevel *float64 `field:"optional" json:"globalConfidenceLevel" yaml:"globalConfidenceLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#global_enforced_urls AwsDataProtectionSettings#global_enforced_urls}.
	// Experimental.
	GlobalEnforcedUrls *[]*string `field:"optional" json:"globalEnforcedUrls" yaml:"globalEnforcedUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#global_exempt_urls AwsDataProtectionSettings#global_exempt_urls}.
	// Experimental.
	GlobalExemptUrls *[]*string `field:"optional" json:"globalExemptUrls" yaml:"globalExemptUrls"`
	// inline_redaction_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#inline_redaction_pattern AwsDataProtectionSettings#inline_redaction_pattern}
	// Experimental.
	InlineRedactionPattern interface{} `field:"optional" json:"inlineRedactionPattern" yaml:"inlineRedactionPattern"`
}

