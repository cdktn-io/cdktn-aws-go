package awsappsync


// Experimental.
type AwsAppsyncDatasource_AuthorizationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#authorization_type AwsAppsyncDatasource#authorization_type}.
	// Experimental.
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// aws_iam_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#aws_iam_config AwsAppsyncDatasource#aws_iam_config}
	// Experimental.
	AwsIamConfig *AwsAppsyncDatasource_AwsIamConfigProperty `field:"optional" json:"awsIamConfig" yaml:"awsIamConfig"`
}

