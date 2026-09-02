package awselasticache

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReplicationGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#description TfReplicationGroup#description}.
	// Experimental.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#replication_group_id TfReplicationGroup#replication_group_id}.
	// Experimental.
	ReplicationGroupId *string `field:"required" json:"replicationGroupId" yaml:"replicationGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#apply_immediately TfReplicationGroup#apply_immediately}.
	// Experimental.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#at_rest_encryption_enabled TfReplicationGroup#at_rest_encryption_enabled}.
	// Experimental.
	AtRestEncryptionEnabled *string `field:"optional" json:"atRestEncryptionEnabled" yaml:"atRestEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#auth_token TfReplicationGroup#auth_token}.
	// Experimental.
	AuthToken *string `field:"optional" json:"authToken" yaml:"authToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#auth_token_update_strategy TfReplicationGroup#auth_token_update_strategy}.
	// Experimental.
	AuthTokenUpdateStrategy *string `field:"optional" json:"authTokenUpdateStrategy" yaml:"authTokenUpdateStrategy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#auth_token_wo TfReplicationGroup#auth_token_wo}.
	// Experimental.
	AuthTokenWo *string `field:"optional" json:"authTokenWo" yaml:"authTokenWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#auth_token_wo_version TfReplicationGroup#auth_token_wo_version}.
	// Experimental.
	AuthTokenWoVersion *float64 `field:"optional" json:"authTokenWoVersion" yaml:"authTokenWoVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#automatic_failover_enabled TfReplicationGroup#automatic_failover_enabled}.
	// Experimental.
	AutomaticFailoverEnabled interface{} `field:"optional" json:"automaticFailoverEnabled" yaml:"automaticFailoverEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#auto_minor_version_upgrade TfReplicationGroup#auto_minor_version_upgrade}.
	// Experimental.
	AutoMinorVersionUpgrade *string `field:"optional" json:"autoMinorVersionUpgrade" yaml:"autoMinorVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#cluster_mode TfReplicationGroup#cluster_mode}.
	// Experimental.
	ClusterMode *string `field:"optional" json:"clusterMode" yaml:"clusterMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#data_tiering_enabled TfReplicationGroup#data_tiering_enabled}.
	// Experimental.
	DataTieringEnabled interface{} `field:"optional" json:"dataTieringEnabled" yaml:"dataTieringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#durability TfReplicationGroup#durability}.
	// Experimental.
	Durability *string `field:"optional" json:"durability" yaml:"durability"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#engine TfReplicationGroup#engine}.
	// Experimental.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#engine_version TfReplicationGroup#engine_version}.
	// Experimental.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#final_snapshot_identifier TfReplicationGroup#final_snapshot_identifier}.
	// Experimental.
	FinalSnapshotIdentifier *string `field:"optional" json:"finalSnapshotIdentifier" yaml:"finalSnapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#global_replication_group_id TfReplicationGroup#global_replication_group_id}.
	// Experimental.
	GlobalReplicationGroupId *string `field:"optional" json:"globalReplicationGroupId" yaml:"globalReplicationGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#id TfReplicationGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#ip_discovery TfReplicationGroup#ip_discovery}.
	// Experimental.
	IpDiscovery *string `field:"optional" json:"ipDiscovery" yaml:"ipDiscovery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#kms_key_id TfReplicationGroup#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// log_delivery_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#log_delivery_configuration TfReplicationGroup#log_delivery_configuration}
	// Experimental.
	LogDeliveryConfiguration interface{} `field:"optional" json:"logDeliveryConfiguration" yaml:"logDeliveryConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#maintenance_window TfReplicationGroup#maintenance_window}.
	// Experimental.
	MaintenanceWindow *string `field:"optional" json:"maintenanceWindow" yaml:"maintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#multi_az_enabled TfReplicationGroup#multi_az_enabled}.
	// Experimental.
	MultiAzEnabled interface{} `field:"optional" json:"multiAzEnabled" yaml:"multiAzEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#network_type TfReplicationGroup#network_type}.
	// Experimental.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// node_group_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#node_group_configuration TfReplicationGroup#node_group_configuration}
	// Experimental.
	NodeGroupConfiguration interface{} `field:"optional" json:"nodeGroupConfiguration" yaml:"nodeGroupConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#node_type TfReplicationGroup#node_type}.
	// Experimental.
	NodeType *string `field:"optional" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#notification_topic_arn TfReplicationGroup#notification_topic_arn}.
	// Experimental.
	NotificationTopicArn *string `field:"optional" json:"notificationTopicArn" yaml:"notificationTopicArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#num_cache_clusters TfReplicationGroup#num_cache_clusters}.
	// Experimental.
	NumCacheClusters *float64 `field:"optional" json:"numCacheClusters" yaml:"numCacheClusters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#num_node_groups TfReplicationGroup#num_node_groups}.
	// Experimental.
	NumNodeGroups *float64 `field:"optional" json:"numNodeGroups" yaml:"numNodeGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#parameter_group_name TfReplicationGroup#parameter_group_name}.
	// Experimental.
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#port TfReplicationGroup#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#preferred_cache_cluster_azs TfReplicationGroup#preferred_cache_cluster_azs}.
	// Experimental.
	PreferredCacheClusterAzs *[]*string `field:"optional" json:"preferredCacheClusterAzs" yaml:"preferredCacheClusterAzs"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#region TfReplicationGroup#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#replicas_per_node_group TfReplicationGroup#replicas_per_node_group}.
	// Experimental.
	ReplicasPerNodeGroup *float64 `field:"optional" json:"replicasPerNodeGroup" yaml:"replicasPerNodeGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#security_group_ids TfReplicationGroup#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#security_group_names TfReplicationGroup#security_group_names}.
	// Experimental.
	SecurityGroupNames *[]*string `field:"optional" json:"securityGroupNames" yaml:"securityGroupNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#snapshot_arns TfReplicationGroup#snapshot_arns}.
	// Experimental.
	SnapshotArns *[]*string `field:"optional" json:"snapshotArns" yaml:"snapshotArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#snapshot_name TfReplicationGroup#snapshot_name}.
	// Experimental.
	SnapshotName *string `field:"optional" json:"snapshotName" yaml:"snapshotName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#snapshot_retention_limit TfReplicationGroup#snapshot_retention_limit}.
	// Experimental.
	SnapshotRetentionLimit *float64 `field:"optional" json:"snapshotRetentionLimit" yaml:"snapshotRetentionLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#snapshot_window TfReplicationGroup#snapshot_window}.
	// Experimental.
	SnapshotWindow *string `field:"optional" json:"snapshotWindow" yaml:"snapshotWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#subnet_group_name TfReplicationGroup#subnet_group_name}.
	// Experimental.
	SubnetGroupName *string `field:"optional" json:"subnetGroupName" yaml:"subnetGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#tags TfReplicationGroup#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#tags_all TfReplicationGroup#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#timeouts TfReplicationGroup#timeouts}
	// Experimental.
	Timeouts *TfReplicationGroup_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#transit_encryption_enabled TfReplicationGroup#transit_encryption_enabled}.
	// Experimental.
	TransitEncryptionEnabled interface{} `field:"optional" json:"transitEncryptionEnabled" yaml:"transitEncryptionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#transit_encryption_mode TfReplicationGroup#transit_encryption_mode}.
	// Experimental.
	TransitEncryptionMode *string `field:"optional" json:"transitEncryptionMode" yaml:"transitEncryptionMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_replication_group#user_group_ids TfReplicationGroup#user_group_ids}.
	// Experimental.
	UserGroupIds *[]*string `field:"optional" json:"userGroupIds" yaml:"userGroupIds"`
}

