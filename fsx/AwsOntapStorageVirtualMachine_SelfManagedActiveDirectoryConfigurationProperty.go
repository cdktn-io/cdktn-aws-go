package fsx


// Experimental.
type AwsOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_storage_virtual_machine#dns_ips AwsOntapStorageVirtualMachine#dns_ips}.
	// Experimental.
	DnsIps *[]*string `field:"required" json:"dnsIps" yaml:"dnsIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_storage_virtual_machine#domain_name AwsOntapStorageVirtualMachine#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_storage_virtual_machine#password AwsOntapStorageVirtualMachine#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_storage_virtual_machine#username AwsOntapStorageVirtualMachine#username}.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_storage_virtual_machine#file_system_administrators_group AwsOntapStorageVirtualMachine#file_system_administrators_group}.
	// Experimental.
	FileSystemAdministratorsGroup *string `field:"optional" json:"fileSystemAdministratorsGroup" yaml:"fileSystemAdministratorsGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_ontap_storage_virtual_machine#organizational_unit_distinguished_name AwsOntapStorageVirtualMachine#organizational_unit_distinguished_name}.
	// Experimental.
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
}

