package awscloudfront


// Experimental.
type AwsCloudfrontContinuousDeploymentPolicy_SingleHeaderConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#header AwsCloudfrontContinuousDeploymentPolicy#header}.
	// Experimental.
	Header *string `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#value AwsCloudfrontContinuousDeploymentPolicy#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

