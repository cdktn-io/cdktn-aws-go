package awsappfabric


// Experimental.
type AwsAppfabricIngestionDestination_ProcessingConfigurationAuditLogProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_ingestion_destination#format AwsAppfabricIngestionDestination#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appfabric_ingestion_destination#schema AwsAppfabricIngestionDestination#schema}.
	// Experimental.
	Schema *string `field:"required" json:"schema" yaml:"schema"`
}

