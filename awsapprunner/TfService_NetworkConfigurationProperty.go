package awsapprunner


// Experimental.
type TfService_NetworkConfigurationProperty struct {
	// egress_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#egress_configuration TfService#egress_configuration}
	// Experimental.
	EgressConfiguration *TfService_EgressConfigurationProperty `field:"optional" json:"egressConfiguration" yaml:"egressConfiguration"`
	// ingress_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#ingress_configuration TfService#ingress_configuration}
	// Experimental.
	IngressConfiguration *TfService_IngressConfigurationProperty `field:"optional" json:"ingressConfiguration" yaml:"ingressConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#ip_address_type TfService#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
}

