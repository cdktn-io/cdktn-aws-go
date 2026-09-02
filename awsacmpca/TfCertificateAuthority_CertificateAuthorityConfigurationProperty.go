package awsacmpca


// Experimental.
type TfCertificateAuthority_CertificateAuthorityConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#key_algorithm TfCertificateAuthority#key_algorithm}.
	// Experimental.
	KeyAlgorithm *string `field:"required" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#signing_algorithm TfCertificateAuthority#signing_algorithm}.
	// Experimental.
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#subject TfCertificateAuthority#subject}
	// Experimental.
	Subject *TfCertificateAuthority_SubjectProperty `field:"required" json:"subject" yaml:"subject"`
}

