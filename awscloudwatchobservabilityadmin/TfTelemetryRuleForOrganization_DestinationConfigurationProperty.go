package awscloudwatchobservabilityadmin


// Experimental.
type TfTelemetryRuleForOrganization_DestinationConfigurationProperty struct {
	// cloudtrail_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#cloudtrail_parameters TfTelemetryRuleForOrganization#cloudtrail_parameters}
	// Experimental.
	CloudtrailParameters interface{} `field:"optional" json:"cloudtrailParameters" yaml:"cloudtrailParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#destination_pattern TfTelemetryRuleForOrganization#destination_pattern}.
	// Experimental.
	DestinationPattern *string `field:"optional" json:"destinationPattern" yaml:"destinationPattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#destination_type TfTelemetryRuleForOrganization#destination_type}.
	// Experimental.
	DestinationType *string `field:"optional" json:"destinationType" yaml:"destinationType"`
	// elb_load_balancer_logging_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#elb_load_balancer_logging_parameters TfTelemetryRuleForOrganization#elb_load_balancer_logging_parameters}
	// Experimental.
	ElbLoadBalancerLoggingParameters interface{} `field:"optional" json:"elbLoadBalancerLoggingParameters" yaml:"elbLoadBalancerLoggingParameters"`
	// log_delivery_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#log_delivery_parameters TfTelemetryRuleForOrganization#log_delivery_parameters}
	// Experimental.
	LogDeliveryParameters interface{} `field:"optional" json:"logDeliveryParameters" yaml:"logDeliveryParameters"`
	// msk_monitoring_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#msk_monitoring_parameters TfTelemetryRuleForOrganization#msk_monitoring_parameters}
	// Experimental.
	MskMonitoringParameters interface{} `field:"optional" json:"mskMonitoringParameters" yaml:"mskMonitoringParameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#retention_in_days TfTelemetryRuleForOrganization#retention_in_days}.
	// Experimental.
	RetentionInDays *float64 `field:"optional" json:"retentionInDays" yaml:"retentionInDays"`
	// vpc_flow_log_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#vpc_flow_log_parameters TfTelemetryRuleForOrganization#vpc_flow_log_parameters}
	// Experimental.
	VpcFlowLogParameters interface{} `field:"optional" json:"vpcFlowLogParameters" yaml:"vpcFlowLogParameters"`
	// waf_logging_parameters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule_for_organization#waf_logging_parameters TfTelemetryRuleForOrganization#waf_logging_parameters}
	// Experimental.
	WafLoggingParameters interface{} `field:"optional" json:"wafLoggingParameters" yaml:"wafLoggingParameters"`
}

