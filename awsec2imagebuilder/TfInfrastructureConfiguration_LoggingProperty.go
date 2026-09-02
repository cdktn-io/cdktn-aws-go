package awsec2imagebuilder


// Experimental.
type TfInfrastructureConfiguration_LoggingProperty struct {
	// s3_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_infrastructure_configuration#s3_logs TfInfrastructureConfiguration#s3_logs}
	// Experimental.
	S3Logs *TfInfrastructureConfiguration_S3LogsProperty `field:"required" json:"s3Logs" yaml:"s3Logs"`
}

