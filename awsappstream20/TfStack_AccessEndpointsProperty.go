package awsappstream20


// Experimental.
type TfStack_AccessEndpointsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#endpoint_type TfStack#endpoint_type}.
	// Experimental.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#vpce_id TfStack#vpce_id}.
	// Experimental.
	VpceId *string `field:"optional" json:"vpceId" yaml:"vpceId"`
}

