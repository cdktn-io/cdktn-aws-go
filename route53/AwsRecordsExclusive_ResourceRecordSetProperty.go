package route53


// Experimental.
type AwsRecordsExclusive_ResourceRecordSetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#name AwsRecordsExclusive#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// alias_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#alias_target AwsRecordsExclusive#alias_target}
	// Experimental.
	AliasTarget interface{} `field:"optional" json:"aliasTarget" yaml:"aliasTarget"`
	// cidr_routing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#cidr_routing_config AwsRecordsExclusive#cidr_routing_config}
	// Experimental.
	CidrRoutingConfig interface{} `field:"optional" json:"cidrRoutingConfig" yaml:"cidrRoutingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#failover AwsRecordsExclusive#failover}.
	// Experimental.
	Failover *string `field:"optional" json:"failover" yaml:"failover"`
	// geolocation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#geolocation AwsRecordsExclusive#geolocation}
	// Experimental.
	Geolocation interface{} `field:"optional" json:"geolocation" yaml:"geolocation"`
	// geoproximity_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#geoproximity_location AwsRecordsExclusive#geoproximity_location}
	// Experimental.
	GeoproximityLocation interface{} `field:"optional" json:"geoproximityLocation" yaml:"geoproximityLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#health_check_id AwsRecordsExclusive#health_check_id}.
	// Experimental.
	HealthCheckId *string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#multi_value_answer AwsRecordsExclusive#multi_value_answer}.
	// Experimental.
	MultiValueAnswer interface{} `field:"optional" json:"multiValueAnswer" yaml:"multiValueAnswer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#region AwsRecordsExclusive#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// resource_records block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#resource_records AwsRecordsExclusive#resource_records}
	// Experimental.
	ResourceRecords interface{} `field:"optional" json:"resourceRecords" yaml:"resourceRecords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#set_identifier AwsRecordsExclusive#set_identifier}.
	// Experimental.
	SetIdentifier *string `field:"optional" json:"setIdentifier" yaml:"setIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#traffic_policy_instance_id AwsRecordsExclusive#traffic_policy_instance_id}.
	// Experimental.
	TrafficPolicyInstanceId *string `field:"optional" json:"trafficPolicyInstanceId" yaml:"trafficPolicyInstanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#ttl AwsRecordsExclusive#ttl}.
	// Experimental.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#type AwsRecordsExclusive#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#weight AwsRecordsExclusive#weight}.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

