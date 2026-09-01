package awscloudfront


// Experimental.
type AwsCloudfrontMultitenantDistribution_FailoverCriteriaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#status_codes AwsCloudfrontMultitenantDistribution#status_codes}.
	// Experimental.
	StatusCodes *[]*float64 `field:"required" json:"statusCodes" yaml:"statusCodes"`
}

