package awscloudfront


// Experimental.
type AwsCloudfrontContinuousDeploymentPolicy_StagingDistributionDnsNamesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#quantity AwsCloudfrontContinuousDeploymentPolicy#quantity}.
	// Experimental.
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#items AwsCloudfrontContinuousDeploymentPolicy#items}.
	// Experimental.
	Items *[]*string `field:"optional" json:"items" yaml:"items"`
}

