package awsvpc


// Experimental.
type TfEndpoint_SubnetConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint#ipv4 TfEndpoint#ipv4}.
	// Experimental.
	Ipv4 *string `field:"optional" json:"ipv4" yaml:"ipv4"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint#ipv6 TfEndpoint#ipv6}.
	// Experimental.
	Ipv6 *string `field:"optional" json:"ipv6" yaml:"ipv6"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint#subnet_id TfEndpoint#subnet_id}.
	// Experimental.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}

