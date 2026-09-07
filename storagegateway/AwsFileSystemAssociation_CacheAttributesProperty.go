package storagegateway


// Experimental.
type AwsFileSystemAssociation_CacheAttributesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_file_system_association#cache_stale_timeout_in_seconds AwsFileSystemAssociation#cache_stale_timeout_in_seconds}.
	// Experimental.
	CacheStaleTimeoutInSeconds *float64 `field:"optional" json:"cacheStaleTimeoutInSeconds" yaml:"cacheStaleTimeoutInSeconds"`
}

