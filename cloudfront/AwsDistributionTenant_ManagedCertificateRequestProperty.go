package cloudfront


// Experimental.
type AwsDistributionTenant_ManagedCertificateRequestProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution_tenant#certificate_transparency_logging_preference AwsDistributionTenant#certificate_transparency_logging_preference}.
	// Experimental.
	CertificateTransparencyLoggingPreference *string `field:"optional" json:"certificateTransparencyLoggingPreference" yaml:"certificateTransparencyLoggingPreference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution_tenant#primary_domain_name AwsDistributionTenant#primary_domain_name}.
	// Experimental.
	PrimaryDomainName *string `field:"optional" json:"primaryDomainName" yaml:"primaryDomainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution_tenant#validation_token_host AwsDistributionTenant#validation_token_host}.
	// Experimental.
	ValidationTokenHost *string `field:"optional" json:"validationTokenHost" yaml:"validationTokenHost"`
}

