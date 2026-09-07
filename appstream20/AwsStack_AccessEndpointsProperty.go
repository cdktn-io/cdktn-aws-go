package appstream20


// Experimental.
type AwsStack_AccessEndpointsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#endpoint_type AwsStack#endpoint_type}.
	// Experimental.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_stack#vpce_id AwsStack#vpce_id}.
	// Experimental.
	VpceId *string `field:"optional" json:"vpceId" yaml:"vpceId"`
}

