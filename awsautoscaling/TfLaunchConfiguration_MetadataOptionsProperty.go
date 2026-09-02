package awsautoscaling


// Experimental.
type TfLaunchConfiguration_MetadataOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_configuration#http_endpoint TfLaunchConfiguration#http_endpoint}.
	// Experimental.
	HttpEndpoint *string `field:"optional" json:"httpEndpoint" yaml:"httpEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_configuration#http_put_response_hop_limit TfLaunchConfiguration#http_put_response_hop_limit}.
	// Experimental.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_configuration#http_tokens TfLaunchConfiguration#http_tokens}.
	// Experimental.
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
}

