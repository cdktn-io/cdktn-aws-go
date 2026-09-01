package awskeyspaces


// Experimental.
type AwsKeyspacesTable_EncryptionSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#kms_key_identifier AwsKeyspacesTable#kms_key_identifier}.
	// Experimental.
	KmsKeyIdentifier *string `field:"optional" json:"kmsKeyIdentifier" yaml:"kmsKeyIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/keyspaces_table#type AwsKeyspacesTable#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

