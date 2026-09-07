package fsx


// Experimental.
type AwsOpenzfsVolume_ClientConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#clients AwsOpenzfsVolume#clients}.
	// Experimental.
	Clients *string `field:"required" json:"clients" yaml:"clients"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_volume#options AwsOpenzfsVolume#options}.
	// Experimental.
	Options *[]*string `field:"required" json:"options" yaml:"options"`
}

