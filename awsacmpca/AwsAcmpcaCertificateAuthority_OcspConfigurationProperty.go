package awsacmpca


// Experimental.
type AwsAcmpcaCertificateAuthority_OcspConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#enabled AwsAcmpcaCertificateAuthority#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acmpca_certificate_authority#ocsp_custom_cname AwsAcmpcaCertificateAuthority#ocsp_custom_cname}.
	// Experimental.
	OcspCustomCname *string `field:"optional" json:"ocspCustomCname" yaml:"ocspCustomCname"`
}

