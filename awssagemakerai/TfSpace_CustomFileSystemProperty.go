package awssagemakerai


// Experimental.
type TfSpace_CustomFileSystemProperty struct {
	// efs_file_system block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_space#efs_file_system TfSpace#efs_file_system}
	// Experimental.
	EfsFileSystem *TfSpace_EfsFileSystemProperty `field:"required" json:"efsFileSystem" yaml:"efsFileSystem"`
}

