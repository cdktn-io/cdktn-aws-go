package awsecs


// Experimental.
type TfCluster_ConfigurationProperty struct {
	// execute_command_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#execute_command_configuration TfCluster#execute_command_configuration}
	// Experimental.
	ExecuteCommandConfiguration *TfCluster_ExecuteCommandConfigurationProperty `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
	// managed_storage_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_cluster#managed_storage_configuration TfCluster#managed_storage_configuration}
	// Experimental.
	ManagedStorageConfiguration *TfCluster_ManagedStorageConfigurationProperty `field:"optional" json:"managedStorageConfiguration" yaml:"managedStorageConfiguration"`
}

