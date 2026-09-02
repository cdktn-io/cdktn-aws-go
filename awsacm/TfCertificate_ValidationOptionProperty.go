package awsacm


// Experimental.
type TfCertificate_ValidationOptionProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acm_certificate#domain_name TfCertificate#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/acm_certificate#validation_domain TfCertificate#validation_domain}.
	// Experimental.
	ValidationDomain *string `field:"required" json:"validationDomain" yaml:"validationDomain"`
}

