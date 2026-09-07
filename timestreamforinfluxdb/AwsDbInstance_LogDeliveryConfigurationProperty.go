package timestreamforinfluxdb


// Experimental.
type AwsDbInstance_LogDeliveryConfigurationProperty struct {
	// s3_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreaminfluxdb_db_instance#s3_configuration AwsDbInstance#s3_configuration}
	// Experimental.
	S3Configuration interface{} `field:"optional" json:"s3Configuration" yaml:"s3Configuration"`
}

