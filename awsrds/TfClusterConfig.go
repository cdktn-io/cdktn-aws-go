package awsrds

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfClusterConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#engine TfCluster#engine}.
	// Experimental.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#allocated_storage TfCluster#allocated_storage}.
	// Experimental.
	AllocatedStorage *float64 `field:"optional" json:"allocatedStorage" yaml:"allocatedStorage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#allow_major_version_upgrade TfCluster#allow_major_version_upgrade}.
	// Experimental.
	AllowMajorVersionUpgrade interface{} `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#apply_immediately TfCluster#apply_immediately}.
	// Experimental.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#auto_minor_version_upgrade TfCluster#auto_minor_version_upgrade}.
	// Experimental.
	AutoMinorVersionUpgrade interface{} `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#availability_zones TfCluster#availability_zones}.
	// Experimental.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#backtrack_window TfCluster#backtrack_window}.
	// Experimental.
	BacktrackWindow *float64 `field:"optional" json:"backtrackWindow" yaml:"backtrackWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#backup_retention_period TfCluster#backup_retention_period}.
	// Experimental.
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#ca_certificate_identifier TfCluster#ca_certificate_identifier}.
	// Experimental.
	CaCertificateIdentifier *string `field:"optional" json:"caCertificateIdentifier" yaml:"caCertificateIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#cluster_identifier TfCluster#cluster_identifier}.
	// Experimental.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#cluster_identifier_prefix TfCluster#cluster_identifier_prefix}.
	// Experimental.
	ClusterIdentifierPrefix *string `field:"optional" json:"clusterIdentifierPrefix" yaml:"clusterIdentifierPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#cluster_members TfCluster#cluster_members}.
	// Experimental.
	ClusterMembers *[]*string `field:"optional" json:"clusterMembers" yaml:"clusterMembers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#cluster_scalability_type TfCluster#cluster_scalability_type}.
	// Experimental.
	ClusterScalabilityType *string `field:"optional" json:"clusterScalabilityType" yaml:"clusterScalabilityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#copy_tags_to_snapshot TfCluster#copy_tags_to_snapshot}.
	// Experimental.
	CopyTagsToSnapshot interface{} `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#database_insights_mode TfCluster#database_insights_mode}.
	// Experimental.
	DatabaseInsightsMode *string `field:"optional" json:"databaseInsightsMode" yaml:"databaseInsightsMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#database_name TfCluster#database_name}.
	// Experimental.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#db_cluster_instance_class TfCluster#db_cluster_instance_class}.
	// Experimental.
	DbClusterInstanceClass *string `field:"optional" json:"dbClusterInstanceClass" yaml:"dbClusterInstanceClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#db_cluster_parameter_group_name TfCluster#db_cluster_parameter_group_name}.
	// Experimental.
	DbClusterParameterGroupName *string `field:"optional" json:"dbClusterParameterGroupName" yaml:"dbClusterParameterGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#db_instance_parameter_group_name TfCluster#db_instance_parameter_group_name}.
	// Experimental.
	DbInstanceParameterGroupName *string `field:"optional" json:"dbInstanceParameterGroupName" yaml:"dbInstanceParameterGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#db_subnet_group_name TfCluster#db_subnet_group_name}.
	// Experimental.
	DbSubnetGroupName *string `field:"optional" json:"dbSubnetGroupName" yaml:"dbSubnetGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#db_system_id TfCluster#db_system_id}.
	// Experimental.
	DbSystemId *string `field:"optional" json:"dbSystemId" yaml:"dbSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#delete_automated_backups TfCluster#delete_automated_backups}.
	// Experimental.
	DeleteAutomatedBackups interface{} `field:"optional" json:"deleteAutomatedBackups" yaml:"deleteAutomatedBackups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#deletion_protection TfCluster#deletion_protection}.
	// Experimental.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#domain TfCluster#domain}.
	// Experimental.
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#domain_iam_role_name TfCluster#domain_iam_role_name}.
	// Experimental.
	DomainIamRoleName *string `field:"optional" json:"domainIamRoleName" yaml:"domainIamRoleName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#enabled_cloudwatch_logs_exports TfCluster#enabled_cloudwatch_logs_exports}.
	// Experimental.
	EnabledCloudwatchLogsExports *[]*string `field:"optional" json:"enabledCloudwatchLogsExports" yaml:"enabledCloudwatchLogsExports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#enable_global_write_forwarding TfCluster#enable_global_write_forwarding}.
	// Experimental.
	EnableGlobalWriteForwarding interface{} `field:"optional" json:"enableGlobalWriteForwarding" yaml:"enableGlobalWriteForwarding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#enable_http_endpoint TfCluster#enable_http_endpoint}.
	// Experimental.
	EnableHttpEndpoint interface{} `field:"optional" json:"enableHttpEndpoint" yaml:"enableHttpEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#enable_local_write_forwarding TfCluster#enable_local_write_forwarding}.
	// Experimental.
	EnableLocalWriteForwarding interface{} `field:"optional" json:"enableLocalWriteForwarding" yaml:"enableLocalWriteForwarding"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#engine_lifecycle_support TfCluster#engine_lifecycle_support}.
	// Experimental.
	EngineLifecycleSupport *string `field:"optional" json:"engineLifecycleSupport" yaml:"engineLifecycleSupport"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#engine_mode TfCluster#engine_mode}.
	// Experimental.
	EngineMode *string `field:"optional" json:"engineMode" yaml:"engineMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#engine_version TfCluster#engine_version}.
	// Experimental.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#final_snapshot_identifier TfCluster#final_snapshot_identifier}.
	// Experimental.
	FinalSnapshotIdentifier *string `field:"optional" json:"finalSnapshotIdentifier" yaml:"finalSnapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#global_cluster_identifier TfCluster#global_cluster_identifier}.
	// Experimental.
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#iam_database_authentication_enabled TfCluster#iam_database_authentication_enabled}.
	// Experimental.
	IamDatabaseAuthenticationEnabled interface{} `field:"optional" json:"iamDatabaseAuthenticationEnabled" yaml:"iamDatabaseAuthenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#iam_roles TfCluster#iam_roles}.
	// Experimental.
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#id TfCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#iops TfCluster#iops}.
	// Experimental.
	Iops *float64 `field:"optional" json:"iops" yaml:"iops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#kms_key_id TfCluster#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#manage_master_user_password TfCluster#manage_master_user_password}.
	// Experimental.
	ManageMasterUserPassword interface{} `field:"optional" json:"manageMasterUserPassword" yaml:"manageMasterUserPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#master_password TfCluster#master_password}.
	// Experimental.
	MasterPassword *string `field:"optional" json:"masterPassword" yaml:"masterPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#master_password_wo TfCluster#master_password_wo}.
	// Experimental.
	MasterPasswordWo *string `field:"optional" json:"masterPasswordWo" yaml:"masterPasswordWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#master_password_wo_version TfCluster#master_password_wo_version}.
	// Experimental.
	MasterPasswordWoVersion *float64 `field:"optional" json:"masterPasswordWoVersion" yaml:"masterPasswordWoVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#master_username TfCluster#master_username}.
	// Experimental.
	MasterUsername *string `field:"optional" json:"masterUsername" yaml:"masterUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#master_user_secret_kms_key_id TfCluster#master_user_secret_kms_key_id}.
	// Experimental.
	MasterUserSecretKmsKeyId *string `field:"optional" json:"masterUserSecretKmsKeyId" yaml:"masterUserSecretKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#monitoring_interval TfCluster#monitoring_interval}.
	// Experimental.
	MonitoringInterval *float64 `field:"optional" json:"monitoringInterval" yaml:"monitoringInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#monitoring_role_arn TfCluster#monitoring_role_arn}.
	// Experimental.
	MonitoringRoleArn *string `field:"optional" json:"monitoringRoleArn" yaml:"monitoringRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#network_type TfCluster#network_type}.
	// Experimental.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#performance_insights_enabled TfCluster#performance_insights_enabled}.
	// Experimental.
	PerformanceInsightsEnabled interface{} `field:"optional" json:"performanceInsightsEnabled" yaml:"performanceInsightsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#performance_insights_kms_key_id TfCluster#performance_insights_kms_key_id}.
	// Experimental.
	PerformanceInsightsKmsKeyId *string `field:"optional" json:"performanceInsightsKmsKeyId" yaml:"performanceInsightsKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#performance_insights_retention_period TfCluster#performance_insights_retention_period}.
	// Experimental.
	PerformanceInsightsRetentionPeriod *float64 `field:"optional" json:"performanceInsightsRetentionPeriod" yaml:"performanceInsightsRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#port TfCluster#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#preferred_backup_window TfCluster#preferred_backup_window}.
	// Experimental.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#preferred_maintenance_window TfCluster#preferred_maintenance_window}.
	// Experimental.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#region TfCluster#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#replication_source_identifier TfCluster#replication_source_identifier}.
	// Experimental.
	ReplicationSourceIdentifier *string `field:"optional" json:"replicationSourceIdentifier" yaml:"replicationSourceIdentifier"`
	// restore_to_point_in_time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#restore_to_point_in_time TfCluster#restore_to_point_in_time}
	// Experimental.
	RestoreToPointInTime *TfCluster_RestoreToPointInTimeProperty `field:"optional" json:"restoreToPointInTime" yaml:"restoreToPointInTime"`
	// s3_import block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#s3_import TfCluster#s3_import}
	// Experimental.
	S3Import *TfCluster_S3ImportProperty `field:"optional" json:"s3Import" yaml:"s3Import"`
	// scaling_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#scaling_configuration TfCluster#scaling_configuration}
	// Experimental.
	ScalingConfiguration *TfCluster_ScalingConfigurationProperty `field:"optional" json:"scalingConfiguration" yaml:"scalingConfiguration"`
	// serverlessv2_scaling_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#serverlessv2_scaling_configuration TfCluster#serverlessv2_scaling_configuration}
	// Experimental.
	Serverlessv2ScalingConfiguration *TfCluster_Serverlessv2ScalingConfigurationProperty `field:"optional" json:"serverlessv2ScalingConfiguration" yaml:"serverlessv2ScalingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#skip_final_snapshot TfCluster#skip_final_snapshot}.
	// Experimental.
	SkipFinalSnapshot interface{} `field:"optional" json:"skipFinalSnapshot" yaml:"skipFinalSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#snapshot_identifier TfCluster#snapshot_identifier}.
	// Experimental.
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#source_region TfCluster#source_region}.
	// Experimental.
	SourceRegion *string `field:"optional" json:"sourceRegion" yaml:"sourceRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#storage_encrypted TfCluster#storage_encrypted}.
	// Experimental.
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#storage_type TfCluster#storage_type}.
	// Experimental.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#tags TfCluster#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#tags_all TfCluster#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#timeouts TfCluster#timeouts}
	// Experimental.
	Timeouts *TfCluster_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/rds_cluster#vpc_security_group_ids TfCluster#vpc_security_group_ids}.
	// Experimental.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

