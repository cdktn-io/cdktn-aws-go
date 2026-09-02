package awscodedeploy


// Experimental.
type TfDeploymentGroup_TargetGroupPairInfoProperty struct {
	// prod_traffic_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#prod_traffic_route TfDeploymentGroup#prod_traffic_route}
	// Experimental.
	ProdTrafficRoute *TfDeploymentGroup_ProdTrafficRouteProperty `field:"required" json:"prodTrafficRoute" yaml:"prodTrafficRoute"`
	// target_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#target_group TfDeploymentGroup#target_group}
	// Experimental.
	TargetGroup interface{} `field:"required" json:"targetGroup" yaml:"targetGroup"`
	// test_traffic_route block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#test_traffic_route TfDeploymentGroup#test_traffic_route}
	// Experimental.
	TestTrafficRoute *TfDeploymentGroup_TestTrafficRouteProperty `field:"optional" json:"testTrafficRoute" yaml:"testTrafficRoute"`
}

