package awselasticache

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfServerlessCacheConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#engine TfServerlessCache#engine}.
	// Experimental.
	Engine *string `field:"required" json:"engine" yaml:"engine"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#name TfServerlessCache#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// cache_usage_limits block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#cache_usage_limits TfServerlessCache#cache_usage_limits}
	// Experimental.
	CacheUsageLimits interface{} `field:"optional" json:"cacheUsageLimits" yaml:"cacheUsageLimits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#daily_snapshot_time TfServerlessCache#daily_snapshot_time}.
	// Experimental.
	DailySnapshotTime *string `field:"optional" json:"dailySnapshotTime" yaml:"dailySnapshotTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#description TfServerlessCache#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#kms_key_id TfServerlessCache#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#major_engine_version TfServerlessCache#major_engine_version}.
	// Experimental.
	MajorEngineVersion *string `field:"optional" json:"majorEngineVersion" yaml:"majorEngineVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#network_type TfServerlessCache#network_type}.
	// Experimental.
	NetworkType *string `field:"optional" json:"networkType" yaml:"networkType"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#region TfServerlessCache#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#security_group_ids TfServerlessCache#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#snapshot_arns_to_restore TfServerlessCache#snapshot_arns_to_restore}.
	// Experimental.
	SnapshotArnsToRestore *[]*string `field:"optional" json:"snapshotArnsToRestore" yaml:"snapshotArnsToRestore"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#snapshot_retention_limit TfServerlessCache#snapshot_retention_limit}.
	// Experimental.
	SnapshotRetentionLimit *float64 `field:"optional" json:"snapshotRetentionLimit" yaml:"snapshotRetentionLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#subnet_ids TfServerlessCache#subnet_ids}.
	// Experimental.
	SubnetIds *[]*string `field:"optional" json:"subnetIds" yaml:"subnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#tags TfServerlessCache#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#timeouts TfServerlessCache#timeouts}
	// Experimental.
	Timeouts *TfServerlessCache_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#user_group_id TfServerlessCache#user_group_id}.
	// Experimental.
	UserGroupId *string `field:"optional" json:"userGroupId" yaml:"userGroupId"`
}

