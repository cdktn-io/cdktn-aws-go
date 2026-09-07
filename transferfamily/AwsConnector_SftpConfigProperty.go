package transferfamily


// Experimental.
type AwsConnector_SftpConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#trusted_host_keys AwsConnector#trusted_host_keys}.
	// Experimental.
	TrustedHostKeys *[]*string `field:"optional" json:"trustedHostKeys" yaml:"trustedHostKeys"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/transfer_connector#user_secret_id AwsConnector#user_secret_id}.
	// Experimental.
	UserSecretId *string `field:"optional" json:"userSecretId" yaml:"userSecretId"`
}

