package awsglobalaccelerator


// Experimental.
type TfEndpointGroup_EndpointConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_endpoint_group#attachment_arn TfEndpointGroup#attachment_arn}.
	// Experimental.
	AttachmentArn *string `field:"optional" json:"attachmentArn" yaml:"attachmentArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_endpoint_group#client_ip_preservation_enabled TfEndpointGroup#client_ip_preservation_enabled}.
	// Experimental.
	ClientIpPreservationEnabled interface{} `field:"optional" json:"clientIpPreservationEnabled" yaml:"clientIpPreservationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_endpoint_group#endpoint_id TfEndpointGroup#endpoint_id}.
	// Experimental.
	EndpointId *string `field:"optional" json:"endpointId" yaml:"endpointId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/globalaccelerator_endpoint_group#weight TfEndpointGroup#weight}.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

