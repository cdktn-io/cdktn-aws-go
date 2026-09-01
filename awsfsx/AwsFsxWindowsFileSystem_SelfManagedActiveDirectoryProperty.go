package awsfsx


// Experimental.
type AwsFsxWindowsFileSystem_SelfManagedActiveDirectoryProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#dns_ips AwsFsxWindowsFileSystem#dns_ips}.
	// Experimental.
	DnsIps *[]*string `field:"required" json:"dnsIps" yaml:"dnsIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#domain_name AwsFsxWindowsFileSystem#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#domain_join_service_account_secret AwsFsxWindowsFileSystem#domain_join_service_account_secret}.
	// Experimental.
	DomainJoinServiceAccountSecret *string `field:"optional" json:"domainJoinServiceAccountSecret" yaml:"domainJoinServiceAccountSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#file_system_administrators_group AwsFsxWindowsFileSystem#file_system_administrators_group}.
	// Experimental.
	FileSystemAdministratorsGroup *string `field:"optional" json:"fileSystemAdministratorsGroup" yaml:"fileSystemAdministratorsGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#organizational_unit_distinguished_name AwsFsxWindowsFileSystem#organizational_unit_distinguished_name}.
	// Experimental.
	OrganizationalUnitDistinguishedName *string `field:"optional" json:"organizationalUnitDistinguishedName" yaml:"organizationalUnitDistinguishedName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#password AwsFsxWindowsFileSystem#password}.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#password_wo AwsFsxWindowsFileSystem#password_wo}.
	// Experimental.
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#password_wo_version AwsFsxWindowsFileSystem#password_wo_version}.
	// Experimental.
	PasswordWoVersion *float64 `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#username AwsFsxWindowsFileSystem#username}.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

