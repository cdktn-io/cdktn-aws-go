package awsecs


// Experimental.
type TfService_ClientAliasProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#port TfService#port}.
	// Experimental.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#dns_name TfService#dns_name}.
	// Experimental.
	DnsName *string `field:"optional" json:"dnsName" yaml:"dnsName"`
	// test_traffic_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ecs_service#test_traffic_rules TfService#test_traffic_rules}
	// Experimental.
	TestTrafficRules interface{} `field:"optional" json:"testTrafficRules" yaml:"testTrafficRules"`
}

