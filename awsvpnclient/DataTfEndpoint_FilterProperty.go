package awsvpnclient


// Experimental.
type DataTfEndpoint_FilterProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_client_vpn_endpoint#name DataTfEndpoint#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/ec2_client_vpn_endpoint#values DataTfEndpoint#values}.
	// Experimental.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}

