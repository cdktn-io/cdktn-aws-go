package awsglue


// Experimental.
type TfCrawler_CatalogTargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#database_name TfCrawler#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#tables TfCrawler#tables}.
	// Experimental.
	Tables *[]*string `field:"required" json:"tables" yaml:"tables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#connection_name TfCrawler#connection_name}.
	// Experimental.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#dlq_event_queue_arn TfCrawler#dlq_event_queue_arn}.
	// Experimental.
	DlqEventQueueArn *string `field:"optional" json:"dlqEventQueueArn" yaml:"dlqEventQueueArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#event_queue_arn TfCrawler#event_queue_arn}.
	// Experimental.
	EventQueueArn *string `field:"optional" json:"eventQueueArn" yaml:"eventQueueArn"`
}

