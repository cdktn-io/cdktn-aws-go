package awsroute53


// Experimental.
type AwsRoute53RecordsExclusive_ResourceRecordSetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#name AwsRoute53RecordsExclusive#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// alias_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#alias_target AwsRoute53RecordsExclusive#alias_target}
	// Experimental.
	AliasTarget interface{} `field:"optional" json:"aliasTarget" yaml:"aliasTarget"`
	// cidr_routing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#cidr_routing_config AwsRoute53RecordsExclusive#cidr_routing_config}
	// Experimental.
	CidrRoutingConfig interface{} `field:"optional" json:"cidrRoutingConfig" yaml:"cidrRoutingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#failover AwsRoute53RecordsExclusive#failover}.
	// Experimental.
	Failover *string `field:"optional" json:"failover" yaml:"failover"`
	// geolocation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#geolocation AwsRoute53RecordsExclusive#geolocation}
	// Experimental.
	Geolocation interface{} `field:"optional" json:"geolocation" yaml:"geolocation"`
	// geoproximity_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#geoproximity_location AwsRoute53RecordsExclusive#geoproximity_location}
	// Experimental.
	GeoproximityLocation interface{} `field:"optional" json:"geoproximityLocation" yaml:"geoproximityLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#health_check_id AwsRoute53RecordsExclusive#health_check_id}.
	// Experimental.
	HealthCheckId *string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#multi_value_answer AwsRoute53RecordsExclusive#multi_value_answer}.
	// Experimental.
	MultiValueAnswer interface{} `field:"optional" json:"multiValueAnswer" yaml:"multiValueAnswer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#region AwsRoute53RecordsExclusive#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// resource_records block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#resource_records AwsRoute53RecordsExclusive#resource_records}
	// Experimental.
	ResourceRecords interface{} `field:"optional" json:"resourceRecords" yaml:"resourceRecords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#set_identifier AwsRoute53RecordsExclusive#set_identifier}.
	// Experimental.
	SetIdentifier *string `field:"optional" json:"setIdentifier" yaml:"setIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#traffic_policy_instance_id AwsRoute53RecordsExclusive#traffic_policy_instance_id}.
	// Experimental.
	TrafficPolicyInstanceId *string `field:"optional" json:"trafficPolicyInstanceId" yaml:"trafficPolicyInstanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#ttl AwsRoute53RecordsExclusive#ttl}.
	// Experimental.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#type AwsRoute53RecordsExclusive#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#weight AwsRoute53RecordsExclusive#weight}.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

