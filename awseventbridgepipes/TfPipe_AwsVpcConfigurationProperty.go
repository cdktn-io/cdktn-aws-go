package awseventbridgepipes


// Experimental.
type TfPipe_AwsVpcConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#assign_public_ip TfPipe#assign_public_ip}.
	// Experimental.
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#security_groups TfPipe#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#subnets TfPipe#subnets}.
	// Experimental.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

