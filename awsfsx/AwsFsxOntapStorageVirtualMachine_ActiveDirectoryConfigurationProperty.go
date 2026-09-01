package awsfsx


// Experimental.
type AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_storage_virtual_machine#netbios_name AwsFsxOntapStorageVirtualMachine#netbios_name}.
	// Experimental.
	NetbiosName *string `field:"optional" json:"netbiosName" yaml:"netbiosName"`
	// self_managed_active_directory_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_storage_virtual_machine#self_managed_active_directory_configuration AwsFsxOntapStorageVirtualMachine#self_managed_active_directory_configuration}
	// Experimental.
	SelfManagedActiveDirectoryConfiguration *AwsFsxOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty `field:"optional" json:"selfManagedActiveDirectoryConfiguration" yaml:"selfManagedActiveDirectoryConfiguration"`
}

