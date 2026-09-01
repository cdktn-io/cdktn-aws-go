package awscloudfront


// Experimental.
type AwsCloudfrontMultitenantDistribution_ViewerCertificateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#acm_certificate_arn AwsCloudfrontMultitenantDistribution#acm_certificate_arn}.
	// Experimental.
	AcmCertificateArn *string `field:"optional" json:"acmCertificateArn" yaml:"acmCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#cloudfront_default_certificate AwsCloudfrontMultitenantDistribution#cloudfront_default_certificate}.
	// Experimental.
	CloudfrontDefaultCertificate interface{} `field:"optional" json:"cloudfrontDefaultCertificate" yaml:"cloudfrontDefaultCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#minimum_protocol_version AwsCloudfrontMultitenantDistribution#minimum_protocol_version}.
	// Experimental.
	MinimumProtocolVersion *string `field:"optional" json:"minimumProtocolVersion" yaml:"minimumProtocolVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#ssl_support_method AwsCloudfrontMultitenantDistribution#ssl_support_method}.
	// Experimental.
	SslSupportMethod *string `field:"optional" json:"sslSupportMethod" yaml:"sslSupportMethod"`
}

