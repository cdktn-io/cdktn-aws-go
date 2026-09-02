package awsrolesanywhere


// Experimental.
type TfTrustAnchor_SourceDataProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#acm_pca_arn TfTrustAnchor#acm_pca_arn}.
	// Experimental.
	AcmPcaArn *string `field:"optional" json:"acmPcaArn" yaml:"acmPcaArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rolesanywhere_trust_anchor#x509_certificate_data TfTrustAnchor#x509_certificate_data}.
	// Experimental.
	X509CertificateData *string `field:"optional" json:"x509CertificateData" yaml:"x509CertificateData"`
}

