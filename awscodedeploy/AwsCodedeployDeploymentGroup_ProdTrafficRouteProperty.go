package awscodedeploy


// Experimental.
type AwsCodedeployDeploymentGroup_ProdTrafficRouteProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#listener_arns AwsCodedeployDeploymentGroup#listener_arns}.
	// Experimental.
	ListenerArns *[]*string `field:"required" json:"listenerArns" yaml:"listenerArns"`
}

