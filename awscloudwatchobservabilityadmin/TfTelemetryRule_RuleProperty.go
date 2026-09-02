package awscloudwatchobservabilityadmin


// Experimental.
type TfTelemetryRule_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#telemetry_type TfTelemetryRule#telemetry_type}.
	// Experimental.
	TelemetryType *string `field:"required" json:"telemetryType" yaml:"telemetryType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#allow_field_updates TfTelemetryRule#allow_field_updates}.
	// Experimental.
	AllowFieldUpdates interface{} `field:"optional" json:"allowFieldUpdates" yaml:"allowFieldUpdates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#all_regions TfTelemetryRule#all_regions}.
	// Experimental.
	AllRegions interface{} `field:"optional" json:"allRegions" yaml:"allRegions"`
	// destination_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#destination_configuration TfTelemetryRule#destination_configuration}
	// Experimental.
	DestinationConfiguration interface{} `field:"optional" json:"destinationConfiguration" yaml:"destinationConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#regions TfTelemetryRule#regions}.
	// Experimental.
	Regions *[]*string `field:"optional" json:"regions" yaml:"regions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#resource_type TfTelemetryRule#resource_type}.
	// Experimental.
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#scope TfTelemetryRule#scope}.
	// Experimental.
	Scope *string `field:"optional" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#selection_criteria TfTelemetryRule#selection_criteria}.
	// Experimental.
	SelectionCriteria *string `field:"optional" json:"selectionCriteria" yaml:"selectionCriteria"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#telemetry_source_types TfTelemetryRule#telemetry_source_types}.
	// Experimental.
	TelemetrySourceTypes *[]*string `field:"optional" json:"telemetrySourceTypes" yaml:"telemetrySourceTypes"`
}

