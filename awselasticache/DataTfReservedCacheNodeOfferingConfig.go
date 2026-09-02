package awselasticache

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfReservedCacheNodeOfferingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/elasticache_reserved_cache_node_offering#cache_node_type DataTfReservedCacheNodeOffering#cache_node_type}.
	// Experimental.
	CacheNodeType *string `field:"required" json:"cacheNodeType" yaml:"cacheNodeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/elasticache_reserved_cache_node_offering#duration DataTfReservedCacheNodeOffering#duration}.
	// Experimental.
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/elasticache_reserved_cache_node_offering#offering_type DataTfReservedCacheNodeOffering#offering_type}.
	// Experimental.
	OfferingType *string `field:"required" json:"offeringType" yaml:"offeringType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/elasticache_reserved_cache_node_offering#product_description DataTfReservedCacheNodeOffering#product_description}.
	// Experimental.
	ProductDescription *string `field:"required" json:"productDescription" yaml:"productDescription"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/elasticache_reserved_cache_node_offering#region DataTfReservedCacheNodeOffering#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

