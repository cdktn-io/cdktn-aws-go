package awsquicksight


// Experimental.
type TfDataSource_CredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#copy_source_arn TfDataSource#copy_source_arn}.
	// Experimental.
	CopySourceArn *string `field:"optional" json:"copySourceArn" yaml:"copySourceArn"`
	// credential_pair block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#credential_pair TfDataSource#credential_pair}
	// Experimental.
	CredentialPair *TfDataSource_CredentialPairProperty `field:"optional" json:"credentialPair" yaml:"credentialPair"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#secret_arn TfDataSource#secret_arn}.
	// Experimental.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

