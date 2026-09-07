package lightsail


// Experimental.
type AwsInstancePublicPorts_PortInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance_public_ports#from_port AwsInstancePublicPorts#from_port}.
	// Experimental.
	FromPort *float64 `field:"required" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance_public_ports#protocol AwsInstancePublicPorts#protocol}.
	// Experimental.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance_public_ports#to_port AwsInstancePublicPorts#to_port}.
	// Experimental.
	ToPort *float64 `field:"required" json:"toPort" yaml:"toPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance_public_ports#cidr_list_aliases AwsInstancePublicPorts#cidr_list_aliases}.
	// Experimental.
	CidrListAliases *[]*string `field:"optional" json:"cidrListAliases" yaml:"cidrListAliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance_public_ports#cidrs AwsInstancePublicPorts#cidrs}.
	// Experimental.
	Cidrs *[]*string `field:"optional" json:"cidrs" yaml:"cidrs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_instance_public_ports#ipv6_cidrs AwsInstancePublicPorts#ipv6_cidrs}.
	// Experimental.
	Ipv6Cidrs *[]*string `field:"optional" json:"ipv6Cidrs" yaml:"ipv6Cidrs"`
}

