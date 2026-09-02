package awscodegurureviewer


// Experimental.
type TfRepositoryAssociation_KmsKeyDetailsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#encryption_option TfRepositoryAssociation#encryption_option}.
	// Experimental.
	EncryptionOption *string `field:"optional" json:"encryptionOption" yaml:"encryptionOption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/codegurureviewer_repository_association#kms_key_id TfRepositoryAssociation#kms_key_id}.
	// Experimental.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

