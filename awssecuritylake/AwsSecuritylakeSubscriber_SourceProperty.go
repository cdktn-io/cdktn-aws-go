package awssecuritylake


// Experimental.
type AwsSecuritylakeSubscriber_SourceProperty struct {
	// aws_log_source_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#aws_log_source_resource AwsSecuritylakeSubscriber#aws_log_source_resource}
	// Experimental.
	AwsLogSourceResource interface{} `field:"optional" json:"awsLogSourceResource" yaml:"awsLogSourceResource"`
	// custom_log_source_resource block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#custom_log_source_resource AwsSecuritylakeSubscriber#custom_log_source_resource}
	// Experimental.
	CustomLogSourceResource interface{} `field:"optional" json:"customLogSourceResource" yaml:"customLogSourceResource"`
}

