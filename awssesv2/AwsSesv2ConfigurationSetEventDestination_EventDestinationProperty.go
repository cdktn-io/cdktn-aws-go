package awssesv2


// Experimental.
type AwsSesv2ConfigurationSetEventDestination_EventDestinationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#matching_event_types AwsSesv2ConfigurationSetEventDestination#matching_event_types}.
	// Experimental.
	MatchingEventTypes *[]*string `field:"required" json:"matchingEventTypes" yaml:"matchingEventTypes"`
	// cloud_watch_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#cloud_watch_destination AwsSesv2ConfigurationSetEventDestination#cloud_watch_destination}
	// Experimental.
	CloudWatchDestination *AwsSesv2ConfigurationSetEventDestination_CloudWatchDestinationProperty `field:"optional" json:"cloudWatchDestination" yaml:"cloudWatchDestination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#enabled AwsSesv2ConfigurationSetEventDestination#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// event_bridge_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#event_bridge_destination AwsSesv2ConfigurationSetEventDestination#event_bridge_destination}
	// Experimental.
	EventBridgeDestination *AwsSesv2ConfigurationSetEventDestination_EventBridgeDestinationProperty `field:"optional" json:"eventBridgeDestination" yaml:"eventBridgeDestination"`
	// kinesis_firehose_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#kinesis_firehose_destination AwsSesv2ConfigurationSetEventDestination#kinesis_firehose_destination}
	// Experimental.
	KinesisFirehoseDestination *AwsSesv2ConfigurationSetEventDestination_KinesisFirehoseDestinationProperty `field:"optional" json:"kinesisFirehoseDestination" yaml:"kinesisFirehoseDestination"`
	// pinpoint_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#pinpoint_destination AwsSesv2ConfigurationSetEventDestination#pinpoint_destination}
	// Experimental.
	PinpointDestination *AwsSesv2ConfigurationSetEventDestination_PinpointDestinationProperty `field:"optional" json:"pinpointDestination" yaml:"pinpointDestination"`
	// sns_destination block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_configuration_set_event_destination#sns_destination AwsSesv2ConfigurationSetEventDestination#sns_destination}
	// Experimental.
	SnsDestination *AwsSesv2ConfigurationSetEventDestination_SnsDestinationProperty `field:"optional" json:"snsDestination" yaml:"snsDestination"`
}

