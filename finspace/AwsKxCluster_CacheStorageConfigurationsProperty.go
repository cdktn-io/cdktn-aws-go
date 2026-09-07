package finspace


// Experimental.
type AwsKxCluster_CacheStorageConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#size AwsKxCluster#size}.
	// Experimental.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_cluster#type AwsKxCluster#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

