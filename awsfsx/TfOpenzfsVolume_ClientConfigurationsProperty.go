package awsfsx


// Experimental.
type TfOpenzfsVolume_ClientConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#clients TfOpenzfsVolume#clients}.
	// Experimental.
	Clients *string `field:"required" json:"clients" yaml:"clients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#options TfOpenzfsVolume#options}.
	// Experimental.
	Options *[]*string `field:"required" json:"options" yaml:"options"`
}

