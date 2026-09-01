package awsacmpca


// Experimental.
type AwsAcmpcaCertificateAuthority_CertificateAuthorityConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#key_algorithm AwsAcmpcaCertificateAuthority#key_algorithm}.
	// Experimental.
	KeyAlgorithm *string `field:"required" json:"keyAlgorithm" yaml:"keyAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#signing_algorithm AwsAcmpcaCertificateAuthority#signing_algorithm}.
	// Experimental.
	SigningAlgorithm *string `field:"required" json:"signingAlgorithm" yaml:"signingAlgorithm"`
	// subject block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#subject AwsAcmpcaCertificateAuthority#subject}
	// Experimental.
	Subject *AwsAcmpcaCertificateAuthority_SubjectProperty `field:"required" json:"subject" yaml:"subject"`
}

