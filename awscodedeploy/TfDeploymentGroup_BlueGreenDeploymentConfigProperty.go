package awscodedeploy


// Experimental.
type TfDeploymentGroup_BlueGreenDeploymentConfigProperty struct {
	// deployment_ready_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#deployment_ready_option TfDeploymentGroup#deployment_ready_option}
	// Experimental.
	DeploymentReadyOption *TfDeploymentGroup_DeploymentReadyOptionProperty `field:"optional" json:"deploymentReadyOption" yaml:"deploymentReadyOption"`
	// green_fleet_provisioning_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#green_fleet_provisioning_option TfDeploymentGroup#green_fleet_provisioning_option}
	// Experimental.
	GreenFleetProvisioningOption *TfDeploymentGroup_GreenFleetProvisioningOptionProperty `field:"optional" json:"greenFleetProvisioningOption" yaml:"greenFleetProvisioningOption"`
	// terminate_blue_instances_on_deployment_success block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codedeploy_deployment_group#terminate_blue_instances_on_deployment_success TfDeploymentGroup#terminate_blue_instances_on_deployment_success}
	// Experimental.
	TerminateBlueInstancesOnDeploymentSuccess *TfDeploymentGroup_TerminateBlueInstancesOnDeploymentSuccessProperty `field:"optional" json:"terminateBlueInstancesOnDeploymentSuccess" yaml:"terminateBlueInstancesOnDeploymentSuccess"`
}

