package awselasticache

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReservedCacheNodeConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_reserved_cache_node#reserved_cache_nodes_offering_id TfReservedCacheNode#reserved_cache_nodes_offering_id}.
	// Experimental.
	ReservedCacheNodesOfferingId *string `field:"required" json:"reservedCacheNodesOfferingId" yaml:"reservedCacheNodesOfferingId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_reserved_cache_node#cache_node_count TfReservedCacheNode#cache_node_count}.
	// Experimental.
	CacheNodeCount *float64 `field:"optional" json:"cacheNodeCount" yaml:"cacheNodeCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_reserved_cache_node#id TfReservedCacheNode#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_reserved_cache_node#region TfReservedCacheNode#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_reserved_cache_node#tags TfReservedCacheNode#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_reserved_cache_node#timeouts TfReservedCacheNode#timeouts}
	// Experimental.
	Timeouts *TfReservedCacheNode_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

