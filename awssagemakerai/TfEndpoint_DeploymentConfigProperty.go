package awssagemakerai


// Experimental.
type TfEndpoint_DeploymentConfigProperty struct {
	// auto_rollback_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#auto_rollback_configuration TfEndpoint#auto_rollback_configuration}
	// Experimental.
	AutoRollbackConfiguration *TfEndpoint_AutoRollbackConfigurationProperty `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
	// blue_green_update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#blue_green_update_policy TfEndpoint#blue_green_update_policy}
	// Experimental.
	BlueGreenUpdatePolicy *TfEndpoint_BlueGreenUpdatePolicyProperty `field:"optional" json:"blueGreenUpdatePolicy" yaml:"blueGreenUpdatePolicy"`
	// rolling_update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#rolling_update_policy TfEndpoint#rolling_update_policy}
	// Experimental.
	RollingUpdatePolicy *TfEndpoint_RollingUpdatePolicyProperty `field:"optional" json:"rollingUpdatePolicy" yaml:"rollingUpdatePolicy"`
}

