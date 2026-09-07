package fsx


// Experimental.
type AwsWindowsFileSystem_AuditLogConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#audit_log_destination AwsWindowsFileSystem#audit_log_destination}.
	// Experimental.
	AuditLogDestination *string `field:"optional" json:"auditLogDestination" yaml:"auditLogDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#file_access_audit_log_level AwsWindowsFileSystem#file_access_audit_log_level}.
	// Experimental.
	FileAccessAuditLogLevel *string `field:"optional" json:"fileAccessAuditLogLevel" yaml:"fileAccessAuditLogLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#file_share_access_audit_log_level AwsWindowsFileSystem#file_share_access_audit_log_level}.
	// Experimental.
	FileShareAccessAuditLogLevel *string `field:"optional" json:"fileShareAccessAuditLogLevel" yaml:"fileShareAccessAuditLogLevel"`
}

