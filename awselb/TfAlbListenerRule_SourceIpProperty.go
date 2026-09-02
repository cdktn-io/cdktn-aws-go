package awselb


// Experimental.
type TfAlbListenerRule_SourceIpProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#ip_address_type TfAlbListenerRule#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener_rule#values TfAlbListenerRule#values}.
	// Experimental.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

