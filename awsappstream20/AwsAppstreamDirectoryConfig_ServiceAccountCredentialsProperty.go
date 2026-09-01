package awsappstream20


// Experimental.
type AwsAppstreamDirectoryConfig_ServiceAccountCredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_directory_config#account_name AwsAppstreamDirectoryConfig#account_name}.
	// Experimental.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_directory_config#account_password AwsAppstreamDirectoryConfig#account_password}.
	// Experimental.
	AccountPassword *string `field:"required" json:"accountPassword" yaml:"accountPassword"`
}

