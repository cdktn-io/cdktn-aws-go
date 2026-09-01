package awsrds


// Experimental.
type AwsDbProxy_AuthProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy#auth_scheme AwsDbProxy#auth_scheme}.
	// Experimental.
	AuthScheme *string `field:"optional" json:"authScheme" yaml:"authScheme"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy#client_password_auth_type AwsDbProxy#client_password_auth_type}.
	// Experimental.
	ClientPasswordAuthType *string `field:"optional" json:"clientPasswordAuthType" yaml:"clientPasswordAuthType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy#description AwsDbProxy#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy#iam_auth AwsDbProxy#iam_auth}.
	// Experimental.
	IamAuth *string `field:"optional" json:"iamAuth" yaml:"iamAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy#secret_arn AwsDbProxy#secret_arn}.
	// Experimental.
	SecretArn *string `field:"optional" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/db_proxy#username AwsDbProxy#username}.
	// Experimental.
	Username *string `field:"optional" json:"username" yaml:"username"`
}

