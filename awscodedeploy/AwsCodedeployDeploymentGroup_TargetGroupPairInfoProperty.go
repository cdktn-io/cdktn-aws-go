package awscodedeploy


// Experimental.
type AwsCodedeployDeploymentGroup_TargetGroupPairInfoProperty struct {
	// prod_traffic_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#prod_traffic_route AwsCodedeployDeploymentGroup#prod_traffic_route}
	// Experimental.
	ProdTrafficRoute *AwsCodedeployDeploymentGroup_ProdTrafficRouteProperty `field:"required" json:"prodTrafficRoute" yaml:"prodTrafficRoute"`
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#target_group AwsCodedeployDeploymentGroup#target_group}
	// Experimental.
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// test_traffic_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#test_traffic_route AwsCodedeployDeploymentGroup#test_traffic_route}
	// Experimental.
	TestTrafficRoute *AwsCodedeployDeploymentGroup_TestTrafficRouteProperty `field:"optional" json:"testTrafficRoute" yaml:"testTrafficRoute"`
}

