package awsappsync


// Experimental.
type TfDatasource_ElasticsearchConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#endpoint TfDatasource#endpoint}.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#region TfDatasource#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

