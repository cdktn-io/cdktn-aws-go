package awssagemakerai


// Experimental.
type TfEndpointConfiguration_CaptureContentTypeHeaderProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#csv_content_types TfEndpointConfiguration#csv_content_types}.
	// Experimental.
	CsvContentTypes *[]*string `field:"optional" json:"csvContentTypes" yaml:"csvContentTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sagemaker_endpoint_configuration#json_content_types TfEndpointConfiguration#json_content_types}.
	// Experimental.
	JsonContentTypes *[]*string `field:"optional" json:"jsonContentTypes" yaml:"jsonContentTypes"`
}

