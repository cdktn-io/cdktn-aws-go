package ecs


// Experimental.
type AwsCluster_ConfigurationProperty struct {
	// execute_command_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#execute_command_configuration AwsCluster#execute_command_configuration}
	// Experimental.
	ExecuteCommandConfiguration *AwsCluster_ExecuteCommandConfigurationProperty `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
	// managed_storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#managed_storage_configuration AwsCluster#managed_storage_configuration}
	// Experimental.
	ManagedStorageConfiguration *AwsCluster_ManagedStorageConfigurationProperty `field:"optional" json:"managedStorageConfiguration" yaml:"managedStorageConfiguration"`
}

