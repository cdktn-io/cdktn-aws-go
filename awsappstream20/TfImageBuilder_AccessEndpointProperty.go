package awsappstream20


// Experimental.
type TfImageBuilder_AccessEndpointProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#endpoint_type TfImageBuilder#endpoint_type}.
	// Experimental.
	EndpointType *string `field:"required" json:"endpointType" yaml:"endpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/appstream_image_builder#vpce_id TfImageBuilder#vpce_id}.
	// Experimental.
	VpceId *string `field:"optional" json:"vpceId" yaml:"vpceId"`
}

