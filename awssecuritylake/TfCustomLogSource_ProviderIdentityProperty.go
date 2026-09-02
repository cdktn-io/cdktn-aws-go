package awssecuritylake


// Experimental.
type TfCustomLogSource_ProviderIdentityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_custom_log_source#external_id TfCustomLogSource#external_id}.
	// Experimental.
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_custom_log_source#principal TfCustomLogSource#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

