package finspace


// Experimental.
type AwsKxVolume_Nas1ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_volume#size AwsKxVolume#size}.
	// Experimental.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_volume#type AwsKxVolume#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

