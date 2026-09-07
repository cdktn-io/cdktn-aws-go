package codedeploy


// Experimental.
type AwsDeploymentGroup_DeploymentStyleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#deployment_option AwsDeploymentGroup#deployment_option}.
	// Experimental.
	DeploymentOption *string `field:"optional" json:"deploymentOption" yaml:"deploymentOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#deployment_type AwsDeploymentGroup#deployment_type}.
	// Experimental.
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
}

