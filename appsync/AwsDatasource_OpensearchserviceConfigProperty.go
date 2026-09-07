package appsync


// Experimental.
type AwsDatasource_OpensearchserviceConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#endpoint AwsDatasource#endpoint}.
	// Experimental.
	Endpoint *string `field:"required" json:"endpoint" yaml:"endpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appsync_datasource#region AwsDatasource#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

