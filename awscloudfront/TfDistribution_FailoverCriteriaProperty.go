package awscloudfront


// Experimental.
type TfDistribution_FailoverCriteriaProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#status_codes TfDistribution#status_codes}.
	// Experimental.
	StatusCodes *[]*float64 `field:"required" json:"statusCodes" yaml:"statusCodes"`
}

