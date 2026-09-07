package cloudfront


// Experimental.
type AwsDistribution_OriginMtlsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#client_certificate_arn AwsDistribution#client_certificate_arn}.
	// Experimental.
	ClientCertificateArn *string `field:"required" json:"clientCertificateArn" yaml:"clientCertificateArn"`
}

