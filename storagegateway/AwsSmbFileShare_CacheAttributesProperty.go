package storagegateway


// Experimental.
type AwsSmbFileShare_CacheAttributesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/storagegateway_smb_file_share#cache_stale_timeout_in_seconds AwsSmbFileShare#cache_stale_timeout_in_seconds}.
	// Experimental.
	CacheStaleTimeoutInSeconds *float64 `field:"optional" json:"cacheStaleTimeoutInSeconds" yaml:"cacheStaleTimeoutInSeconds"`
}

