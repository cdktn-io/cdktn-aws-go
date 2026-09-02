package awssesmailmanager


// Experimental.
type TfIngressPoint_PrivateNetworkConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/mailmanager_ingress_point#vpc_endpoint_id TfIngressPoint#vpc_endpoint_id}.
	// Experimental.
	VpcEndpointId *string `field:"required" json:"vpcEndpointId" yaml:"vpcEndpointId"`
}

