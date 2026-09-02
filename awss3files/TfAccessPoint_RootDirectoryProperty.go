package awss3files


// Experimental.
type TfAccessPoint_RootDirectoryProperty struct {
	// creation_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_access_point#creation_permissions TfAccessPoint#creation_permissions}
	// Experimental.
	CreationPermissions interface{} `field:"optional" json:"creationPermissions" yaml:"creationPermissions"`
	// Root directory path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_access_point#path TfAccessPoint#path}
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

