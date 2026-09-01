package awscloudfront


// Experimental.
type AwsCloudfrontDistribution_ViewerCertificateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#acm_certificate_arn AwsCloudfrontDistribution#acm_certificate_arn}.
	// Experimental.
	AcmCertificateArn *string `field:"optional" json:"acmCertificateArn" yaml:"acmCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#cloudfront_default_certificate AwsCloudfrontDistribution#cloudfront_default_certificate}.
	// Experimental.
	CloudfrontDefaultCertificate interface{} `field:"optional" json:"cloudfrontDefaultCertificate" yaml:"cloudfrontDefaultCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#iam_certificate_id AwsCloudfrontDistribution#iam_certificate_id}.
	// Experimental.
	IamCertificateId *string `field:"optional" json:"iamCertificateId" yaml:"iamCertificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#minimum_protocol_version AwsCloudfrontDistribution#minimum_protocol_version}.
	// Experimental.
	MinimumProtocolVersion *string `field:"optional" json:"minimumProtocolVersion" yaml:"minimumProtocolVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#ssl_support_method AwsCloudfrontDistribution#ssl_support_method}.
	// Experimental.
	SslSupportMethod *string `field:"optional" json:"sslSupportMethod" yaml:"sslSupportMethod"`
}

