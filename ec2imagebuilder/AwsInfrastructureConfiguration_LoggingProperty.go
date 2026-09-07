package ec2imagebuilder


// Experimental.
type AwsInfrastructureConfiguration_LoggingProperty struct {
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_infrastructure_configuration#s3_logs AwsInfrastructureConfiguration#s3_logs}
	// Experimental.
	S3Logs *AwsInfrastructureConfiguration_S3LogsProperty `field:"required" json:"s3Logs" yaml:"s3Logs"`
}

