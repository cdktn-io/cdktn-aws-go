package sesv2


// Experimental.
type AwsConfigurationSetEventDestination_EventDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#matching_event_types AwsConfigurationSetEventDestination#matching_event_types}.
	// Experimental.
	MatchingEventTypes *[]*string `field:"required" json:"matchingEventTypes" yaml:"matchingEventTypes"`
	// cloud_watch_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#cloud_watch_destination AwsConfigurationSetEventDestination#cloud_watch_destination}
	// Experimental.
	CloudWatchDestination *AwsConfigurationSetEventDestination_CloudWatchDestinationProperty `field:"optional" json:"cloudWatchDestination" yaml:"cloudWatchDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#enabled AwsConfigurationSetEventDestination#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// event_bridge_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#event_bridge_destination AwsConfigurationSetEventDestination#event_bridge_destination}
	// Experimental.
	EventBridgeDestination *AwsConfigurationSetEventDestination_EventBridgeDestinationProperty `field:"optional" json:"eventBridgeDestination" yaml:"eventBridgeDestination"`
	// kinesis_firehose_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#kinesis_firehose_destination AwsConfigurationSetEventDestination#kinesis_firehose_destination}
	// Experimental.
	KinesisFirehoseDestination *AwsConfigurationSetEventDestination_KinesisFirehoseDestinationProperty `field:"optional" json:"kinesisFirehoseDestination" yaml:"kinesisFirehoseDestination"`
	// pinpoint_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#pinpoint_destination AwsConfigurationSetEventDestination#pinpoint_destination}
	// Experimental.
	PinpointDestination *AwsConfigurationSetEventDestination_PinpointDestinationProperty `field:"optional" json:"pinpointDestination" yaml:"pinpointDestination"`
	// sns_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#sns_destination AwsConfigurationSetEventDestination#sns_destination}
	// Experimental.
	SnsDestination *AwsConfigurationSetEventDestination_SnsDestinationProperty `field:"optional" json:"snsDestination" yaml:"snsDestination"`
}

