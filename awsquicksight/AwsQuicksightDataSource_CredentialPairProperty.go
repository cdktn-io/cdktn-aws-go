package awsquicksight


// Experimental.
type AwsQuicksightDataSource_CredentialPairProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#password AwsQuicksightDataSource#password}.
	// Experimental.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_source#username AwsQuicksightDataSource#username}.
	// Experimental.
	Username *string `field:"required" json:"username" yaml:"username"`
}

