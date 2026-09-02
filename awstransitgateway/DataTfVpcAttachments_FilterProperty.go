package awstransitgateway


// Experimental.
type DataTfVpcAttachments_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_vpc_attachments#name DataTfVpcAttachments#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_transit_gateway_vpc_attachments#values DataTfVpcAttachments#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

