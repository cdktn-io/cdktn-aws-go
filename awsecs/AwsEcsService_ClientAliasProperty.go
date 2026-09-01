package awsecs


// Experimental.
type AwsEcsService_ClientAliasProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#port AwsEcsService#port}.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#dns_name AwsEcsService#dns_name}.
	// Experimental.
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// test_traffic_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#test_traffic_rules AwsEcsService#test_traffic_rules}
	// Experimental.
	TestTrafficRules interface{} `field:"optional" json:"testTrafficRules" yaml:"testTrafficRules"`
}

