package awsroute53

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRecordConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#name TfRecord#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#type TfRecord#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#zone_id TfRecord#zone_id}.
	// Experimental.
	ZoneId *string `field:"required" json:"zoneId" yaml:"zoneId"`
	// alias block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#alias TfRecord#alias}
	// Experimental.
	Alias *TfRecord_AliasProperty `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#allow_overwrite TfRecord#allow_overwrite}.
	// Experimental.
	AllowOverwrite interface{} `field:"optional" json:"allowOverwrite" yaml:"allowOverwrite"`
	// cidr_routing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#cidr_routing_policy TfRecord#cidr_routing_policy}
	// Experimental.
	CidrRoutingPolicy *TfRecord_CidrRoutingPolicyProperty `field:"optional" json:"cidrRoutingPolicy" yaml:"cidrRoutingPolicy"`
	// failover_routing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#failover_routing_policy TfRecord#failover_routing_policy}
	// Experimental.
	FailoverRoutingPolicy *TfRecord_FailoverRoutingPolicyProperty `field:"optional" json:"failoverRoutingPolicy" yaml:"failoverRoutingPolicy"`
	// geolocation_routing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#geolocation_routing_policy TfRecord#geolocation_routing_policy}
	// Experimental.
	GeolocationRoutingPolicy *TfRecord_GeolocationRoutingPolicyProperty `field:"optional" json:"geolocationRoutingPolicy" yaml:"geolocationRoutingPolicy"`
	// geoproximity_routing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#geoproximity_routing_policy TfRecord#geoproximity_routing_policy}
	// Experimental.
	GeoproximityRoutingPolicy *TfRecord_GeoproximityRoutingPolicyProperty `field:"optional" json:"geoproximityRoutingPolicy" yaml:"geoproximityRoutingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#health_check_id TfRecord#health_check_id}.
	// Experimental.
	HealthCheckId *string `field:"optional" json:"healthCheckId" yaml:"healthCheckId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#id TfRecord#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// latency_routing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#latency_routing_policy TfRecord#latency_routing_policy}
	// Experimental.
	LatencyRoutingPolicy *TfRecord_LatencyRoutingPolicyProperty `field:"optional" json:"latencyRoutingPolicy" yaml:"latencyRoutingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#multivalue_answer_routing_policy TfRecord#multivalue_answer_routing_policy}.
	// Experimental.
	MultivalueAnswerRoutingPolicy interface{} `field:"optional" json:"multivalueAnswerRoutingPolicy" yaml:"multivalueAnswerRoutingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#records TfRecord#records}.
	// Experimental.
	Records *[]*string `field:"optional" json:"records" yaml:"records"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#set_identifier TfRecord#set_identifier}.
	// Experimental.
	SetIdentifier *string `field:"optional" json:"setIdentifier" yaml:"setIdentifier"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#timeouts TfRecord#timeouts}
	// Experimental.
	Timeouts *TfRecord_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#ttl TfRecord#ttl}.
	// Experimental.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
	// weighted_routing_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_record#weighted_routing_policy TfRecord#weighted_routing_policy}
	// Experimental.
	WeightedRoutingPolicy *TfRecord_WeightedRoutingPolicyProperty `field:"optional" json:"weightedRoutingPolicy" yaml:"weightedRoutingPolicy"`
}

