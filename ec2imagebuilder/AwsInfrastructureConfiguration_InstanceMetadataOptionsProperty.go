package ec2imagebuilder


// Experimental.
type AwsInfrastructureConfiguration_InstanceMetadataOptionsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_infrastructure_configuration#http_put_response_hop_limit AwsInfrastructureConfiguration#http_put_response_hop_limit}.
	// Experimental.
	HttpPutResponseHopLimit *float64 `field:"optional" json:"httpPutResponseHopLimit" yaml:"httpPutResponseHopLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/imagebuilder_infrastructure_configuration#http_tokens AwsInfrastructureConfiguration#http_tokens}.
	// Experimental.
	HttpTokens *string `field:"optional" json:"httpTokens" yaml:"httpTokens"`
}

