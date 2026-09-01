package awscloudwatchlogs


// Experimental.
type AwsCloudwatchLogS3TableIntegrationSource_TimeoutsProperty struct {
	// A string that can be [parsed as a duration](https://pkg.go.dev/time#ParseDuration) consisting of numbers and unit suffixes, such as "30s" or "2h45m". Valid time units are "s" (seconds), "m" (minutes), "h" (hours). Setting a timeout for a Delete operation is only applicable if changes are saved into state before the destroy operation occurs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudwatch_log_s3_table_integration_source#delete AwsCloudwatchLogS3TableIntegrationSource#delete}
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

