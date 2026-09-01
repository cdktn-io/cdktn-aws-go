package awsworkspacesweb


// Experimental.
type AwsWorkspaceswebSessionLogger_S3Property struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_session_logger#bucket AwsWorkspaceswebSessionLogger#bucket}.
	// Experimental.
	Bucket *string `field:"required" json:"bucket" yaml:"bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_session_logger#folder_structure AwsWorkspaceswebSessionLogger#folder_structure}.
	// Experimental.
	FolderStructure *string `field:"required" json:"folderStructure" yaml:"folderStructure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_session_logger#log_file_format AwsWorkspaceswebSessionLogger#log_file_format}.
	// Experimental.
	LogFileFormat *string `field:"required" json:"logFileFormat" yaml:"logFileFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_session_logger#bucket_owner AwsWorkspaceswebSessionLogger#bucket_owner}.
	// Experimental.
	BucketOwner *string `field:"optional" json:"bucketOwner" yaml:"bucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspacesweb_session_logger#key_prefix AwsWorkspaceswebSessionLogger#key_prefix}.
	// Experimental.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

