package securitylake


// Experimental.
type AwsCustomLogSource_ProviderIdentityProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_custom_log_source#external_id AwsCustomLogSource#external_id}.
	// Experimental.
	ExternalId *string `field:"required" json:"externalId" yaml:"externalId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_custom_log_source#principal AwsCustomLogSource#principal}.
	// Experimental.
	Principal *string `field:"required" json:"principal" yaml:"principal"`
}

