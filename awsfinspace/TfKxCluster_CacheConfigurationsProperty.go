package awsfinspace


// Experimental.
type TfKxCluster_CacheConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#cache_type TfKxCluster#cache_type}.
	// Experimental.
	CacheType *string `field:"required" json:"cacheType" yaml:"cacheType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#db_paths TfKxCluster#db_paths}.
	// Experimental.
	DbPaths *[]*string `field:"optional" json:"dbPaths" yaml:"dbPaths"`
}

