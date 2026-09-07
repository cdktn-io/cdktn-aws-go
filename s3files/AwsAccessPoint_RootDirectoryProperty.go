package s3files


// Experimental.
type AwsAccessPoint_RootDirectoryProperty struct {
	// creation_permissions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_access_point#creation_permissions AwsAccessPoint#creation_permissions}
	// Experimental.
	CreationPermissions interface{} `field:"optional" json:"creationPermissions" yaml:"creationPermissions"`
	// Root directory path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/s3files_access_point#path AwsAccessPoint#path}
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

