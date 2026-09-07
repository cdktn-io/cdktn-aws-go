package acmpca


// Experimental.
type AwsCertificateAuthority_CertificateAuthorityConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#key_algorithm AwsCertificateAuthority#key_algorithm}.
	// Experimental.
	KeyAlgorithm *string `field:"required" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#signing_algorithm AwsCertificateAuthority#signing_algorithm}.
	// Experimental.
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#subject AwsCertificateAuthority#subject}
	// Experimental.
	Subject *AwsCertificateAuthority_SubjectProperty `field:"required" json:"subject" yaml:"subject"`
}

