package awsappsync


// Experimental.
type TfDatasource_AwsIamConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#signing_region TfDatasource#signing_region}.
	// Experimental.
	SigningRegion *string `field:"optional" json:"signingRegion" yaml:"signingRegion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#signing_service_name TfDatasource#signing_service_name}.
	// Experimental.
	SigningServiceName *string `field:"optional" json:"signingServiceName" yaml:"signingServiceName"`
}

