package awsdevopsguru


// Experimental.
type TfNotificationChannel_FiltersProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devopsguru_notification_channel#message_types TfNotificationChannel#message_types}.
	// Experimental.
	MessageTypes *[]*string `field:"optional" json:"messageTypes" yaml:"messageTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/devopsguru_notification_channel#severities TfNotificationChannel#severities}.
	// Experimental.
	Severities *[]*string `field:"optional" json:"severities" yaml:"severities"`
}

