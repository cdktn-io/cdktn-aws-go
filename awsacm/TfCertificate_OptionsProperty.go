package awsacm


// Experimental.
type TfCertificate_OptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acm_certificate#certificate_transparency_logging_preference TfCertificate#certificate_transparency_logging_preference}.
	// Experimental.
	CertificateTransparencyLoggingPreference *string `field:"optional" json:"certificateTransparencyLoggingPreference" yaml:"certificateTransparencyLoggingPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acm_certificate#export TfCertificate#export}.
	// Experimental.
	Export *string `field:"optional" json:"export" yaml:"export"`
}

