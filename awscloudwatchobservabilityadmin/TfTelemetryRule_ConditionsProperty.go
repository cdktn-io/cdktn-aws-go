package awscloudwatchobservabilityadmin


// Experimental.
type TfTelemetryRule_ConditionsProperty struct {
	// action_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#action_condition TfTelemetryRule#action_condition}
	// Experimental.
	ActionCondition interface{} `field:"optional" json:"actionCondition" yaml:"actionCondition"`
	// label_name_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/observabilityadmin_telemetry_rule#label_name_condition TfTelemetryRule#label_name_condition}
	// Experimental.
	LabelNameCondition interface{} `field:"optional" json:"labelNameCondition" yaml:"labelNameCondition"`
}

