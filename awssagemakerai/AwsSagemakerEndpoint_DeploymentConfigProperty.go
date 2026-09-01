package awssagemakerai


// Experimental.
type AwsSagemakerEndpoint_DeploymentConfigProperty struct {
	// auto_rollback_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#auto_rollback_configuration AwsSagemakerEndpoint#auto_rollback_configuration}
	// Experimental.
	AutoRollbackConfiguration *AwsSagemakerEndpoint_AutoRollbackConfigurationProperty `field:"optional" json:"autoRollbackConfiguration" yaml:"autoRollbackConfiguration"`
	// blue_green_update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#blue_green_update_policy AwsSagemakerEndpoint#blue_green_update_policy}
	// Experimental.
	BlueGreenUpdatePolicy *AwsSagemakerEndpoint_BlueGreenUpdatePolicyProperty `field:"optional" json:"blueGreenUpdatePolicy" yaml:"blueGreenUpdatePolicy"`
	// rolling_update_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint#rolling_update_policy AwsSagemakerEndpoint#rolling_update_policy}
	// Experimental.
	RollingUpdatePolicy *AwsSagemakerEndpoint_RollingUpdatePolicyProperty `field:"optional" json:"rollingUpdatePolicy" yaml:"rollingUpdatePolicy"`
}

