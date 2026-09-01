package awscloudwatchobservabilityadmin


// Experimental.
type AwsObservabilityadminTelemetryRule_RedactedFieldsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#method AwsObservabilityadminTelemetryRule#method}.
	// Experimental.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#query_string AwsObservabilityadminTelemetryRule#query_string}.
	// Experimental.
	QueryString *string `field:"optional" json:"queryString" yaml:"queryString"`
	// single_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#single_header AwsObservabilityadminTelemetryRule#single_header}
	// Experimental.
	SingleHeader interface{} `field:"optional" json:"singleHeader" yaml:"singleHeader"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#uri_path AwsObservabilityadminTelemetryRule#uri_path}.
	// Experimental.
	UriPath *string `field:"optional" json:"uriPath" yaml:"uriPath"`
}

