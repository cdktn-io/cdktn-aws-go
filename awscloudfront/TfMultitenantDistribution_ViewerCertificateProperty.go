package awscloudfront


// Experimental.
type TfMultitenantDistribution_ViewerCertificateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#acm_certificate_arn TfMultitenantDistribution#acm_certificate_arn}.
	// Experimental.
	AcmCertificateArn *string `field:"optional" json:"acmCertificateArn" yaml:"acmCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#cloudfront_default_certificate TfMultitenantDistribution#cloudfront_default_certificate}.
	// Experimental.
	CloudfrontDefaultCertificate interface{} `field:"optional" json:"cloudfrontDefaultCertificate" yaml:"cloudfrontDefaultCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#minimum_protocol_version TfMultitenantDistribution#minimum_protocol_version}.
	// Experimental.
	MinimumProtocolVersion *string `field:"optional" json:"minimumProtocolVersion" yaml:"minimumProtocolVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#ssl_support_method TfMultitenantDistribution#ssl_support_method}.
	// Experimental.
	SslSupportMethod *string `field:"optional" json:"sslSupportMethod" yaml:"sslSupportMethod"`
}

