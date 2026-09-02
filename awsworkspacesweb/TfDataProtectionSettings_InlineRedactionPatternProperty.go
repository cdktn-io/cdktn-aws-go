package awsworkspacesweb


// Experimental.
type TfDataProtectionSettings_InlineRedactionPatternProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#built_in_pattern_id TfDataProtectionSettings#built_in_pattern_id}.
	// Experimental.
	BuiltInPatternId *string `field:"optional" json:"builtInPatternId" yaml:"builtInPatternId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#confidence_level TfDataProtectionSettings#confidence_level}.
	// Experimental.
	ConfidenceLevel *float64 `field:"optional" json:"confidenceLevel" yaml:"confidenceLevel"`
	// custom_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#custom_pattern TfDataProtectionSettings#custom_pattern}
	// Experimental.
	CustomPattern interface{} `field:"optional" json:"customPattern" yaml:"customPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#enforced_urls TfDataProtectionSettings#enforced_urls}.
	// Experimental.
	EnforcedUrls *[]*string `field:"optional" json:"enforcedUrls" yaml:"enforcedUrls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#exempt_urls TfDataProtectionSettings#exempt_urls}.
	// Experimental.
	ExemptUrls *[]*string `field:"optional" json:"exemptUrls" yaml:"exemptUrls"`
	// redaction_place_holder block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#redaction_place_holder TfDataProtectionSettings#redaction_place_holder}
	// Experimental.
	RedactionPlaceHolder interface{} `field:"optional" json:"redactionPlaceHolder" yaml:"redactionPlaceHolder"`
}

