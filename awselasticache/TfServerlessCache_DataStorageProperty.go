package awselasticache


// Experimental.
type TfServerlessCache_DataStorageProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#unit TfServerlessCache#unit}.
	// Experimental.
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#maximum TfServerlessCache#maximum}.
	// Experimental.
	Maximum *float64 `field:"optional" json:"maximum" yaml:"maximum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#minimum TfServerlessCache#minimum}.
	// Experimental.
	Minimum *float64 `field:"optional" json:"minimum" yaml:"minimum"`
}

