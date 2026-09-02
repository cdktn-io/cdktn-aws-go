package awselb

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTargetGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#connection_termination TfTargetGroup#connection_termination}.
	// Experimental.
	ConnectionTermination interface{} `field:"optional" json:"connectionTermination" yaml:"connectionTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#deregistration_delay TfTargetGroup#deregistration_delay}.
	// Experimental.
	DeregistrationDelay *string `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#health_check TfTargetGroup#health_check}
	// Experimental.
	HealthCheck *TfTargetGroup_HealthCheckProperty `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#id TfTargetGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#ip_address_type TfTargetGroup#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#lambda_multi_value_headers_enabled TfTargetGroup#lambda_multi_value_headers_enabled}.
	// Experimental.
	LambdaMultiValueHeadersEnabled interface{} `field:"optional" json:"lambdaMultiValueHeadersEnabled" yaml:"lambdaMultiValueHeadersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#load_balancing_algorithm_type TfTargetGroup#load_balancing_algorithm_type}.
	// Experimental.
	LoadBalancingAlgorithmType *string `field:"optional" json:"loadBalancingAlgorithmType" yaml:"loadBalancingAlgorithmType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#load_balancing_anomaly_mitigation TfTargetGroup#load_balancing_anomaly_mitigation}.
	// Experimental.
	LoadBalancingAnomalyMitigation *string `field:"optional" json:"loadBalancingAnomalyMitigation" yaml:"loadBalancingAnomalyMitigation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#load_balancing_cross_zone_enabled TfTargetGroup#load_balancing_cross_zone_enabled}.
	// Experimental.
	LoadBalancingCrossZoneEnabled *string `field:"optional" json:"loadBalancingCrossZoneEnabled" yaml:"loadBalancingCrossZoneEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#name TfTargetGroup#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#name_prefix TfTargetGroup#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#port TfTargetGroup#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#preserve_client_ip TfTargetGroup#preserve_client_ip}.
	// Experimental.
	PreserveClientIp *string `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#protocol TfTargetGroup#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#protocol_version TfTargetGroup#protocol_version}.
	// Experimental.
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#proxy_protocol_v2 TfTargetGroup#proxy_protocol_v2}.
	// Experimental.
	ProxyProtocolV2 interface{} `field:"optional" json:"proxyProtocolV2" yaml:"proxyProtocolV2"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#region TfTargetGroup#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#slow_start TfTargetGroup#slow_start}.
	// Experimental.
	SlowStart *float64 `field:"optional" json:"slowStart" yaml:"slowStart"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#stickiness TfTargetGroup#stickiness}
	// Experimental.
	Stickiness *TfTargetGroup_StickinessProperty `field:"optional" json:"stickiness" yaml:"stickiness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#tags TfTargetGroup#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#tags_all TfTargetGroup#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#target_control_port TfTargetGroup#target_control_port}.
	// Experimental.
	TargetControlPort *float64 `field:"optional" json:"targetControlPort" yaml:"targetControlPort"`
	// target_failover block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#target_failover TfTargetGroup#target_failover}
	// Experimental.
	TargetFailover interface{} `field:"optional" json:"targetFailover" yaml:"targetFailover"`
	// target_group_health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#target_group_health TfTargetGroup#target_group_health}
	// Experimental.
	TargetGroupHealth *TfTargetGroup_TargetGroupHealthProperty `field:"optional" json:"targetGroupHealth" yaml:"targetGroupHealth"`
	// target_health_state block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#target_health_state TfTargetGroup#target_health_state}
	// Experimental.
	TargetHealthState interface{} `field:"optional" json:"targetHealthState" yaml:"targetHealthState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#target_type TfTargetGroup#target_type}.
	// Experimental.
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#vpc_id TfTargetGroup#vpc_id}.
	// Experimental.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

