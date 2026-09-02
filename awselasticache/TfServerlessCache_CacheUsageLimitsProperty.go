package awselasticache


// Experimental.
type TfServerlessCache_CacheUsageLimitsProperty struct {
	// data_storage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#data_storage TfServerlessCache#data_storage}
	// Experimental.
	DataStorage interface{} `field:"optional" json:"dataStorage" yaml:"dataStorage"`
	// ecpu_per_second block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elasticache_serverless_cache#ecpu_per_second TfServerlessCache#ecpu_per_second}
	// Experimental.
	EcpuPerSecond interface{} `field:"optional" json:"ecpuPerSecond" yaml:"ecpuPerSecond"`
}

