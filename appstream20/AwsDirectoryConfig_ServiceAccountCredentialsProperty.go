package appstream20


// Experimental.
type AwsDirectoryConfig_ServiceAccountCredentialsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_directory_config#account_name AwsDirectoryConfig#account_name}.
	// Experimental.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_directory_config#account_password AwsDirectoryConfig#account_password}.
	// Experimental.
	AccountPassword *string `field:"required" json:"accountPassword" yaml:"accountPassword"`
}

