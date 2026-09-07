package cloudfront


// Experimental.
type AwsMultitenantDistribution_OriginMtlsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#client_certificate_arn AwsMultitenantDistribution#client_certificate_arn}.
	// Experimental.
	ClientCertificateArn *string `field:"required" json:"clientCertificateArn" yaml:"clientCertificateArn"`
}

