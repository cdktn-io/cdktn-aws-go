package awscloudwatchinternetmonitor


// Experimental.
type TfMonitor_InternetMeasurementsLogDeliveryProperty struct {
	// s3_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/internetmonitor_monitor#s3_config TfMonitor#s3_config}
	// Experimental.
	S3Config *TfMonitor_S3ConfigProperty `field:"optional" json:"s3Config" yaml:"s3Config"`
}

