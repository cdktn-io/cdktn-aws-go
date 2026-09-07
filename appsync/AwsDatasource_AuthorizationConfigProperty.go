package appsync


// Experimental.
type AwsDatasource_AuthorizationConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#authorization_type AwsDatasource#authorization_type}.
	// Experimental.
	AuthorizationType *string `field:"optional" json:"authorizationType" yaml:"authorizationType"`
	// aws_iam_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#aws_iam_config AwsDatasource#aws_iam_config}
	// Experimental.
	AwsIamConfig *AwsDatasource_AwsIamConfigProperty `field:"optional" json:"awsIamConfig" yaml:"awsIamConfig"`
}

