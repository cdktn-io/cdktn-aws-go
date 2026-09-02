package awsrds

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDbInstanceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#instance_class TfDbInstance#instance_class}.
	// Experimental.
	InstanceClass *string `field:"required" json:"instanceClass" yaml:"instanceClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#allocated_storage TfDbInstance#allocated_storage}.
	// Experimental.
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#allow_major_version_upgrade TfDbInstance#allow_major_version_upgrade}.
	// Experimental.
	AllowMajorVersionUpgrade interface{} `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#apply_immediately TfDbInstance#apply_immediately}.
	// Experimental.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#auto_minor_version_upgrade TfDbInstance#auto_minor_version_upgrade}.
	// Experimental.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#availability_zone TfDbInstance#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#backup_retention_period TfDbInstance#backup_retention_period}.
	// Experimental.
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#backup_target TfDbInstance#backup_target}.
	// Experimental.
	BackupTarget *string `field:"optional" json:"backupTarget" yaml:"backupTarget"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#backup_window TfDbInstance#backup_window}.
	// Experimental.
	BackupWindow *string `field:"optional" json:"backupWindow" yaml:"backupWindow"`
	// blue_green_update block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#blue_green_update TfDbInstance#blue_green_update}
	// Experimental.
	BlueGreenUpdate *TfDbInstance_BlueGreenUpdateProperty `field:"optional" json:"blueGreenUpdate" yaml:"blueGreenUpdate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#ca_cert_identifier TfDbInstance#ca_cert_identifier}.
	// Experimental.
	CaCertIdentifier *string `field:"optional" json:"caCertIdentifier" yaml:"caCertIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#character_set_name TfDbInstance#character_set_name}.
	// Experimental.
	CharacterSetName *string `field:"optional" json:"characterSetName" yaml:"characterSetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#copy_tags_to_snapshot TfDbInstance#copy_tags_to_snapshot}.
	// Experimental.
	CopyTagsToSnapshot interface{} `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#customer_owned_ip_enabled TfDbInstance#customer_owned_ip_enabled}.
	// Experimental.
	CustomerOwnedIpEnabled interface{} `field:"optional" json:"customerOwnedIpEnabled" yaml:"customerOwnedIpEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#custom_iam_instance_profile TfDbInstance#custom_iam_instance_profile}.
	// Experimental.
	CustomIamInstanceProfile *string `field:"optional" json:"customIamInstanceProfile" yaml:"customIamInstanceProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#database_insights_mode TfDbInstance#database_insights_mode}.
	// Experimental.
	DatabaseInsightsMode *string `field:"optional" json:"databaseInsightsMode" yaml:"databaseInsightsMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#db_name TfDbInstance#db_name}.
	// Experimental.
	DbName *string `field:"optional" json:"dbName" yaml:"dbName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#db_subnet_group_name TfDbInstance#db_subnet_group_name}.
	// Experimental.
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#dedicated_log_volume TfDbInstance#dedicated_log_volume}.
	// Experimental.
	DedicatedLogVolume interface{} `field:"optional" json:"dedicatedLogVolume" yaml:"dedicatedLogVolume"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#delete_automated_backups TfDbInstance#delete_automated_backups}.
	// Experimental.
	DeleteAutomatedBackups interface{} `field:"optional" json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#deletion_protection TfDbInstance#deletion_protection}.
	// Experimental.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#domain TfDbInstance#domain}.
	// Experimental.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#domain_auth_secret_arn TfDbInstance#domain_auth_secret_arn}.
	// Experimental.
	DomainAuthSecretArn *string `field:"optional" json:"domainAuthSecretArn" yaml:"domainAuthSecretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#domain_dns_ips TfDbInstance#domain_dns_ips}.
	// Experimental.
	DomainDnsIps *[]*string `field:"optional" json:"domainDnsIps" yaml:"domainDnsIps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#domain_fqdn TfDbInstance#domain_fqdn}.
	// Experimental.
	DomainFqdn *string `field:"optional" json:"domainFqdn" yaml:"domainFqdn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#domain_iam_role_name TfDbInstance#domain_iam_role_name}.
	// Experimental.
	DomainIamRoleName *string `field:"optional" json:"domainIamRoleName" yaml:"domainIamRoleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#domain_ou TfDbInstance#domain_ou}.
	// Experimental.
	DomainOu *string `field:"optional" json:"domainOu" yaml:"domainOu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#enabled_cloudwatch_logs_exports TfDbInstance#enabled_cloudwatch_logs_exports}.
	// Experimental.
	EnabledCloudwatchLogsExports *[]*string `field:"optional" json:"enabledCloudwatchLogsExports" yaml:"enabledCloudwatchLogsExports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#engine TfDbInstance#engine}.
	// Experimental.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#engine_lifecycle_support TfDbInstance#engine_lifecycle_support}.
	// Experimental.
	EngineLifecycleSupport *string `field:"optional" json:"engineLifecycleSupport" yaml:"engineLifecycleSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#engine_version TfDbInstance#engine_version}.
	// Experimental.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#final_snapshot_identifier TfDbInstance#final_snapshot_identifier}.
	// Experimental.
	FinalSnapshotIdentifier *string `field:"optional" json:"finalSnapshotIdentifier" yaml:"finalSnapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#iam_database_authentication_enabled TfDbInstance#iam_database_authentication_enabled}.
	// Experimental.
	IamDatabaseAuthenticationEnabled interface{} `field:"optional" json:"iamDatabaseAuthenticationEnabled" yaml:"iamDatabaseAuthenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#id TfDbInstance#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#identifier TfDbInstance#identifier}.
	// Experimental.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#identifier_prefix TfDbInstance#identifier_prefix}.
	// Experimental.
	IdentifierPrefix *string `field:"optional" json:"identifierPrefix" yaml:"identifierPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#iops TfDbInstance#iops}.
	// Experimental.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#kms_key_id TfDbInstance#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#license_model TfDbInstance#license_model}.
	// Experimental.
	LicenseModel *string `field:"optional" json:"licenseModel" yaml:"licenseModel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#maintenance_window TfDbInstance#maintenance_window}.
	// Experimental.
	MaintenanceWindow *string `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#manage_master_user_password TfDbInstance#manage_master_user_password}.
	// Experimental.
	ManageMasterUserPassword interface{} `field:"optional" json:"manageMasterUserPassword" yaml:"manageMasterUserPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#master_user_secret_kms_key_id TfDbInstance#master_user_secret_kms_key_id}.
	// Experimental.
	MasterUserSecretKmsKeyId *string `field:"optional" json:"masterUserSecretKmsKeyId" yaml:"masterUserSecretKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#max_allocated_storage TfDbInstance#max_allocated_storage}.
	// Experimental.
	MaxAllocatedStorage *float64 `field:"optional" json:"maxAllocatedStorage" yaml:"maxAllocatedStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#monitoring_interval TfDbInstance#monitoring_interval}.
	// Experimental.
	MonitoringInterval *float64 `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#monitoring_role_arn TfDbInstance#monitoring_role_arn}.
	// Experimental.
	MonitoringRoleArn *string `field:"optional" json:"monitoringRoleArn" yaml:"monitoringRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#multi_az TfDbInstance#multi_az}.
	// Experimental.
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#nchar_character_set_name TfDbInstance#nchar_character_set_name}.
	// Experimental.
	NcharCharacterSetName *string `field:"optional" json:"ncharCharacterSetName" yaml:"ncharCharacterSetName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#network_type TfDbInstance#network_type}.
	// Experimental.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#option_group_name TfDbInstance#option_group_name}.
	// Experimental.
	OptionGroupName *string `field:"optional" json:"optionGroupName" yaml:"optionGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#parameter_group_name TfDbInstance#parameter_group_name}.
	// Experimental.
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#password TfDbInstance#password}.
	// Experimental.
	Password *string `field:"optional" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#password_wo TfDbInstance#password_wo}.
	// Experimental.
	PasswordWo *string `field:"optional" json:"passwordWo" yaml:"passwordWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#password_wo_version TfDbInstance#password_wo_version}.
	// Experimental.
	PasswordWoVersion *float64 `field:"optional" json:"passwordWoVersion" yaml:"passwordWoVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#performance_insights_enabled TfDbInstance#performance_insights_enabled}.
	// Experimental.
	PerformanceInsightsEnabled interface{} `field:"optional" json:"performanceInsightsEnabled" yaml:"performanceInsightsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#performance_insights_kms_key_id TfDbInstance#performance_insights_kms_key_id}.
	// Experimental.
	PerformanceInsightsKmsKeyId *string `field:"optional" json:"performanceInsightsKmsKeyId" yaml:"performanceInsightsKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#performance_insights_retention_period TfDbInstance#performance_insights_retention_period}.
	// Experimental.
	PerformanceInsightsRetentionPeriod *float64 `field:"optional" json:"performanceInsightsRetentionPeriod" yaml:"performanceInsightsRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#port TfDbInstance#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#publicly_accessible TfDbInstance#publicly_accessible}.
	// Experimental.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#region TfDbInstance#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#replica_mode TfDbInstance#replica_mode}.
	// Experimental.
	ReplicaMode *string `field:"optional" json:"replicaMode" yaml:"replicaMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#replicate_source_db TfDbInstance#replicate_source_db}.
	// Experimental.
	ReplicateSourceDb *string `field:"optional" json:"replicateSourceDb" yaml:"replicateSourceDb"`
	// restore_to_point_in_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#restore_to_point_in_time TfDbInstance#restore_to_point_in_time}
	// Experimental.
	RestoreToPointInTime *TfDbInstance_RestoreToPointInTimeProperty `field:"optional" json:"restoreToPointInTime" yaml:"restoreToPointInTime"`
	// s3_import block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#s3_import TfDbInstance#s3_import}
	// Experimental.
	S3Import *TfDbInstance_S3ImportProperty `field:"optional" json:"s3Import" yaml:"s3Import"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#skip_final_snapshot TfDbInstance#skip_final_snapshot}.
	// Experimental.
	SkipFinalSnapshot interface{} `field:"optional" json:"skipFinalSnapshot" yaml:"skipFinalSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#snapshot_identifier TfDbInstance#snapshot_identifier}.
	// Experimental.
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#storage_encrypted TfDbInstance#storage_encrypted}.
	// Experimental.
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#storage_throughput TfDbInstance#storage_throughput}.
	// Experimental.
	StorageThroughput *float64 `field:"optional" json:"storageThroughput" yaml:"storageThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#storage_type TfDbInstance#storage_type}.
	// Experimental.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#tags TfDbInstance#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#tags_all TfDbInstance#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#timeouts TfDbInstance#timeouts}
	// Experimental.
	Timeouts *TfDbInstance_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#timezone TfDbInstance#timezone}.
	// Experimental.
	Timezone *string `field:"optional" json:"timezone" yaml:"timezone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#upgrade_storage_config TfDbInstance#upgrade_storage_config}.
	// Experimental.
	UpgradeStorageConfig interface{} `field:"optional" json:"upgradeStorageConfig" yaml:"upgradeStorageConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#username TfDbInstance#username}.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_instance#vpc_security_group_ids TfDbInstance#vpc_security_group_ids}.
	// Experimental.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

