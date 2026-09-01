package awsefs


// Experimental.
type AwsEfsFileSystem_LifecyclePolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_file_system#transition_to_archive AwsEfsFileSystem#transition_to_archive}.
	// Experimental.
	TransitionToArchive *string `field:"optional" json:"transitionToArchive" yaml:"transitionToArchive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_file_system#transition_to_ia AwsEfsFileSystem#transition_to_ia}.
	// Experimental.
	TransitionToIa *string `field:"optional" json:"transitionToIa" yaml:"transitionToIa"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_file_system#transition_to_primary_storage_class AwsEfsFileSystem#transition_to_primary_storage_class}.
	// Experimental.
	TransitionToPrimaryStorageClass *string `field:"optional" json:"transitionToPrimaryStorageClass" yaml:"transitionToPrimaryStorageClass"`
}

