package acmpca


// Experimental.
type AwsCertificateAuthority_RevocationConfigurationProperty struct {
	// crl_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#crl_configuration AwsCertificateAuthority#crl_configuration}
	// Experimental.
	CrlConfiguration *AwsCertificateAuthority_CrlConfigurationProperty `field:"optional" json:"crlConfiguration" yaml:"crlConfiguration"`
	// ocsp_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#ocsp_configuration AwsCertificateAuthority#ocsp_configuration}
	// Experimental.
	OcspConfiguration *AwsCertificateAuthority_OcspConfigurationProperty `field:"optional" json:"ocspConfiguration" yaml:"ocspConfiguration"`
}

