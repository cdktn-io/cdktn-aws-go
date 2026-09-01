package awscloudwatchobservabilityadmin


// Experimental.
type AwsObservabilityadminTelemetryRule_VpcFlowLogParametersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#log_format AwsObservabilityadminTelemetryRule#log_format}.
	// Experimental.
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#max_aggregation_interval AwsObservabilityadminTelemetryRule#max_aggregation_interval}.
	// Experimental.
	MaxAggregationInterval *float64 `field:"optional" json:"maxAggregationInterval" yaml:"maxAggregationInterval"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#traffic_type AwsObservabilityadminTelemetryRule#traffic_type}.
	// Experimental.
	TrafficType *string `field:"optional" json:"trafficType" yaml:"trafficType"`
}

