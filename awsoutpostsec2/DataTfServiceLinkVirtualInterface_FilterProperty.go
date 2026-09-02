package awsoutpostsec2


// Experimental.
type DataTfServiceLinkVirtualInterface_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_service_link_virtual_interface#name DataTfServiceLinkVirtualInterface#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_service_link_virtual_interface#values DataTfServiceLinkVirtualInterface#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

