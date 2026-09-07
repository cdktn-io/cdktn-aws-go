package lambda


// Experimental.
type AwsEventSourceMapping_SourceAccessConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#type AwsEventSourceMapping#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#uri AwsEventSourceMapping#uri}.
	// Experimental.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
}

