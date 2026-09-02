package awsefs


// Experimental.
type TfAccessPoint_RootDirectoryProperty struct {
	// creation_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#creation_info TfAccessPoint#creation_info}
	// Experimental.
	CreationInfo *TfAccessPoint_CreationInfoProperty `field:"optional" json:"creationInfo" yaml:"creationInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#path TfAccessPoint#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

