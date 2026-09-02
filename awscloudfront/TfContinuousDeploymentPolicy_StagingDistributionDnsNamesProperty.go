package awscloudfront


// Experimental.
type TfContinuousDeploymentPolicy_StagingDistributionDnsNamesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#quantity TfContinuousDeploymentPolicy#quantity}.
	// Experimental.
	Quantity *float64 `field:"required" json:"quantity" yaml:"quantity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#items TfContinuousDeploymentPolicy#items}.
	// Experimental.
	Items *[]*string `field:"optional" json:"items" yaml:"items"`
}

