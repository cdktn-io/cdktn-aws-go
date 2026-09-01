package awsacmpca


// Experimental.
type AwsAcmpcaCertificateAuthority_RevocationConfigurationProperty struct {
	// crl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#crl_configuration AwsAcmpcaCertificateAuthority#crl_configuration}
	// Experimental.
	CrlConfiguration *AwsAcmpcaCertificateAuthority_CrlConfigurationProperty `field:"optional" json:"crlConfiguration" yaml:"crlConfiguration"`
	// ocsp_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#ocsp_configuration AwsAcmpcaCertificateAuthority#ocsp_configuration}
	// Experimental.
	OcspConfiguration *AwsAcmpcaCertificateAuthority_OcspConfigurationProperty `field:"optional" json:"ocspConfiguration" yaml:"ocspConfiguration"`
}

