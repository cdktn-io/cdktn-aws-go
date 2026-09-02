package awscloudfront


// Experimental.
type TfContinuousDeploymentPolicy_SingleHeaderConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#header TfContinuousDeploymentPolicy#header}.
	// Experimental.
	Header *string `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#value TfContinuousDeploymentPolicy#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

