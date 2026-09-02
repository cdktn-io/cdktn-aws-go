package awsredshift

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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#cluster_identifier TfCluster#cluster_identifier}.
	// Experimental.
	ClusterIdentifier *string `field:"required" json:"clusterIdentifier" yaml:"clusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#node_type TfCluster#node_type}.
	// Experimental.
	NodeType *string `field:"required" json:"nodeType" yaml:"nodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#allow_version_upgrade TfCluster#allow_version_upgrade}.
	// Experimental.
	AllowVersionUpgrade interface{} `field:"optional" json:"allowVersionUpgrade" yaml:"allowVersionUpgrade"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#apply_immediately TfCluster#apply_immediately}.
	// Experimental.
	ApplyImmediately interface{} `field:"optional" json:"applyImmediately" yaml:"applyImmediately"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#aqua_configuration_status TfCluster#aqua_configuration_status}.
	// Experimental.
	AquaConfigurationStatus *string `field:"optional" json:"aquaConfigurationStatus" yaml:"aquaConfigurationStatus"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#automated_snapshot_retention_period TfCluster#automated_snapshot_retention_period}.
	// Experimental.
	AutomatedSnapshotRetentionPeriod *float64 `field:"optional" json:"automatedSnapshotRetentionPeriod" yaml:"automatedSnapshotRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#availability_zone TfCluster#availability_zone}.
	// Experimental.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#availability_zone_relocation_enabled TfCluster#availability_zone_relocation_enabled}.
	// Experimental.
	AvailabilityZoneRelocationEnabled interface{} `field:"optional" json:"availabilityZoneRelocationEnabled" yaml:"availabilityZoneRelocationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#cluster_parameter_group_name TfCluster#cluster_parameter_group_name}.
	// Experimental.
	ClusterParameterGroupName *string `field:"optional" json:"clusterParameterGroupName" yaml:"clusterParameterGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#cluster_subnet_group_name TfCluster#cluster_subnet_group_name}.
	// Experimental.
	ClusterSubnetGroupName *string `field:"optional" json:"clusterSubnetGroupName" yaml:"clusterSubnetGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#cluster_type TfCluster#cluster_type}.
	// Experimental.
	ClusterType *string `field:"optional" json:"clusterType" yaml:"clusterType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#cluster_version TfCluster#cluster_version}.
	// Experimental.
	ClusterVersion *string `field:"optional" json:"clusterVersion" yaml:"clusterVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#database_name TfCluster#database_name}.
	// Experimental.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#default_iam_role_arn TfCluster#default_iam_role_arn}.
	// Experimental.
	DefaultIamRoleArn *string `field:"optional" json:"defaultIamRoleArn" yaml:"defaultIamRoleArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#elastic_ip TfCluster#elastic_ip}.
	// Experimental.
	ElasticIp *string `field:"optional" json:"elasticIp" yaml:"elasticIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#encrypted TfCluster#encrypted}.
	// Experimental.
	Encrypted *string `field:"optional" json:"encrypted" yaml:"encrypted"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#enhanced_vpc_routing TfCluster#enhanced_vpc_routing}.
	// Experimental.
	EnhancedVpcRouting interface{} `field:"optional" json:"enhancedVpcRouting" yaml:"enhancedVpcRouting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#final_snapshot_identifier TfCluster#final_snapshot_identifier}.
	// Experimental.
	FinalSnapshotIdentifier *string `field:"optional" json:"finalSnapshotIdentifier" yaml:"finalSnapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#iam_roles TfCluster#iam_roles}.
	// Experimental.
	IamRoles *[]*string `field:"optional" json:"iamRoles" yaml:"iamRoles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#id TfCluster#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#kms_key_id TfCluster#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#maintenance_track_name TfCluster#maintenance_track_name}.
	// Experimental.
	MaintenanceTrackName *string `field:"optional" json:"maintenanceTrackName" yaml:"maintenanceTrackName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#manage_master_password TfCluster#manage_master_password}.
	// Experimental.
	ManageMasterPassword interface{} `field:"optional" json:"manageMasterPassword" yaml:"manageMasterPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#manual_snapshot_retention_period TfCluster#manual_snapshot_retention_period}.
	// Experimental.
	ManualSnapshotRetentionPeriod *float64 `field:"optional" json:"manualSnapshotRetentionPeriod" yaml:"manualSnapshotRetentionPeriod"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#master_password TfCluster#master_password}.
	// Experimental.
	MasterPassword *string `field:"optional" json:"masterPassword" yaml:"masterPassword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#master_password_secret_kms_key_id TfCluster#master_password_secret_kms_key_id}.
	// Experimental.
	MasterPasswordSecretKmsKeyId *string `field:"optional" json:"masterPasswordSecretKmsKeyId" yaml:"masterPasswordSecretKmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#master_password_wo TfCluster#master_password_wo}.
	// Experimental.
	MasterPasswordWo *string `field:"optional" json:"masterPasswordWo" yaml:"masterPasswordWo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#master_password_wo_version TfCluster#master_password_wo_version}.
	// Experimental.
	MasterPasswordWoVersion *float64 `field:"optional" json:"masterPasswordWoVersion" yaml:"masterPasswordWoVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#master_username TfCluster#master_username}.
	// Experimental.
	MasterUsername *string `field:"optional" json:"masterUsername" yaml:"masterUsername"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#multi_az TfCluster#multi_az}.
	// Experimental.
	MultiAz interface{} `field:"optional" json:"multiAz" yaml:"multiAz"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#number_of_nodes TfCluster#number_of_nodes}.
	// Experimental.
	NumberOfNodes *float64 `field:"optional" json:"numberOfNodes" yaml:"numberOfNodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#owner_account TfCluster#owner_account}.
	// Experimental.
	OwnerAccount *string `field:"optional" json:"ownerAccount" yaml:"ownerAccount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#port TfCluster#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#preferred_maintenance_window TfCluster#preferred_maintenance_window}.
	// Experimental.
	PreferredMaintenanceWindow *string `field:"optional" json:"preferredMaintenanceWindow" yaml:"preferredMaintenanceWindow"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#publicly_accessible TfCluster#publicly_accessible}.
	// Experimental.
	PubliclyAccessible interface{} `field:"optional" json:"publiclyAccessible" yaml:"publiclyAccessible"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#region TfCluster#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#skip_final_snapshot TfCluster#skip_final_snapshot}.
	// Experimental.
	SkipFinalSnapshot interface{} `field:"optional" json:"skipFinalSnapshot" yaml:"skipFinalSnapshot"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#snapshot_arn TfCluster#snapshot_arn}.
	// Experimental.
	SnapshotArn *string `field:"optional" json:"snapshotArn" yaml:"snapshotArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#snapshot_cluster_identifier TfCluster#snapshot_cluster_identifier}.
	// Experimental.
	SnapshotClusterIdentifier *string `field:"optional" json:"snapshotClusterIdentifier" yaml:"snapshotClusterIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#snapshot_identifier TfCluster#snapshot_identifier}.
	// Experimental.
	SnapshotIdentifier *string `field:"optional" json:"snapshotIdentifier" yaml:"snapshotIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#tags TfCluster#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#tags_all TfCluster#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#timeouts TfCluster#timeouts}
	// Experimental.
	Timeouts *TfCluster_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/redshift_cluster#vpc_security_group_ids TfCluster#vpc_security_group_ids}.
	// Experimental.
	VpcSecurityGroupIds *[]*string `field:"optional" json:"vpcSecurityGroupIds" yaml:"vpcSecurityGroupIds"`
}

