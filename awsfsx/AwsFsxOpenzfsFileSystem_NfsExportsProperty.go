package awsfsx


// Experimental.
type AwsFsxOpenzfsFileSystem_NfsExportsProperty struct {
	// client_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/fsx_openzfs_file_system#client_configurations AwsFsxOpenzfsFileSystem#client_configurations}
	// Experimental.
	ClientConfigurations interface{} `field:"required" json:"clientConfigurations" yaml:"clientConfigurations"`
}

