package awsroute53


// Experimental.
type TfRecordsExclusive_ResourceRecordSetProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#name TfRecordsExclusive#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// alias_target block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#alias_target TfRecordsExclusive#alias_target}
	// Experimental.
	AliasTarget interface{} `field:"optional" json:"aliasTarget" yaml:"aliasTarget"`
	// cidr_routing_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#cidr_routing_config TfRecordsExclusive#cidr_routing_config}
	// Experimental.
	CidrRoutingConfig interface{} `field:"optional" json:"cidrRoutingConfig" yaml:"cidrRoutingConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#failover TfRecordsExclusive#failover}.
	// Experimental.
	Failover *string `field:"optional" json:"failover" yaml:"failover"`
	// geolocation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#geolocation TfRecordsExclusive#geolocation}
	// Experimental.
	Geolocation interface{} `field:"optional" json:"geolocation" yaml:"geolocation"`
	// geoproximity_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#geoproximity_location TfRecordsExclusive#geoproximity_location}
	// Experimental.
	GeoproximityLocation interface{} `field:"optional" json:"geoproximityLocation" yaml:"geoproximityLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#health_check_id TfRecordsExclusive#health_check_id}.
	// Experimental.
	HealthCheckId *string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#multi_value_answer TfRecordsExclusive#multi_value_answer}.
	// Experimental.
	MultiValueAnswer interface{} `field:"optional" json:"multiValueAnswer" yaml:"multiValueAnswer"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#region TfRecordsExclusive#region}.
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// resource_records block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#resource_records TfRecordsExclusive#resource_records}
	// Experimental.
	ResourceRecords interface{} `field:"optional" json:"resourceRecords" yaml:"resourceRecords"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#set_identifier TfRecordsExclusive#set_identifier}.
	// Experimental.
	SetIdentifier *string `field:"optional" json:"setIdentifier" yaml:"setIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#traffic_policy_instance_id TfRecordsExclusive#traffic_policy_instance_id}.
	// Experimental.
	TrafficPolicyInstanceId *string `field:"optional" json:"trafficPolicyInstanceId" yaml:"trafficPolicyInstanceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#ttl TfRecordsExclusive#ttl}.
	// Experimental.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#type TfRecordsExclusive#type}.
	// Experimental.
	Type *string `field:"optional" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_records_exclusive#weight TfRecordsExclusive#weight}.
	// Experimental.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

