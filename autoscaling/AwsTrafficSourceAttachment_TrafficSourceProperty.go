package autoscaling


// Experimental.
type AwsTrafficSourceAttachment_TrafficSourceProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_traffic_source_attachment#identifier AwsTrafficSourceAttachment#identifier}.
	// Experimental.
	Identifier *string `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/autoscaling_traffic_source_attachment#type AwsTrafficSourceAttachment#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

