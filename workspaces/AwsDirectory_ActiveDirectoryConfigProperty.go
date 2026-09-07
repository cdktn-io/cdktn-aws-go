package workspaces


// Experimental.
type AwsDirectory_ActiveDirectoryConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#domain_name AwsDirectory#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/workspaces_directory#service_account_secret_arn AwsDirectory#service_account_secret_arn}.
	// Experimental.
	ServiceAccountSecretArn *string `field:"required" json:"serviceAccountSecretArn" yaml:"serviceAccountSecretArn"`
}

