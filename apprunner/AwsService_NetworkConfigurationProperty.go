package apprunner


// Experimental.
type AwsService_NetworkConfigurationProperty struct {
	// egress_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#egress_configuration AwsService#egress_configuration}
	// Experimental.
	EgressConfiguration *AwsService_EgressConfigurationProperty `field:"optional" json:"egressConfiguration" yaml:"egressConfiguration"`
	// ingress_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#ingress_configuration AwsService#ingress_configuration}
	// Experimental.
	IngressConfiguration *AwsService_IngressConfigurationProperty `field:"optional" json:"ingressConfiguration" yaml:"ingressConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#ip_address_type AwsService#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
}

