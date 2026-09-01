package awsacm


// Experimental.
type AwsAcmCertificate_ValidationOptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acm_certificate#domain_name AwsAcmCertificate#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acm_certificate#validation_domain AwsAcmCertificate#validation_domain}.
	// Experimental.
	ValidationDomain *string `field:"required" json:"validationDomain" yaml:"validationDomain"`
}

