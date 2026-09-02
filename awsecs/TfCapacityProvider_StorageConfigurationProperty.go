package awsecs


// Experimental.
type TfCapacityProvider_StorageConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_capacity_provider#storage_size_gib TfCapacityProvider#storage_size_gib}.
	// Experimental.
	StorageSizeGib *float64 `field:"required" json:"storageSizeGib" yaml:"storageSizeGib"`
}

