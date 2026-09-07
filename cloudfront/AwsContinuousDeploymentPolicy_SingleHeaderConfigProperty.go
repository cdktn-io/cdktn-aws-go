package cloudfront


// Experimental.
type AwsContinuousDeploymentPolicy_SingleHeaderConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#header AwsContinuousDeploymentPolicy#header}.
	// Experimental.
	Header *string `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_continuous_deployment_policy#value AwsContinuousDeploymentPolicy#value}.
	// Experimental.
	Value *string `field:"required" json:"value" yaml:"value"`
}

