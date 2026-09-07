package efs


// Experimental.
type AwsAccessPoint_RootDirectoryProperty struct {
	// creation_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#creation_info AwsAccessPoint#creation_info}
	// Experimental.
	CreationInfo *AwsAccessPoint_CreationInfoProperty `field:"optional" json:"creationInfo" yaml:"creationInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#path AwsAccessPoint#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

