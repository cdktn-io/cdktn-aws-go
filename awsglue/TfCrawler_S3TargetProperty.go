package awsglue


// Experimental.
type TfCrawler_S3TargetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#path TfCrawler#path}.
	// Experimental.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#connection_name TfCrawler#connection_name}.
	// Experimental.
	ConnectionName *string `field:"optional" json:"connectionName" yaml:"connectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#dlq_event_queue_arn TfCrawler#dlq_event_queue_arn}.
	// Experimental.
	DlqEventQueueArn *string `field:"optional" json:"dlqEventQueueArn" yaml:"dlqEventQueueArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#event_queue_arn TfCrawler#event_queue_arn}.
	// Experimental.
	EventQueueArn *string `field:"optional" json:"eventQueueArn" yaml:"eventQueueArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#exclusions TfCrawler#exclusions}.
	// Experimental.
	Exclusions *[]*string `field:"optional" json:"exclusions" yaml:"exclusions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/glue_crawler#sample_size TfCrawler#sample_size}.
	// Experimental.
	SampleSize *float64 `field:"optional" json:"sampleSize" yaml:"sampleSize"`
}

