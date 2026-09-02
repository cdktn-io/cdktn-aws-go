package awscodebuild


// Experimental.
type TfProject_FileSystemLocationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#identifier TfProject#identifier}.
	// Experimental.
	Identifier *string `field:"optional" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#location TfProject#location}.
	// Experimental.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#mount_options TfProject#mount_options}.
	// Experimental.
	MountOptions *string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#mount_point TfProject#mount_point}.
	// Experimental.
	MountPoint *string `field:"optional" json:"mountPoint" yaml:"mountPoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#type TfProject#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

