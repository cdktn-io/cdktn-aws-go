package awskeyspaces


// Experimental.
type TfTable_EncryptionSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#kms_key_identifier TfTable#kms_key_identifier}.
	// Experimental.
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#type TfTable#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

