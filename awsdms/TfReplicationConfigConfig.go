package awsdms

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReplicationConfigConfig struct {
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
	// compute_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#compute_config TfReplicationConfig#compute_config}
	// Experimental.
	ComputeConfig *TfReplicationConfig_ComputeConfigProperty `field:"required" json:"computeConfig" yaml:"computeConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#replication_config_identifier TfReplicationConfig#replication_config_identifier}.
	// Experimental.
	ReplicationConfigIdentifier *string `field:"required" json:"replicationConfigIdentifier" yaml:"replicationConfigIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#replication_type TfReplicationConfig#replication_type}.
	// Experimental.
	ReplicationType *string `field:"required" json:"replicationType" yaml:"replicationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#source_endpoint_arn TfReplicationConfig#source_endpoint_arn}.
	// Experimental.
	SourceEndpointArn *string `field:"required" json:"sourceEndpointArn" yaml:"sourceEndpointArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#table_mappings TfReplicationConfig#table_mappings}.
	// Experimental.
	TableMappings *string `field:"required" json:"tableMappings" yaml:"tableMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#target_endpoint_arn TfReplicationConfig#target_endpoint_arn}.
	// Experimental.
	TargetEndpointArn *string `field:"required" json:"targetEndpointArn" yaml:"targetEndpointArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#id TfReplicationConfig#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#region TfReplicationConfig#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#replication_settings TfReplicationConfig#replication_settings}.
	// Experimental.
	ReplicationSettings *string `field:"optional" json:"replicationSettings" yaml:"replicationSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#resource_identifier TfReplicationConfig#resource_identifier}.
	// Experimental.
	ResourceIdentifier *string `field:"optional" json:"resourceIdentifier" yaml:"resourceIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#start_replication TfReplicationConfig#start_replication}.
	// Experimental.
	StartReplication interface{} `field:"optional" json:"startReplication" yaml:"startReplication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#supplemental_settings TfReplicationConfig#supplemental_settings}.
	// Experimental.
	SupplementalSettings *string `field:"optional" json:"supplementalSettings" yaml:"supplementalSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#tags TfReplicationConfig#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#tags_all TfReplicationConfig#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/dms_replication_config#timeouts TfReplicationConfig#timeouts}
	// Experimental.
	Timeouts *TfReplicationConfig_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

