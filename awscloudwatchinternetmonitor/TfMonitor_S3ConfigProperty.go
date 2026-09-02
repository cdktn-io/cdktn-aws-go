package awscloudwatchinternetmonitor


// Experimental.
type TfMonitor_S3ConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/internetmonitor_monitor#bucket_name TfMonitor#bucket_name}.
	// Experimental.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/internetmonitor_monitor#bucket_prefix TfMonitor#bucket_prefix}.
	// Experimental.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/internetmonitor_monitor#log_delivery_status TfMonitor#log_delivery_status}.
	// Experimental.
	LogDeliveryStatus *string `field:"optional" json:"logDeliveryStatus" yaml:"logDeliveryStatus"`
}

