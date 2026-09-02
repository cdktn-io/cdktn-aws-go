package awsfsx

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWindowsFileSystemConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#subnet_ids TfWindowsFileSystem#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#throughput_capacity TfWindowsFileSystem#throughput_capacity}.
	// Experimental.
	ThroughputCapacity *float64 `field:"required" json:"throughputCapacity" yaml:"throughputCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#active_directory_id TfWindowsFileSystem#active_directory_id}.
	// Experimental.
	ActiveDirectoryId *string `field:"optional" json:"activeDirectoryId" yaml:"activeDirectoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#aliases TfWindowsFileSystem#aliases}.
	// Experimental.
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// audit_log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#audit_log_configuration TfWindowsFileSystem#audit_log_configuration}
	// Experimental.
	AuditLogConfiguration *TfWindowsFileSystem_AuditLogConfigurationProperty `field:"optional" json:"auditLogConfiguration" yaml:"auditLogConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#automatic_backup_retention_days TfWindowsFileSystem#automatic_backup_retention_days}.
	// Experimental.
	AutomaticBackupRetentionDays *float64 `field:"optional" json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#backup_id TfWindowsFileSystem#backup_id}.
	// Experimental.
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#copy_tags_to_backups TfWindowsFileSystem#copy_tags_to_backups}.
	// Experimental.
	CopyTagsToBackups interface{} `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#daily_automatic_backup_start_time TfWindowsFileSystem#daily_automatic_backup_start_time}.
	// Experimental.
	DailyAutomaticBackupStartTime *string `field:"optional" json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#deployment_type TfWindowsFileSystem#deployment_type}.
	// Experimental.
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// disk_iops_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#disk_iops_configuration TfWindowsFileSystem#disk_iops_configuration}
	// Experimental.
	DiskIopsConfiguration *TfWindowsFileSystem_DiskIopsConfigurationProperty `field:"optional" json:"diskIopsConfiguration" yaml:"diskIopsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#final_backup_tags TfWindowsFileSystem#final_backup_tags}.
	// Experimental.
	FinalBackupTags *map[string]*string `field:"optional" json:"finalBackupTags" yaml:"finalBackupTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#id TfWindowsFileSystem#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#kms_key_id TfWindowsFileSystem#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#network_type TfWindowsFileSystem#network_type}.
	// Experimental.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#preferred_subnet_id TfWindowsFileSystem#preferred_subnet_id}.
	// Experimental.
	PreferredSubnetId *string `field:"optional" json:"preferredSubnetId" yaml:"preferredSubnetId"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#region TfWindowsFileSystem#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#security_group_ids TfWindowsFileSystem#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// self_managed_active_directory block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#self_managed_active_directory TfWindowsFileSystem#self_managed_active_directory}
	// Experimental.
	SelfManagedActiveDirectory *TfWindowsFileSystem_SelfManagedActiveDirectoryProperty `field:"optional" json:"selfManagedActiveDirectory" yaml:"selfManagedActiveDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#skip_final_backup TfWindowsFileSystem#skip_final_backup}.
	// Experimental.
	SkipFinalBackup interface{} `field:"optional" json:"skipFinalBackup" yaml:"skipFinalBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#storage_capacity TfWindowsFileSystem#storage_capacity}.
	// Experimental.
	StorageCapacity *float64 `field:"optional" json:"storageCapacity" yaml:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#storage_type TfWindowsFileSystem#storage_type}.
	// Experimental.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#tags TfWindowsFileSystem#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#tags_all TfWindowsFileSystem#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#timeouts TfWindowsFileSystem#timeouts}
	// Experimental.
	Timeouts *TfWindowsFileSystem_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_windows_file_system#weekly_maintenance_start_time TfWindowsFileSystem#weekly_maintenance_start_time}.
	// Experimental.
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

