package awsdrs

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReplicationConfigurationTemplateConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#associate_default_security_group TfReplicationConfigurationTemplate#associate_default_security_group}.
	// Experimental.
	AssociateDefaultSecurityGroup interface{} `field:"required" json:"associateDefaultSecurityGroup" yaml:"associateDefaultSecurityGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#bandwidth_throttling TfReplicationConfigurationTemplate#bandwidth_throttling}.
	// Experimental.
	BandwidthThrottling *float64 `field:"required" json:"bandwidthThrottling" yaml:"bandwidthThrottling"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#create_public_ip TfReplicationConfigurationTemplate#create_public_ip}.
	// Experimental.
	CreatePublicIp interface{} `field:"required" json:"createPublicIp" yaml:"createPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#data_plane_routing TfReplicationConfigurationTemplate#data_plane_routing}.
	// Experimental.
	DataPlaneRouting *string `field:"required" json:"dataPlaneRouting" yaml:"dataPlaneRouting"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#default_large_staging_disk_type TfReplicationConfigurationTemplate#default_large_staging_disk_type}.
	// Experimental.
	DefaultLargeStagingDiskType *string `field:"required" json:"defaultLargeStagingDiskType" yaml:"defaultLargeStagingDiskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#ebs_encryption TfReplicationConfigurationTemplate#ebs_encryption}.
	// Experimental.
	EbsEncryption *string `field:"required" json:"ebsEncryption" yaml:"ebsEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#replication_server_instance_type TfReplicationConfigurationTemplate#replication_server_instance_type}.
	// Experimental.
	ReplicationServerInstanceType *string `field:"required" json:"replicationServerInstanceType" yaml:"replicationServerInstanceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#replication_servers_security_groups_ids TfReplicationConfigurationTemplate#replication_servers_security_groups_ids}.
	// Experimental.
	ReplicationServersSecurityGroupsIds *[]*string `field:"required" json:"replicationServersSecurityGroupsIds" yaml:"replicationServersSecurityGroupsIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#staging_area_subnet_id TfReplicationConfigurationTemplate#staging_area_subnet_id}.
	// Experimental.
	StagingAreaSubnetId *string `field:"required" json:"stagingAreaSubnetId" yaml:"stagingAreaSubnetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#staging_area_tags TfReplicationConfigurationTemplate#staging_area_tags}.
	// Experimental.
	StagingAreaTags *map[string]*string `field:"required" json:"stagingAreaTags" yaml:"stagingAreaTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#use_dedicated_replication_server TfReplicationConfigurationTemplate#use_dedicated_replication_server}.
	// Experimental.
	UseDedicatedReplicationServer interface{} `field:"required" json:"useDedicatedReplicationServer" yaml:"useDedicatedReplicationServer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#auto_replicate_new_disks TfReplicationConfigurationTemplate#auto_replicate_new_disks}.
	// Experimental.
	AutoReplicateNewDisks interface{} `field:"optional" json:"autoReplicateNewDisks" yaml:"autoReplicateNewDisks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#ebs_encryption_key_arn TfReplicationConfigurationTemplate#ebs_encryption_key_arn}.
	// Experimental.
	EbsEncryptionKeyArn *string `field:"optional" json:"ebsEncryptionKeyArn" yaml:"ebsEncryptionKeyArn"`
	// pit_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#pit_policy TfReplicationConfigurationTemplate#pit_policy}
	// Experimental.
	PitPolicy interface{} `field:"optional" json:"pitPolicy" yaml:"pitPolicy"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#region TfReplicationConfigurationTemplate#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#tags TfReplicationConfigurationTemplate#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/drs_replication_configuration_template#timeouts TfReplicationConfigurationTemplate#timeouts}
	// Experimental.
	Timeouts *TfReplicationConfigurationTemplate_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

