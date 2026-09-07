package amplify


// Experimental.
type AwsDomainAssociation_CertificateSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_domain_association#type AwsDomainAssociation#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_domain_association#custom_certificate_arn AwsDomainAssociation#custom_certificate_arn}.
	// Experimental.
	CustomCertificateArn *string `field:"optional" json:"customCertificateArn" yaml:"customCertificateArn"`
}

