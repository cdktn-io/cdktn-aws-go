package awsredshift


// Experimental.
type TfIdcApplication_AuthorizedTokenIssuerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_idc_application#authorized_audiences_list TfIdcApplication#authorized_audiences_list}.
	// Experimental.
	AuthorizedAudiencesList *[]*string `field:"optional" json:"authorizedAudiencesList" yaml:"authorizedAudiencesList"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_idc_application#trusted_token_issuer_arn TfIdcApplication#trusted_token_issuer_arn}.
	// Experimental.
	TrustedTokenIssuerArn *string `field:"optional" json:"trustedTokenIssuerArn" yaml:"trustedTokenIssuerArn"`
}

