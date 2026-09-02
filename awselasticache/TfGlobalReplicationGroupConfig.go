package awselasticache

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGlobalReplicationGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#global_replication_group_id_suffix TfGlobalReplicationGroup#global_replication_group_id_suffix}.
	// Experimental.
	GlobalReplicationGroupIdSuffix *string `field:"required" json:"globalReplicationGroupIdSuffix" yaml:"globalReplicationGroupIdSuffix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#primary_replication_group_id TfGlobalReplicationGroup#primary_replication_group_id}.
	// Experimental.
	PrimaryReplicationGroupId *string `field:"required" json:"primaryReplicationGroupId" yaml:"primaryReplicationGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#automatic_failover_enabled TfGlobalReplicationGroup#automatic_failover_enabled}.
	// Experimental.
	AutomaticFailoverEnabled interface{} `field:"optional" json:"automaticFailoverEnabled" yaml:"automaticFailoverEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#cache_node_type TfGlobalReplicationGroup#cache_node_type}.
	// Experimental.
	CacheNodeType *string `field:"optional" json:"cacheNodeType" yaml:"cacheNodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#engine TfGlobalReplicationGroup#engine}.
	// Experimental.
	Engine *string `field:"optional" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#engine_version TfGlobalReplicationGroup#engine_version}.
	// Experimental.
	EngineVersion *string `field:"optional" json:"engineVersion" yaml:"engineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#global_replication_group_description TfGlobalReplicationGroup#global_replication_group_description}.
	// Experimental.
	GlobalReplicationGroupDescription *string `field:"optional" json:"globalReplicationGroupDescription" yaml:"globalReplicationGroupDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#id TfGlobalReplicationGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#num_node_groups TfGlobalReplicationGroup#num_node_groups}.
	// Experimental.
	NumNodeGroups *float64 `field:"optional" json:"numNodeGroups" yaml:"numNodeGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#parameter_group_name TfGlobalReplicationGroup#parameter_group_name}.
	// Experimental.
	ParameterGroupName *string `field:"optional" json:"parameterGroupName" yaml:"parameterGroupName"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#region TfGlobalReplicationGroup#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_global_replication_group#timeouts TfGlobalReplicationGroup#timeouts}
	// Experimental.
	Timeouts *TfGlobalReplicationGroup_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

