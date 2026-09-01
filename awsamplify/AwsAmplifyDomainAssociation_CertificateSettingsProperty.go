package awsamplify


// Experimental.
type AwsAmplifyDomainAssociation_CertificateSettingsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_domain_association#type AwsAmplifyDomainAssociation#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/amplify_domain_association#custom_certificate_arn AwsAmplifyDomainAssociation#custom_certificate_arn}.
	// Experimental.
	CustomCertificateArn *string `field:"optional" json:"customCertificateArn" yaml:"customCertificateArn"`
}

