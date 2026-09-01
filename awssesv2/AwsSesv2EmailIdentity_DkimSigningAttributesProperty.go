package awssesv2


// Experimental.
type AwsSesv2EmailIdentity_DkimSigningAttributesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_email_identity#domain_signing_private_key AwsSesv2EmailIdentity#domain_signing_private_key}.
	// Experimental.
	DomainSigningPrivateKey *string `field:"optional" json:"domainSigningPrivateKey" yaml:"domainSigningPrivateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_email_identity#domain_signing_selector AwsSesv2EmailIdentity#domain_signing_selector}.
	// Experimental.
	DomainSigningSelector *string `field:"optional" json:"domainSigningSelector" yaml:"domainSigningSelector"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_email_identity#next_signing_key_length AwsSesv2EmailIdentity#next_signing_key_length}.
	// Experimental.
	NextSigningKeyLength *string `field:"optional" json:"nextSigningKeyLength" yaml:"nextSigningKeyLength"`
}

