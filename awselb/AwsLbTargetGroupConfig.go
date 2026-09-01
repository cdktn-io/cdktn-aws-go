package awselb

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLbTargetGroupConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#connection_termination AwsLbTargetGroup#connection_termination}.
	// Experimental.
	ConnectionTermination interface{} `field:"optional" json:"connectionTermination" yaml:"connectionTermination"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#deregistration_delay AwsLbTargetGroup#deregistration_delay}.
	// Experimental.
	DeregistrationDelay *string `field:"optional" json:"deregistrationDelay" yaml:"deregistrationDelay"`
	// health_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#health_check AwsLbTargetGroup#health_check}
	// Experimental.
	HealthCheck *AwsLbTargetGroup_HealthCheckProperty `field:"optional" json:"healthCheck" yaml:"healthCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#id AwsLbTargetGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#ip_address_type AwsLbTargetGroup#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#lambda_multi_value_headers_enabled AwsLbTargetGroup#lambda_multi_value_headers_enabled}.
	// Experimental.
	LambdaMultiValueHeadersEnabled interface{} `field:"optional" json:"lambdaMultiValueHeadersEnabled" yaml:"lambdaMultiValueHeadersEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#load_balancing_algorithm_type AwsLbTargetGroup#load_balancing_algorithm_type}.
	// Experimental.
	LoadBalancingAlgorithmType *string `field:"optional" json:"loadBalancingAlgorithmType" yaml:"loadBalancingAlgorithmType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#load_balancing_anomaly_mitigation AwsLbTargetGroup#load_balancing_anomaly_mitigation}.
	// Experimental.
	LoadBalancingAnomalyMitigation *string `field:"optional" json:"loadBalancingAnomalyMitigation" yaml:"loadBalancingAnomalyMitigation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#load_balancing_cross_zone_enabled AwsLbTargetGroup#load_balancing_cross_zone_enabled}.
	// Experimental.
	LoadBalancingCrossZoneEnabled *string `field:"optional" json:"loadBalancingCrossZoneEnabled" yaml:"loadBalancingCrossZoneEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#name AwsLbTargetGroup#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#name_prefix AwsLbTargetGroup#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#port AwsLbTargetGroup#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#preserve_client_ip AwsLbTargetGroup#preserve_client_ip}.
	// Experimental.
	PreserveClientIp *string `field:"optional" json:"preserveClientIp" yaml:"preserveClientIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#protocol AwsLbTargetGroup#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#protocol_version AwsLbTargetGroup#protocol_version}.
	// Experimental.
	ProtocolVersion *string `field:"optional" json:"protocolVersion" yaml:"protocolVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#proxy_protocol_v2 AwsLbTargetGroup#proxy_protocol_v2}.
	// Experimental.
	ProxyProtocolV2 interface{} `field:"optional" json:"proxyProtocolV2" yaml:"proxyProtocolV2"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#region AwsLbTargetGroup#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#slow_start AwsLbTargetGroup#slow_start}.
	// Experimental.
	SlowStart *float64 `field:"optional" json:"slowStart" yaml:"slowStart"`
	// stickiness block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#stickiness AwsLbTargetGroup#stickiness}
	// Experimental.
	Stickiness *AwsLbTargetGroup_StickinessProperty `field:"optional" json:"stickiness" yaml:"stickiness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#tags AwsLbTargetGroup#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#tags_all AwsLbTargetGroup#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#target_control_port AwsLbTargetGroup#target_control_port}.
	// Experimental.
	TargetControlPort *float64 `field:"optional" json:"targetControlPort" yaml:"targetControlPort"`
	// target_failover block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#target_failover AwsLbTargetGroup#target_failover}
	// Experimental.
	TargetFailover interface{} `field:"optional" json:"targetFailover" yaml:"targetFailover"`
	// target_group_health block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#target_group_health AwsLbTargetGroup#target_group_health}
	// Experimental.
	TargetGroupHealth *AwsLbTargetGroup_TargetGroupHealthProperty `field:"optional" json:"targetGroupHealth" yaml:"targetGroupHealth"`
	// target_health_state block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#target_health_state AwsLbTargetGroup#target_health_state}
	// Experimental.
	TargetHealthState interface{} `field:"optional" json:"targetHealthState" yaml:"targetHealthState"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#target_type AwsLbTargetGroup#target_type}.
	// Experimental.
	TargetType *string `field:"optional" json:"targetType" yaml:"targetType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb_target_group#vpc_id AwsLbTargetGroup#vpc_id}.
	// Experimental.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
}

