package sesmailmanager


// Experimental.
type AwsIngressPoint_NetworkConfigurationProperty struct {
	// private_network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#private_network_configuration AwsIngressPoint#private_network_configuration}
	// Experimental.
	PrivateNetworkConfiguration interface{} `field:"optional" json:"privateNetworkConfiguration" yaml:"privateNetworkConfiguration"`
	// public_network_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#public_network_configuration AwsIngressPoint#public_network_configuration}
	// Experimental.
	PublicNetworkConfiguration interface{} `field:"optional" json:"publicNetworkConfiguration" yaml:"publicNetworkConfiguration"`
}

