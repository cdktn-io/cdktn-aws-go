package awscodebuild


// Experimental.
type TfProject_S3LogsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#bucket_owner_access TfProject#bucket_owner_access}.
	// Experimental.
	BucketOwnerAccess *string `field:"optional" json:"bucketOwnerAccess" yaml:"bucketOwnerAccess"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#encryption_disabled TfProject#encryption_disabled}.
	// Experimental.
	EncryptionDisabled interface{} `field:"optional" json:"encryptionDisabled" yaml:"encryptionDisabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#location TfProject#location}.
	// Experimental.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codebuild_project#status TfProject#status}.
	// Experimental.
	Status *string `field:"optional" json:"status" yaml:"status"`
}

