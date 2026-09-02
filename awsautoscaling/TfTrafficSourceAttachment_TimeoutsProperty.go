package awsautoscaling


// Experimental.
type TfTrafficSourceAttachment_TimeoutsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_traffic_source_attachment#create TfTrafficSourceAttachment#create}.
	// Experimental.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_traffic_source_attachment#delete TfTrafficSourceAttachment#delete}.
	// Experimental.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
}

