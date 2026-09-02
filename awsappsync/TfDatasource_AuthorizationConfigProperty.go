package awsappsync


// Experimental.
type TfDatasource_AuthorizationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#authorization_type TfDatasource#authorization_type}.
	// Experimental.
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// aws_iam_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#aws_iam_config TfDatasource#aws_iam_config}
	// Experimental.
	AwsIamConfig *TfDatasource_AwsIamConfigProperty `field:"optional" json:"awsIamConfig" yaml:"awsIamConfig"`
}

