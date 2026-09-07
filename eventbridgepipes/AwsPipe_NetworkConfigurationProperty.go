package eventbridgepipes


// Experimental.
type AwsPipe_NetworkConfigurationProperty struct {
	// aws_vpc_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pipes_pipe#aws_vpc_configuration AwsPipe#aws_vpc_configuration}
	// Experimental.
	AwsVpcConfiguration *AwsPipe_AwsVpcConfigurationProperty `field:"optional" json:"awsVpcConfiguration" yaml:"awsVpcConfiguration"`
}

