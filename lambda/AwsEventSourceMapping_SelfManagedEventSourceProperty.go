package lambda


// Experimental.
type AwsEventSourceMapping_SelfManagedEventSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lambda_event_source_mapping#endpoints AwsEventSourceMapping#endpoints}.
	// Experimental.
	Endpoints *map[string]*string `field:"required" json:"endpoints" yaml:"endpoints"`
}

