package awsneptune

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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#allow_major_version_upgrade TfCluster#allow_major_version_upgrade}.
	// Experimental.
	AllowMajorVersionUpgrade interface{} `field:"optional" json:"allowMajorVersionUpgrade" yaml:"allowMajorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#apply_immediately TfCluster#apply_immediately}.
	// Experimental.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#availability_zones TfCluster#availability_zones}.
	// Experimental.
	AvailabilityZones *[]*string `field:"optional" json:"availabilityZones" yaml:"availabilityZones"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#backup_retention_period TfCluster#backup_retention_period}.
	// Experimental.
	BackupRetentionPeriod *float64 `field:"optional" json:"backupRetentionPeriod" yaml:"backupRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#cluster_identifier TfCluster#cluster_identifier}.
	// Experimental.
	ClusterIdentifier *string `field:"optional" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#cluster_identifier_prefix TfCluster#cluster_identifier_prefix}.
	// Experimental.
	ClusterIdentifierPrefix *string `field:"optional" json:"clusterIdentifierPrefix" yaml:"clusterIdentifierPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#copy_tags_to_snapshot TfCluster#copy_tags_to_snapshot}.
	// Experimental.
	CopyTagsToSnapshot interface{} `field:"optional" json:"copyTagsToSnapshot" yaml:"copyTagsToSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#deletion_protection TfCluster#deletion_protection}.
	// Experimental.
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#enable_cloudwatch_logs_exports TfCluster#enable_cloudwatch_logs_exports}.
	// Experimental.
	EnableCloudwatchLogsExports *[]*string `field:"optional" json:"enableCloudwatchLogsExports" yaml:"enableCloudwatchLogsExports"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#engine TfCluster#engine}.
	// Experimental.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#engine_version TfCluster#engine_version}.
	// Experimental.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#final_snapshot_identifier TfCluster#final_snapshot_identifier}.
	// Experimental.
	FinalSnapshotIdentifier *string `field:"optional" json:"finalSnapshotIdentifier" yaml:"finalSnapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#global_cluster_identifier TfCluster#global_cluster_identifier}.
	// Experimental.
	GlobalClusterIdentifier *string `field:"optional" json:"globalClusterIdentifier" yaml:"globalClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#iam_database_authentication_enabled TfCluster#iam_database_authentication_enabled}.
	// Experimental.
	IamDatabaseAuthenticationEnabled interface{} `field:"optional" json:"iamDatabaseAuthenticationEnabled" yaml:"iamDatabaseAuthenticationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#iam_roles TfCluster#iam_roles}.
	// Experimental.
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#id TfCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#kms_key_arn TfCluster#kms_key_arn}.
	// Experimental.
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#neptune_cluster_parameter_group_name TfCluster#neptune_cluster_parameter_group_name}.
	// Experimental.
	NeptuneClusterParameterGroupName *string `field:"optional" json:"neptuneClusterParameterGroupName" yaml:"neptuneClusterParameterGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#neptune_instance_parameter_group_name TfCluster#neptune_instance_parameter_group_name}.
	// Experimental.
	NeptuneInstanceParameterGroupName *string `field:"optional" json:"neptuneInstanceParameterGroupName" yaml:"neptuneInstanceParameterGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#neptune_subnet_group_name TfCluster#neptune_subnet_group_name}.
	// Experimental.
	NeptuneSubnetGroupName *string `field:"optional" json:"neptuneSubnetGroupName" yaml:"neptuneSubnetGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#port TfCluster#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#preferred_backup_window TfCluster#preferred_backup_window}.
	// Experimental.
	PreferredBackupWindow *string `field:"optional" json:"preferredBackupWindow" yaml:"preferredBackupWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#preferred_maintenance_window TfCluster#preferred_maintenance_window}.
	// Experimental.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#region TfCluster#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#replication_source_identifier TfCluster#replication_source_identifier}.
	// Experimental.
	ReplicationSourceIdentifier *string `field:"optional" json:"replicationSourceIdentifier" yaml:"replicationSourceIdentifier"`
	// serverless_v2_scaling_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#serverless_v2_scaling_configuration TfCluster#serverless_v2_scaling_configuration}
	// Experimental.
	ServerlessV2ScalingConfiguration *TfCluster_ServerlessV2ScalingConfigurationProperty `field:"optional" json:"serverlessV2ScalingConfiguration" yaml:"serverlessV2ScalingConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#skip_final_snapshot TfCluster#skip_final_snapshot}.
	// Experimental.
	SkipFinalSnapshot interface{} `field:"optional" json:"skipFinalSnapshot" yaml:"skipFinalSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#snapshot_identifier TfCluster#snapshot_identifier}.
	// Experimental.
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#storage_encrypted TfCluster#storage_encrypted}.
	// Experimental.
	StorageEncrypted interface{} `field:"optional" json:"storageEncrypted" yaml:"storageEncrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#storage_type TfCluster#storage_type}.
	// Experimental.
	StorageType *string `field:"optional" json:"storageType" yaml:"storageType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#tags TfCluster#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#tags_all TfCluster#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#timeouts TfCluster#timeouts}
	// Experimental.
	Timeouts *TfCluster_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/neptune_cluster#vpc_security_group_ids TfCluster#vpc_security_group_ids}.
	// Experimental.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

