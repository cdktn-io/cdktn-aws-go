package awsdms

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReplicationTaskConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#migration_type TfReplicationTask#migration_type}.
	// Experimental.
	MigrationType *string `field:"required" json:"migrationType" yaml:"migrationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#replication_instance_arn TfReplicationTask#replication_instance_arn}.
	// Experimental.
	ReplicationInstanceArn *string `field:"required" json:"replicationInstanceArn" yaml:"replicationInstanceArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#replication_task_id TfReplicationTask#replication_task_id}.
	// Experimental.
	ReplicationTaskId *string `field:"required" json:"replicationTaskId" yaml:"replicationTaskId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#source_endpoint_arn TfReplicationTask#source_endpoint_arn}.
	// Experimental.
	SourceEndpointArn *string `field:"required" json:"sourceEndpointArn" yaml:"sourceEndpointArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#table_mappings TfReplicationTask#table_mappings}.
	// Experimental.
	TableMappings *string `field:"required" json:"tableMappings" yaml:"tableMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#target_endpoint_arn TfReplicationTask#target_endpoint_arn}.
	// Experimental.
	TargetEndpointArn *string `field:"required" json:"targetEndpointArn" yaml:"targetEndpointArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#cdc_start_position TfReplicationTask#cdc_start_position}.
	// Experimental.
	CdcStartPosition *string `field:"optional" json:"cdcStartPosition" yaml:"cdcStartPosition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#cdc_start_time TfReplicationTask#cdc_start_time}.
	// Experimental.
	CdcStartTime *string `field:"optional" json:"cdcStartTime" yaml:"cdcStartTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#id TfReplicationTask#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#region TfReplicationTask#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#replication_task_settings TfReplicationTask#replication_task_settings}.
	// Experimental.
	ReplicationTaskSettings *string `field:"optional" json:"replicationTaskSettings" yaml:"replicationTaskSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#resource_identifier TfReplicationTask#resource_identifier}.
	// Experimental.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#start_replication_task TfReplicationTask#start_replication_task}.
	// Experimental.
	StartReplicationTask interface{} `field:"optional" json:"startReplicationTask" yaml:"startReplicationTask"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#tags TfReplicationTask#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_task#tags_all TfReplicationTask#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
}

