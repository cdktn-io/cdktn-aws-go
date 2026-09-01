package awsapprunner


// Experimental.
type AwsApprunnerService_NetworkConfigurationProperty struct {
	// egress_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#egress_configuration AwsApprunnerService#egress_configuration}
	// Experimental.
	EgressConfiguration *AwsApprunnerService_EgressConfigurationProperty `field:"optional" json:"egressConfiguration" yaml:"egressConfiguration"`
	// ingress_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#ingress_configuration AwsApprunnerService#ingress_configuration}
	// Experimental.
	IngressConfiguration *AwsApprunnerService_IngressConfigurationProperty `field:"optional" json:"ingressConfiguration" yaml:"ingressConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/apprunner_service#ip_address_type AwsApprunnerService#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
}

