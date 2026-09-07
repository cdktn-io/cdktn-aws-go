package storagegateway


// Experimental.
type AwsNfsFileShare_CacheAttributesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_nfs_file_share#cache_stale_timeout_in_seconds AwsNfsFileShare#cache_stale_timeout_in_seconds}.
	// Experimental.
	CacheStaleTimeoutInSeconds *float64 `field:"optional" json:"cacheStaleTimeoutInSeconds" yaml:"cacheStaleTimeoutInSeconds"`
}

