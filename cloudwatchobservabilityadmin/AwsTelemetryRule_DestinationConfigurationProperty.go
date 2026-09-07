package cloudwatchobservabilityadmin


// Experimental.
type AwsTelemetryRule_DestinationConfigurationProperty struct {
	// cloudtrail_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#cloudtrail_parameters AwsTelemetryRule#cloudtrail_parameters}
	// Experimental.
	CloudtrailParameters interface{} `field:"optional" json:"cloudtrailParameters" yaml:"cloudtrailParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#destination_pattern AwsTelemetryRule#destination_pattern}.
	// Experimental.
	DestinationPattern *string `field:"optional" json:"destinationPattern" yaml:"destinationPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#destination_type AwsTelemetryRule#destination_type}.
	// Experimental.
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// elb_load_balancer_logging_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#elb_load_balancer_logging_parameters AwsTelemetryRule#elb_load_balancer_logging_parameters}
	// Experimental.
	ElbLoadBalancerLoggingParameters interface{} `field:"optional" json:"elbLoadBalancerLoggingParameters" yaml:"elbLoadBalancerLoggingParameters"`
	// log_delivery_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#log_delivery_parameters AwsTelemetryRule#log_delivery_parameters}
	// Experimental.
	LogDeliveryParameters interface{} `field:"optional" json:"logDeliveryParameters" yaml:"logDeliveryParameters"`
	// msk_monitoring_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#msk_monitoring_parameters AwsTelemetryRule#msk_monitoring_parameters}
	// Experimental.
	MskMonitoringParameters interface{} `field:"optional" json:"mskMonitoringParameters" yaml:"mskMonitoringParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#retention_in_days AwsTelemetryRule#retention_in_days}.
	// Experimental.
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// vpc_flow_log_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#vpc_flow_log_parameters AwsTelemetryRule#vpc_flow_log_parameters}
	// Experimental.
	VpcFlowLogParameters interface{} `field:"optional" json:"vpcFlowLogParameters" yaml:"vpcFlowLogParameters"`
	// waf_logging_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#waf_logging_parameters AwsTelemetryRule#waf_logging_parameters}
	// Experimental.
	WafLoggingParameters interface{} `field:"optional" json:"wafLoggingParameters" yaml:"wafLoggingParameters"`
}

