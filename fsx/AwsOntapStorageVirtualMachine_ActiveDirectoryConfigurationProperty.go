package fsx


// Experimental.
type AwsOntapStorageVirtualMachine_ActiveDirectoryConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_storage_virtual_machine#netbios_name AwsOntapStorageVirtualMachine#netbios_name}.
	// Experimental.
	NetbiosName *string `field:"optional" json:"netbiosName" yaml:"netbiosName"`
	// self_managed_active_directory_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_storage_virtual_machine#self_managed_active_directory_configuration AwsOntapStorageVirtualMachine#self_managed_active_directory_configuration}
	// Experimental.
	SelfManagedActiveDirectoryConfiguration *AwsOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty `field:"optional" json:"selfManagedActiveDirectoryConfiguration" yaml:"selfManagedActiveDirectoryConfiguration"`
}

