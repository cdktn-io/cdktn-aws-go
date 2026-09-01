package awscloudfront


// Experimental.
type AwsCloudfrontDistribution_FailoverCriteriaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#status_codes AwsCloudfrontDistribution#status_codes}.
	// Experimental.
	StatusCodes *[]*float64 `field:"required" json:"statusCodes" yaml:"statusCodes"`
}

