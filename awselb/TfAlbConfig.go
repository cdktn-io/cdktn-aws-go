package awselb

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfAlbConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#access_logs TfAlb#access_logs}
	// Experimental.
	AccessLogs *TfAlb_AccessLogsProperty `field:"optional" json:"accessLogs" yaml:"accessLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#client_keep_alive TfAlb#client_keep_alive}.
	// Experimental.
	ClientKeepAlive *float64 `field:"optional" json:"clientKeepAlive" yaml:"clientKeepAlive"`
	// connection_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#connection_logs TfAlb#connection_logs}
	// Experimental.
	ConnectionLogs *TfAlb_ConnectionLogsProperty `field:"optional" json:"connectionLogs" yaml:"connectionLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#customer_owned_ipv4_pool TfAlb#customer_owned_ipv4_pool}.
	// Experimental.
	CustomerOwnedIpv4Pool *string `field:"optional" json:"customerOwnedIpv4Pool" yaml:"customerOwnedIpv4Pool"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#desync_mitigation_mode TfAlb#desync_mitigation_mode}.
	// Experimental.
	DesyncMitigationMode *string `field:"optional" json:"desyncMitigationMode" yaml:"desyncMitigationMode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#dns_record_client_routing_policy TfAlb#dns_record_client_routing_policy}.
	// Experimental.
	DnsRecordClientRoutingPolicy *string `field:"optional" json:"dnsRecordClientRoutingPolicy" yaml:"dnsRecordClientRoutingPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#drop_invalid_header_fields TfAlb#drop_invalid_header_fields}.
	// Experimental.
	DropInvalidHeaderFields interface{} `field:"optional" json:"dropInvalidHeaderFields" yaml:"dropInvalidHeaderFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#enable_cross_zone_load_balancing TfAlb#enable_cross_zone_load_balancing}.
	// Experimental.
	EnableCrossZoneLoadBalancing interface{} `field:"optional" json:"enableCrossZoneLoadBalancing" yaml:"enableCrossZoneLoadBalancing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#enable_deletion_protection TfAlb#enable_deletion_protection}.
	// Experimental.
	EnableDeletionProtection interface{} `field:"optional" json:"enableDeletionProtection" yaml:"enableDeletionProtection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#enable_http2 TfAlb#enable_http2}.
	// Experimental.
	EnableHttp2 interface{} `field:"optional" json:"enableHttp2" yaml:"enableHttp2"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#enable_prefix_for_ipv6_source_nat TfAlb#enable_prefix_for_ipv6_source_nat}.
	// Experimental.
	EnablePrefixForIpv6SourceNat *string `field:"optional" json:"enablePrefixForIpv6SourceNat" yaml:"enablePrefixForIpv6SourceNat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#enable_tls_version_and_cipher_suite_headers TfAlb#enable_tls_version_and_cipher_suite_headers}.
	// Experimental.
	EnableTlsVersionAndCipherSuiteHeaders interface{} `field:"optional" json:"enableTlsVersionAndCipherSuiteHeaders" yaml:"enableTlsVersionAndCipherSuiteHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#enable_waf_fail_open TfAlb#enable_waf_fail_open}.
	// Experimental.
	EnableWafFailOpen interface{} `field:"optional" json:"enableWafFailOpen" yaml:"enableWafFailOpen"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#enable_xff_client_port TfAlb#enable_xff_client_port}.
	// Experimental.
	EnableXffClientPort interface{} `field:"optional" json:"enableXffClientPort" yaml:"enableXffClientPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#enable_zonal_shift TfAlb#enable_zonal_shift}.
	// Experimental.
	EnableZonalShift interface{} `field:"optional" json:"enableZonalShift" yaml:"enableZonalShift"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#enforce_security_group_inbound_rules_on_private_link_traffic TfAlb#enforce_security_group_inbound_rules_on_private_link_traffic}.
	// Experimental.
	EnforceSecurityGroupInboundRulesOnPrivateLinkTraffic *string `field:"optional" json:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic" yaml:"enforceSecurityGroupInboundRulesOnPrivateLinkTraffic"`
	// health_check_logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#health_check_logs TfAlb#health_check_logs}
	// Experimental.
	HealthCheckLogs *TfAlb_HealthCheckLogsProperty `field:"optional" json:"healthCheckLogs" yaml:"healthCheckLogs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#id TfAlb#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#idle_timeout TfAlb#idle_timeout}.
	// Experimental.
	IdleTimeout *float64 `field:"optional" json:"idleTimeout" yaml:"idleTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#internal TfAlb#internal}.
	// Experimental.
	Internal interface{} `field:"optional" json:"internal" yaml:"internal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#ip_address_type TfAlb#ip_address_type}.
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// ipam_pools block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#ipam_pools TfAlb#ipam_pools}
	// Experimental.
	IpamPools *TfAlb_IpamPoolsProperty `field:"optional" json:"ipamPools" yaml:"ipamPools"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#load_balancer_type TfAlb#load_balancer_type}.
	// Experimental.
	LoadBalancerType *string `field:"optional" json:"loadBalancerType" yaml:"loadBalancerType"`
	// minimum_load_balancer_capacity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#minimum_load_balancer_capacity TfAlb#minimum_load_balancer_capacity}
	// Experimental.
	MinimumLoadBalancerCapacity *TfAlb_MinimumLoadBalancerCapacityProperty `field:"optional" json:"minimumLoadBalancerCapacity" yaml:"minimumLoadBalancerCapacity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#name TfAlb#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#name_prefix TfAlb#name_prefix}.
	// Experimental.
	NamePrefix *string `field:"optional" json:"namePrefix" yaml:"namePrefix"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#preserve_host_header TfAlb#preserve_host_header}.
	// Experimental.
	PreserveHostHeader interface{} `field:"optional" json:"preserveHostHeader" yaml:"preserveHostHeader"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#region TfAlb#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#secondary_ips_auto_assigned_per_subnet TfAlb#secondary_ips_auto_assigned_per_subnet}.
	// Experimental.
	SecondaryIpsAutoAssignedPerSubnet *float64 `field:"optional" json:"secondaryIpsAutoAssignedPerSubnet" yaml:"secondaryIpsAutoAssignedPerSubnet"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#security_groups TfAlb#security_groups}.
	// Experimental.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
	// subnet_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#subnet_mapping TfAlb#subnet_mapping}
	// Experimental.
	SubnetMapping interface{} `field:"optional" json:"subnetMapping" yaml:"subnetMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#subnets TfAlb#subnets}.
	// Experimental.
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#tags TfAlb#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#tags_all TfAlb#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#timeouts TfAlb#timeouts}
	// Experimental.
	Timeouts *TfAlb_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb#xff_header_processing_mode TfAlb#xff_header_processing_mode}.
	// Experimental.
	XffHeaderProcessingMode *string `field:"optional" json:"xffHeaderProcessingMode" yaml:"xffHeaderProcessingMode"`
}

