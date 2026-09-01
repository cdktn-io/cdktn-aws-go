package awselb

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLbConfig struct {
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
	// access_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#access_logs AwsLb#access_logs}
	// Experimental.
	AccessLogs *AwsLb_AccessLogsProperty `field:"optional" json:"accessLogs" yaml:"accessLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#client_keep_alive AwsLb#client_keep_alive}.
	// Experimental.
	ClientKeepAlive *float64 `field:"optional" json:"clientKeepAlive" yaml:"clientKeepAlive"`
	// connection_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#connection_logs AwsLb#connection_logs}
	// Experimental.
	ConnectionLogs *AwsLb_ConnectionLogsProperty `field:"optional" json:"connectionLogs" yaml:"connectionLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#customer_owned_ipv4_pool AwsLb#customer_owned_ipv4_pool}.
	// Experimental.
	CustomerOwnedIpv4Pool *string `field:"optional" json:"customerOwnedIpv4Pool" yaml:"customerOwnedIpv4Pool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#desync_mitigation_mode AwsLb#desync_mitigation_mode}.
	// Experimental.
	DesyncMitigationMode *string `field:"optional" json:"desyncMitigationMode" yaml:"desyncMitigationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#dns_record_client_routing_policy AwsLb#dns_record_client_routing_policy}.
	// Experimental.
	DnsRecordClientRoutingPolicy *string `field:"optional" json:"dnsRecordClientRoutingPolicy" yaml:"dnsRecordClientRoutingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#drop_invalid_header_fields AwsLb#drop_invalid_header_fields}.
	// Experimental.
	DropInvalidHeaderFields interface{} `field:"optional" json:"dropInvalidHeaderFields" yaml:"dropInvalidHeaderFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#enable_cross_zone_load_balancing AwsLb#enable_cross_zone_load_balancing}.
	// Experimental.
	EnableCrossZoneLoadBalancing interface{} `field:"optional" json:"enableCrossZoneLoadBalancing" yaml:"enableCrossZoneLoadBalancing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#enable_deletion_protection AwsLb#enable_deletion_protection}.
	// Experimental.
	EnableDeletionProtection interface{} `field:"optional" json:"enableDeletionProtection" yaml:"enableDeletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#enable_http2 AwsLb#enable_http2}.
	// Experimental.
	EnableHttp2 interface{} `field:"optional" json:"enableHttp2" yaml:"enableHttp2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#enable_prefix_for_ipv6_source_nat AwsLb#enable_prefix_for_ipv6_source_nat}.
	// Experimental.
	EnablePrefixForIpv6SourceNat *string `field:"optional" json:"enablePrefixForIpv6SourceNat" yaml:"enablePrefixForIpv6SourceNat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#enable_tls_version_and_cipher_suite_headers AwsLb#enable_tls_version_and_cipher_suite_headers}.
	// Experimental.
	EnableTlsVersionAndCipherSuiteHeaders interface{} `field:"optional" json:"enableTlsVersionAndCipherSuiteHeaders" yaml:"enableTlsVersionAndCipherSuiteHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#enable_waf_fail_open AwsLb#enable_waf_fail_open}.
	// Experimental.
	EnableWafFailOpen interface{} `field:"optional" json:"enableWafFailOpen" yaml:"enableWafFailOpen"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#enable_xff_client_port AwsLb#enable_xff_client_port}.
	// Experimental.
	EnableXffClientPort interface{} `field:"optional" json:"enableXffClientPort" yaml:"enableXffClientPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#enable_zonal_shift AwsLb#enable_zonal_shift}.
	// Experimental.
	EnableZonalShift interface{} `field:"optional" json:"enableZonalShift" yaml:"enableZonalShift"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#enforce_security_group_inbound_rules_on_private_link_traffic AwsLb#enforce_security_group_inbound_rules_on_private_link_traffic}.
	// Experimental.
	EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic *string `field:"optional" json:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic" yaml:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic"`
	// health_check_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#health_check_logs AwsLb#health_check_logs}
	// Experimental.
	HealthCheckLogs *AwsLb_HealthCheckLogsProperty `field:"optional" json:"healthCheckLogs" yaml:"healthCheckLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#id AwsLb#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#idle_timeout AwsLb#idle_timeout}.
	// Experimental.
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#internal AwsLb#internal}.
	// Experimental.
	Internal interface{} `field:"optional" json:"internal" yaml:"internal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#ip_address_type AwsLb#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// ipam_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#ipam_pools AwsLb#ipam_pools}
	// Experimental.
	IpamPools *AwsLb_IpamPoolsProperty `field:"optional" json:"ipamPools" yaml:"ipamPools"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#load_balancer_type AwsLb#load_balancer_type}.
	// Experimental.
	LoadBalancerType *string `field:"optional" json:"loadBalancerType" yaml:"loadBalancerType"`
	// minimum_load_balancer_capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#minimum_load_balancer_capacity AwsLb#minimum_load_balancer_capacity}
	// Experimental.
	MinimumLoadBalancerCapacity *AwsLb_MinimumLoadBalancerCapacityProperty `field:"optional" json:"minimumLoadBalancerCapacity" yaml:"minimumLoadBalancerCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#name AwsLb#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#name_prefix AwsLb#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#preserve_host_header AwsLb#preserve_host_header}.
	// Experimental.
	PreserveHostHeader interface{} `field:"optional" json:"preserveHostHeader" yaml:"preserveHostHeader"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#region AwsLb#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#secondary_ips_auto_assigned_per_subnet AwsLb#secondary_ips_auto_assigned_per_subnet}.
	// Experimental.
	SecondaryIpsAutoAssignedPerSubnet *float64 `field:"optional" json:"secondaryIpsAutoAssignedPerSubnet" yaml:"secondaryIpsAutoAssignedPerSubnet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#security_groups AwsLb#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// subnet_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#subnet_mapping AwsLb#subnet_mapping}
	// Experimental.
	SubnetMapping interface{} `field:"optional" json:"subnetMapping" yaml:"subnetMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#subnets AwsLb#subnets}.
	// Experimental.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#tags AwsLb#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#tags_all AwsLb#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#timeouts AwsLb#timeouts}
	// Experimental.
	Timeouts *AwsLb_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lb#xff_header_processing_mode AwsLb#xff_header_processing_mode}.
	// Experimental.
	XffHeaderProcessingMode *string `field:"optional" json:"xffHeaderProcessingMode" yaml:"xffHeaderProcessingMode"`
}

