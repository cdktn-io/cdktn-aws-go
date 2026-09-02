package awsfsx


// Experimental.
type TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_storage_virtual_machine#netbios_name TfOntapStorageVirtualMachine#netbios_name}.
	// Experimental.
	NetbiosName *string `field:"optional" json:"netbiosName" yaml:"netbiosName"`
	// self_managed_active_directory_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_storage_virtual_machine#self_managed_active_directory_configuration TfOntapStorageVirtualMachine#self_managed_active_directory_configuration}
	// Experimental.
	SelfManagedActiveDirectoryConfiguration *TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty `field:"optional" json:"selfManagedActiveDirectoryConfiguration" yaml:"selfManagedActiveDirectoryConfiguration"`
}

