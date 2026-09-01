package awsvpc


// Experimental.
type AwsEc2TrafficMirrorFilterRule_SourcePortRangeProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#from_port AwsEc2TrafficMirrorFilterRule#from_port}.
	// Experimental.
	FromPort *float64 `field:"optional" json:"fromPort" yaml:"fromPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_traffic_mirror_filter_rule#to_port AwsEc2TrafficMirrorFilterRule#to_port}.
	// Experimental.
	ToPort *float64 `field:"optional" json:"toPort" yaml:"toPort"`
}

