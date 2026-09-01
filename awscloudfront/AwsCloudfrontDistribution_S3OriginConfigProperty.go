package awscloudfront


// Experimental.
type AwsCloudfrontDistribution_S3OriginConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#origin_access_identity AwsCloudfrontDistribution#origin_access_identity}.
	// Experimental.
	OriginAccessIdentity *string `field:"required" json:"originAccessIdentity" yaml:"originAccessIdentity"`
}

