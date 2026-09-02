package awscloudfront


// Experimental.
type TfDistribution_S3OriginConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#origin_access_identity TfDistribution#origin_access_identity}.
	// Experimental.
	OriginAccessIdentity *string `field:"required" json:"originAccessIdentity" yaml:"originAccessIdentity"`
}

