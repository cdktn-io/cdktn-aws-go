package cloudwatchinternetmonitor


// Experimental.
type AwsMonitor_InternetMeasurementsLogDeliveryProperty struct {
	// s3_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/internetmonitor_monitor#s3_config AwsMonitor#s3_config}
	// Experimental.
	S3Config *AwsMonitor_S3ConfigProperty `field:"optional" json:"s3Config" yaml:"s3Config"`
}

