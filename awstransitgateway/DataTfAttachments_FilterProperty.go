package awstransitgateway


// Experimental.
type DataTfAttachments_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_attachments#name DataTfAttachments#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_attachments#values DataTfAttachments#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

