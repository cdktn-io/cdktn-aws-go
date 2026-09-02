package awsfsx

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLustreFileSystemConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#subnet_ids TfLustreFileSystem#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"required" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#auto_import_policy TfLustreFileSystem#auto_import_policy}.
	// Experimental.
	AutoImportPolicy *string `field:"optional" json:"autoImportPolicy" yaml:"autoImportPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#automatic_backup_retention_days TfLustreFileSystem#automatic_backup_retention_days}.
	// Experimental.
	AutomaticBackupRetentionDays *float64 `field:"optional" json:"automaticBackupRetentionDays" yaml:"automaticBackupRetentionDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#backup_id TfLustreFileSystem#backup_id}.
	// Experimental.
	BackupId *string `field:"optional" json:"backupId" yaml:"backupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#copy_tags_to_backups TfLustreFileSystem#copy_tags_to_backups}.
	// Experimental.
	CopyTagsToBackups interface{} `field:"optional" json:"copyTagsToBackups" yaml:"copyTagsToBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#daily_automatic_backup_start_time TfLustreFileSystem#daily_automatic_backup_start_time}.
	// Experimental.
	DailyAutomaticBackupStartTime *string `field:"optional" json:"dailyAutomaticBackupStartTime" yaml:"dailyAutomaticBackupStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#data_compression_type TfLustreFileSystem#data_compression_type}.
	// Experimental.
	DataCompressionType *string `field:"optional" json:"dataCompressionType" yaml:"dataCompressionType"`
	// data_read_cache_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#data_read_cache_configuration TfLustreFileSystem#data_read_cache_configuration}
	// Experimental.
	DataReadCacheConfiguration *TfLustreFileSystem_DataReadCacheConfigurationProperty `field:"optional" json:"dataReadCacheConfiguration" yaml:"dataReadCacheConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#deployment_type TfLustreFileSystem#deployment_type}.
	// Experimental.
	DeploymentType *string `field:"optional" json:"deploymentType" yaml:"deploymentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#drive_cache_type TfLustreFileSystem#drive_cache_type}.
	// Experimental.
	DriveCacheType *string `field:"optional" json:"driveCacheType" yaml:"driveCacheType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#efa_enabled TfLustreFileSystem#efa_enabled}.
	// Experimental.
	EfaEnabled interface{} `field:"optional" json:"efaEnabled" yaml:"efaEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#export_path TfLustreFileSystem#export_path}.
	// Experimental.
	ExportPath *string `field:"optional" json:"exportPath" yaml:"exportPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#file_system_type_version TfLustreFileSystem#file_system_type_version}.
	// Experimental.
	FileSystemTypeVersion *string `field:"optional" json:"fileSystemTypeVersion" yaml:"fileSystemTypeVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#final_backup_tags TfLustreFileSystem#final_backup_tags}.
	// Experimental.
	FinalBackupTags *map[string]*string `field:"optional" json:"finalBackupTags" yaml:"finalBackupTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#id TfLustreFileSystem#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#imported_file_chunk_size TfLustreFileSystem#imported_file_chunk_size}.
	// Experimental.
	ImportedFileChunkSize *float64 `field:"optional" json:"importedFileChunkSize" yaml:"importedFileChunkSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#import_path TfLustreFileSystem#import_path}.
	// Experimental.
	ImportPath *string `field:"optional" json:"importPath" yaml:"importPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#kms_key_id TfLustreFileSystem#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// log_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#log_configuration TfLustreFileSystem#log_configuration}
	// Experimental.
	LogConfiguration *TfLustreFileSystem_LogConfigurationProperty `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// metadata_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#metadata_configuration TfLustreFileSystem#metadata_configuration}
	// Experimental.
	MetadataConfiguration *TfLustreFileSystem_MetadataConfigurationProperty `field:"optional" json:"metadataConfiguration" yaml:"metadataConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#per_unit_storage_throughput TfLustreFileSystem#per_unit_storage_throughput}.
	// Experimental.
	PerUnitStorageThroughput *float64 `field:"optional" json:"perUnitStorageThroughput" yaml:"perUnitStorageThroughput"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#region TfLustreFileSystem#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// root_squash_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#root_squash_configuration TfLustreFileSystem#root_squash_configuration}
	// Experimental.
	RootSquashConfiguration *TfLustreFileSystem_RootSquashConfigurationProperty `field:"optional" json:"rootSquashConfiguration" yaml:"rootSquashConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#security_group_ids TfLustreFileSystem#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#skip_final_backup TfLustreFileSystem#skip_final_backup}.
	// Experimental.
	SkipFinalBackup interface{} `field:"optional" json:"skipFinalBackup" yaml:"skipFinalBackup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#storage_capacity TfLustreFileSystem#storage_capacity}.
	// Experimental.
	StorageCapacity *float64 `field:"optional" json:"storageCapacity" yaml:"storageCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#storage_type TfLustreFileSystem#storage_type}.
	// Experimental.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#tags TfLustreFileSystem#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#tags_all TfLustreFileSystem#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#throughput_capacity TfLustreFileSystem#throughput_capacity}.
	// Experimental.
	ThroughputCapacity *float64 `field:"optional" json:"throughputCapacity" yaml:"throughputCapacity"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#timeouts TfLustreFileSystem#timeouts}
	// Experimental.
	Timeouts *TfLustreFileSystem_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_lustre_file_system#weekly_maintenance_start_time TfLustreFileSystem#weekly_maintenance_start_time}.
	// Experimental.
	WeeklyMaintenanceStartTime *string `field:"optional" json:"weeklyMaintenanceStartTime" yaml:"weeklyMaintenanceStartTime"`
}

