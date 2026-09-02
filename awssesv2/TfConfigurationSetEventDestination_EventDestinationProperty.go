package awssesv2


// Experimental.
type TfConfigurationSetEventDestination_EventDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#matching_event_types TfConfigurationSetEventDestination#matching_event_types}.
	// Experimental.
	MatchingEventTypes *[]*string `field:"required" json:"matchingEventTypes" yaml:"matchingEventTypes"`
	// cloud_watch_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#cloud_watch_destination TfConfigurationSetEventDestination#cloud_watch_destination}
	// Experimental.
	CloudWatchDestination *TfConfigurationSetEventDestination_CloudWatchDestinationProperty `field:"optional" json:"cloudWatchDestination" yaml:"cloudWatchDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#enabled TfConfigurationSetEventDestination#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// event_bridge_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#event_bridge_destination TfConfigurationSetEventDestination#event_bridge_destination}
	// Experimental.
	EventBridgeDestination *TfConfigurationSetEventDestination_EventBridgeDestinationProperty `field:"optional" json:"eventBridgeDestination" yaml:"eventBridgeDestination"`
	// kinesis_firehose_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#kinesis_firehose_destination TfConfigurationSetEventDestination#kinesis_firehose_destination}
	// Experimental.
	KinesisFirehoseDestination *TfConfigurationSetEventDestination_KinesisFirehoseDestinationProperty `field:"optional" json:"kinesisFirehoseDestination" yaml:"kinesisFirehoseDestination"`
	// pinpoint_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#pinpoint_destination TfConfigurationSetEventDestination#pinpoint_destination}
	// Experimental.
	PinpointDestination *TfConfigurationSetEventDestination_PinpointDestinationProperty `field:"optional" json:"pinpointDestination" yaml:"pinpointDestination"`
	// sns_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#sns_destination TfConfigurationSetEventDestination#sns_destination}
	// Experimental.
	SnsDestination *TfConfigurationSetEventDestination_SnsDestinationProperty `field:"optional" json:"snsDestination" yaml:"snsDestination"`
}

