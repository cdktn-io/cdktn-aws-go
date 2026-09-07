package connect


// Experimental.
type AwsRoutingProfile_QueueConfigsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_routing_profile#channel AwsRoutingProfile#channel}.
	// Experimental.
	Channel *string `field:"required" json:"channel" yaml:"channel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_routing_profile#delay AwsRoutingProfile#delay}.
	// Experimental.
	Delay *float64 `field:"required" json:"delay" yaml:"delay"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_routing_profile#priority AwsRoutingProfile#priority}.
	// Experimental.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/connect_routing_profile#queue_id AwsRoutingProfile#queue_id}.
	// Experimental.
	QueueId *string `field:"required" json:"queueId" yaml:"queueId"`
}

