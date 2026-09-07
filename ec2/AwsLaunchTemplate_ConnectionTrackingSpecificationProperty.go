package ec2


// Experimental.
type AwsLaunchTemplate_ConnectionTrackingSpecificationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#tcp_established_timeout AwsLaunchTemplate#tcp_established_timeout}.
	// Experimental.
	TcpEstablishedTimeout *float64 `field:"optional" json:"tcpEstablishedTimeout" yaml:"tcpEstablishedTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#udp_stream_timeout AwsLaunchTemplate#udp_stream_timeout}.
	// Experimental.
	UdpStreamTimeout *float64 `field:"optional" json:"udpStreamTimeout" yaml:"udpStreamTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/launch_template#udp_timeout AwsLaunchTemplate#udp_timeout}.
	// Experimental.
	UdpTimeout *float64 `field:"optional" json:"udpTimeout" yaml:"udpTimeout"`
}

