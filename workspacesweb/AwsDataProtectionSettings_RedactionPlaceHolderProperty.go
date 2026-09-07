package workspacesweb


// Experimental.
type AwsDataProtectionSettings_RedactionPlaceHolderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#redaction_place_holder_type AwsDataProtectionSettings#redaction_place_holder_type}.
	// Experimental.
	RedactionPlaceHolderType *string `field:"required" json:"redactionPlaceHolderType" yaml:"redactionPlaceHolderType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_data_protection_settings#redaction_place_holder_text AwsDataProtectionSettings#redaction_place_holder_text}.
	// Experimental.
	RedactionPlaceHolderText *string `field:"optional" json:"redactionPlaceHolderText" yaml:"redactionPlaceHolderText"`
}

