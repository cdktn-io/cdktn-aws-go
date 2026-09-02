package awseventbridgepipes


// Experimental.
type TfPipe_NetworkConfigurationProperty struct {
	// aws_vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#aws_vpc_configuration TfPipe#aws_vpc_configuration}
	// Experimental.
	AwsVpcConfiguration *TfPipe_AwsVpcConfigurationProperty `field:"optional" json:"awsVpcConfiguration" yaml:"awsVpcConfiguration"`
}

