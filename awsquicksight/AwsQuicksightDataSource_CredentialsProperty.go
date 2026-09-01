package awsquicksight


// Experimental.
type AwsQuicksightDataSource_CredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#copy_source_arn AwsQuicksightDataSource#copy_source_arn}.
	// Experimental.
	CopySourceArn *string `field:"optional" json:"copySourceArn" yaml:"copySourceArn"`
	// credential_pair block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#credential_pair AwsQuicksightDataSource#credential_pair}
	// Experimental.
	CredentialPair *AwsQuicksightDataSource_CredentialPairProperty `field:"optional" json:"credentialPair" yaml:"credentialPair"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#secret_arn AwsQuicksightDataSource#secret_arn}.
	// Experimental.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
}

