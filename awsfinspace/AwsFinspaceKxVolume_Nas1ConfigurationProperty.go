package awsfinspace


// Experimental.
type AwsFinspaceKxVolume_Nas1ConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_volume#size AwsFinspaceKxVolume#size}.
	// Experimental.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/finspace_kx_volume#type AwsFinspaceKxVolume#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

