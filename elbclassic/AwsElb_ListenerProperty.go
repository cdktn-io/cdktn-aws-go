package elbclassic


// Experimental.
type AwsElb_ListenerProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#instance_port AwsElb#instance_port}.
	// Experimental.
	InstancePort *float64 `field:"required" json:"instancePort" yaml:"instancePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#instance_protocol AwsElb#instance_protocol}.
	// Experimental.
	InstanceProtocol *string `field:"required" json:"instanceProtocol" yaml:"instanceProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#lb_port AwsElb#lb_port}.
	// Experimental.
	LbPort *float64 `field:"required" json:"lbPort" yaml:"lbPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#lb_protocol AwsElb#lb_protocol}.
	// Experimental.
	LbProtocol *string `field:"required" json:"lbProtocol" yaml:"lbProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/elb#ssl_certificate_id AwsElb#ssl_certificate_id}.
	// Experimental.
	SslCertificateId *string `field:"optional" json:"sslCertificateId" yaml:"sslCertificateId"`
}

