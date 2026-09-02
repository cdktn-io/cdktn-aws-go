package awsacmpca


// Experimental.
type TfCertificateAuthority_RevocationConfigurationProperty struct {
	// crl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#crl_configuration TfCertificateAuthority#crl_configuration}
	// Experimental.
	CrlConfiguration *TfCertificateAuthority_CrlConfigurationProperty `field:"optional" json:"crlConfiguration" yaml:"crlConfiguration"`
	// ocsp_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#ocsp_configuration TfCertificateAuthority#ocsp_configuration}
	// Experimental.
	OcspConfiguration *TfCertificateAuthority_OcspConfigurationProperty `field:"optional" json:"ocspConfiguration" yaml:"ocspConfiguration"`
}

