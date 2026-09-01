package awsefs


// Experimental.
type AwsEfsAccessPoint_RootDirectoryProperty struct {
	// creation_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#creation_info AwsEfsAccessPoint#creation_info}
	// Experimental.
	CreationInfo *AwsEfsAccessPoint_CreationInfoProperty `field:"optional" json:"creationInfo" yaml:"creationInfo"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_access_point#path AwsEfsAccessPoint#path}.
	// Experimental.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

