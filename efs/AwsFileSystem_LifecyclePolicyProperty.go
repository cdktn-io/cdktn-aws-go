package efs


// Experimental.
type AwsFileSystem_LifecyclePolicyProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_file_system#transition_to_archive AwsFileSystem#transition_to_archive}.
	// Experimental.
	TransitionToArchive *string `field:"optional" json:"transitionToArchive" yaml:"transitionToArchive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_file_system#transition_to_ia AwsFileSystem#transition_to_ia}.
	// Experimental.
	TransitionToIa *string `field:"optional" json:"transitionToIa" yaml:"transitionToIa"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/efs_file_system#transition_to_primary_storage_class AwsFileSystem#transition_to_primary_storage_class}.
	// Experimental.
	TransitionToPrimaryStorageClass *string `field:"optional" json:"transitionToPrimaryStorageClass" yaml:"transitionToPrimaryStorageClass"`
}

