package awsvpnsitetosite

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsVpnConnectionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#customer_gateway_id AwsVpnConnection#customer_gateway_id}.
	// Experimental.
	CustomerGatewayId *string `field:"required" json:"customerGatewayId" yaml:"customerGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#type AwsVpnConnection#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#enable_acceleration AwsVpnConnection#enable_acceleration}.
	// Experimental.
	EnableAcceleration interface{} `field:"optional" json:"enableAcceleration" yaml:"enableAcceleration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#id AwsVpnConnection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#local_ipv4_network_cidr AwsVpnConnection#local_ipv4_network_cidr}.
	// Experimental.
	LocalIpv4NetworkCidr *string `field:"optional" json:"localIpv4NetworkCidr" yaml:"localIpv4NetworkCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#local_ipv6_network_cidr AwsVpnConnection#local_ipv6_network_cidr}.
	// Experimental.
	LocalIpv6NetworkCidr *string `field:"optional" json:"localIpv6NetworkCidr" yaml:"localIpv6NetworkCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#outside_ip_address_type AwsVpnConnection#outside_ip_address_type}.
	// Experimental.
	OutsideIpAddressType *string `field:"optional" json:"outsideIpAddressType" yaml:"outsideIpAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#preshared_key_storage AwsVpnConnection#preshared_key_storage}.
	// Experimental.
	PresharedKeyStorage *string `field:"optional" json:"presharedKeyStorage" yaml:"presharedKeyStorage"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#region AwsVpnConnection#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#remote_ipv4_network_cidr AwsVpnConnection#remote_ipv4_network_cidr}.
	// Experimental.
	RemoteIpv4NetworkCidr *string `field:"optional" json:"remoteIpv4NetworkCidr" yaml:"remoteIpv4NetworkCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#remote_ipv6_network_cidr AwsVpnConnection#remote_ipv6_network_cidr}.
	// Experimental.
	RemoteIpv6NetworkCidr *string `field:"optional" json:"remoteIpv6NetworkCidr" yaml:"remoteIpv6NetworkCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#static_routes_only AwsVpnConnection#static_routes_only}.
	// Experimental.
	StaticRoutesOnly interface{} `field:"optional" json:"staticRoutesOnly" yaml:"staticRoutesOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tags AwsVpnConnection#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tags_all AwsVpnConnection#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#transit_gateway_id AwsVpnConnection#transit_gateway_id}.
	// Experimental.
	TransitGatewayId *string `field:"optional" json:"transitGatewayId" yaml:"transitGatewayId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#transport_transit_gateway_attachment_id AwsVpnConnection#transport_transit_gateway_attachment_id}.
	// Experimental.
	TransportTransitGatewayAttachmentId *string `field:"optional" json:"transportTransitGatewayAttachmentId" yaml:"transportTransitGatewayAttachmentId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_dpd_timeout_action AwsVpnConnection#tunnel1_dpd_timeout_action}.
	// Experimental.
	Tunnel1DpdTimeoutAction *string `field:"optional" json:"tunnel1DpdTimeoutAction" yaml:"tunnel1DpdTimeoutAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_dpd_timeout_seconds AwsVpnConnection#tunnel1_dpd_timeout_seconds}.
	// Experimental.
	Tunnel1DpdTimeoutSeconds *float64 `field:"optional" json:"tunnel1DpdTimeoutSeconds" yaml:"tunnel1DpdTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_enable_tunnel_lifecycle_control AwsVpnConnection#tunnel1_enable_tunnel_lifecycle_control}.
	// Experimental.
	Tunnel1EnableTunnelLifecycleControl interface{} `field:"optional" json:"tunnel1EnableTunnelLifecycleControl" yaml:"tunnel1EnableTunnelLifecycleControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_ike_versions AwsVpnConnection#tunnel1_ike_versions}.
	// Experimental.
	Tunnel1IkeVersions *[]*string `field:"optional" json:"tunnel1IkeVersions" yaml:"tunnel1IkeVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_inside_cidr AwsVpnConnection#tunnel1_inside_cidr}.
	// Experimental.
	Tunnel1InsideCidr *string `field:"optional" json:"tunnel1InsideCidr" yaml:"tunnel1InsideCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_inside_ipv6_cidr AwsVpnConnection#tunnel1_inside_ipv6_cidr}.
	// Experimental.
	Tunnel1InsideIpv6Cidr *string `field:"optional" json:"tunnel1InsideIpv6Cidr" yaml:"tunnel1InsideIpv6Cidr"`
	// tunnel1_log_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_log_options AwsVpnConnection#tunnel1_log_options}
	// Experimental.
	Tunnel1LogOptions *AwsVpnConnection_Tunnel1LogOptionsProperty `field:"optional" json:"tunnel1LogOptions" yaml:"tunnel1LogOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_phase1_dh_group_numbers AwsVpnConnection#tunnel1_phase1_dh_group_numbers}.
	// Experimental.
	Tunnel1Phase1DhGroupNumbers *[]*float64 `field:"optional" json:"tunnel1Phase1DhGroupNumbers" yaml:"tunnel1Phase1DhGroupNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_phase1_encryption_algorithms AwsVpnConnection#tunnel1_phase1_encryption_algorithms}.
	// Experimental.
	Tunnel1Phase1EncryptionAlgorithms *[]*string `field:"optional" json:"tunnel1Phase1EncryptionAlgorithms" yaml:"tunnel1Phase1EncryptionAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_phase1_integrity_algorithms AwsVpnConnection#tunnel1_phase1_integrity_algorithms}.
	// Experimental.
	Tunnel1Phase1IntegrityAlgorithms *[]*string `field:"optional" json:"tunnel1Phase1IntegrityAlgorithms" yaml:"tunnel1Phase1IntegrityAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_phase1_lifetime_seconds AwsVpnConnection#tunnel1_phase1_lifetime_seconds}.
	// Experimental.
	Tunnel1Phase1LifetimeSeconds *float64 `field:"optional" json:"tunnel1Phase1LifetimeSeconds" yaml:"tunnel1Phase1LifetimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_phase2_dh_group_numbers AwsVpnConnection#tunnel1_phase2_dh_group_numbers}.
	// Experimental.
	Tunnel1Phase2DhGroupNumbers *[]*float64 `field:"optional" json:"tunnel1Phase2DhGroupNumbers" yaml:"tunnel1Phase2DhGroupNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_phase2_encryption_algorithms AwsVpnConnection#tunnel1_phase2_encryption_algorithms}.
	// Experimental.
	Tunnel1Phase2EncryptionAlgorithms *[]*string `field:"optional" json:"tunnel1Phase2EncryptionAlgorithms" yaml:"tunnel1Phase2EncryptionAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_phase2_integrity_algorithms AwsVpnConnection#tunnel1_phase2_integrity_algorithms}.
	// Experimental.
	Tunnel1Phase2IntegrityAlgorithms *[]*string `field:"optional" json:"tunnel1Phase2IntegrityAlgorithms" yaml:"tunnel1Phase2IntegrityAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_phase2_lifetime_seconds AwsVpnConnection#tunnel1_phase2_lifetime_seconds}.
	// Experimental.
	Tunnel1Phase2LifetimeSeconds *float64 `field:"optional" json:"tunnel1Phase2LifetimeSeconds" yaml:"tunnel1Phase2LifetimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_preshared_key AwsVpnConnection#tunnel1_preshared_key}.
	// Experimental.
	Tunnel1PresharedKey *string `field:"optional" json:"tunnel1PresharedKey" yaml:"tunnel1PresharedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_rekey_fuzz_percentage AwsVpnConnection#tunnel1_rekey_fuzz_percentage}.
	// Experimental.
	Tunnel1RekeyFuzzPercentage *float64 `field:"optional" json:"tunnel1RekeyFuzzPercentage" yaml:"tunnel1RekeyFuzzPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_rekey_margin_time_seconds AwsVpnConnection#tunnel1_rekey_margin_time_seconds}.
	// Experimental.
	Tunnel1RekeyMarginTimeSeconds *float64 `field:"optional" json:"tunnel1RekeyMarginTimeSeconds" yaml:"tunnel1RekeyMarginTimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_replay_window_size AwsVpnConnection#tunnel1_replay_window_size}.
	// Experimental.
	Tunnel1ReplayWindowSize *float64 `field:"optional" json:"tunnel1ReplayWindowSize" yaml:"tunnel1ReplayWindowSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel1_startup_action AwsVpnConnection#tunnel1_startup_action}.
	// Experimental.
	Tunnel1StartupAction *string `field:"optional" json:"tunnel1StartupAction" yaml:"tunnel1StartupAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_dpd_timeout_action AwsVpnConnection#tunnel2_dpd_timeout_action}.
	// Experimental.
	Tunnel2DpdTimeoutAction *string `field:"optional" json:"tunnel2DpdTimeoutAction" yaml:"tunnel2DpdTimeoutAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_dpd_timeout_seconds AwsVpnConnection#tunnel2_dpd_timeout_seconds}.
	// Experimental.
	Tunnel2DpdTimeoutSeconds *float64 `field:"optional" json:"tunnel2DpdTimeoutSeconds" yaml:"tunnel2DpdTimeoutSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_enable_tunnel_lifecycle_control AwsVpnConnection#tunnel2_enable_tunnel_lifecycle_control}.
	// Experimental.
	Tunnel2EnableTunnelLifecycleControl interface{} `field:"optional" json:"tunnel2EnableTunnelLifecycleControl" yaml:"tunnel2EnableTunnelLifecycleControl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_ike_versions AwsVpnConnection#tunnel2_ike_versions}.
	// Experimental.
	Tunnel2IkeVersions *[]*string `field:"optional" json:"tunnel2IkeVersions" yaml:"tunnel2IkeVersions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_inside_cidr AwsVpnConnection#tunnel2_inside_cidr}.
	// Experimental.
	Tunnel2InsideCidr *string `field:"optional" json:"tunnel2InsideCidr" yaml:"tunnel2InsideCidr"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_inside_ipv6_cidr AwsVpnConnection#tunnel2_inside_ipv6_cidr}.
	// Experimental.
	Tunnel2InsideIpv6Cidr *string `field:"optional" json:"tunnel2InsideIpv6Cidr" yaml:"tunnel2InsideIpv6Cidr"`
	// tunnel2_log_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_log_options AwsVpnConnection#tunnel2_log_options}
	// Experimental.
	Tunnel2LogOptions *AwsVpnConnection_Tunnel2LogOptionsProperty `field:"optional" json:"tunnel2LogOptions" yaml:"tunnel2LogOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_phase1_dh_group_numbers AwsVpnConnection#tunnel2_phase1_dh_group_numbers}.
	// Experimental.
	Tunnel2Phase1DhGroupNumbers *[]*float64 `field:"optional" json:"tunnel2Phase1DhGroupNumbers" yaml:"tunnel2Phase1DhGroupNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_phase1_encryption_algorithms AwsVpnConnection#tunnel2_phase1_encryption_algorithms}.
	// Experimental.
	Tunnel2Phase1EncryptionAlgorithms *[]*string `field:"optional" json:"tunnel2Phase1EncryptionAlgorithms" yaml:"tunnel2Phase1EncryptionAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_phase1_integrity_algorithms AwsVpnConnection#tunnel2_phase1_integrity_algorithms}.
	// Experimental.
	Tunnel2Phase1IntegrityAlgorithms *[]*string `field:"optional" json:"tunnel2Phase1IntegrityAlgorithms" yaml:"tunnel2Phase1IntegrityAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_phase1_lifetime_seconds AwsVpnConnection#tunnel2_phase1_lifetime_seconds}.
	// Experimental.
	Tunnel2Phase1LifetimeSeconds *float64 `field:"optional" json:"tunnel2Phase1LifetimeSeconds" yaml:"tunnel2Phase1LifetimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_phase2_dh_group_numbers AwsVpnConnection#tunnel2_phase2_dh_group_numbers}.
	// Experimental.
	Tunnel2Phase2DhGroupNumbers *[]*float64 `field:"optional" json:"tunnel2Phase2DhGroupNumbers" yaml:"tunnel2Phase2DhGroupNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_phase2_encryption_algorithms AwsVpnConnection#tunnel2_phase2_encryption_algorithms}.
	// Experimental.
	Tunnel2Phase2EncryptionAlgorithms *[]*string `field:"optional" json:"tunnel2Phase2EncryptionAlgorithms" yaml:"tunnel2Phase2EncryptionAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_phase2_integrity_algorithms AwsVpnConnection#tunnel2_phase2_integrity_algorithms}.
	// Experimental.
	Tunnel2Phase2IntegrityAlgorithms *[]*string `field:"optional" json:"tunnel2Phase2IntegrityAlgorithms" yaml:"tunnel2Phase2IntegrityAlgorithms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_phase2_lifetime_seconds AwsVpnConnection#tunnel2_phase2_lifetime_seconds}.
	// Experimental.
	Tunnel2Phase2LifetimeSeconds *float64 `field:"optional" json:"tunnel2Phase2LifetimeSeconds" yaml:"tunnel2Phase2LifetimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_preshared_key AwsVpnConnection#tunnel2_preshared_key}.
	// Experimental.
	Tunnel2PresharedKey *string `field:"optional" json:"tunnel2PresharedKey" yaml:"tunnel2PresharedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_rekey_fuzz_percentage AwsVpnConnection#tunnel2_rekey_fuzz_percentage}.
	// Experimental.
	Tunnel2RekeyFuzzPercentage *float64 `field:"optional" json:"tunnel2RekeyFuzzPercentage" yaml:"tunnel2RekeyFuzzPercentage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_rekey_margin_time_seconds AwsVpnConnection#tunnel2_rekey_margin_time_seconds}.
	// Experimental.
	Tunnel2RekeyMarginTimeSeconds *float64 `field:"optional" json:"tunnel2RekeyMarginTimeSeconds" yaml:"tunnel2RekeyMarginTimeSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_replay_window_size AwsVpnConnection#tunnel2_replay_window_size}.
	// Experimental.
	Tunnel2ReplayWindowSize *float64 `field:"optional" json:"tunnel2ReplayWindowSize" yaml:"tunnel2ReplayWindowSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel2_startup_action AwsVpnConnection#tunnel2_startup_action}.
	// Experimental.
	Tunnel2StartupAction *string `field:"optional" json:"tunnel2StartupAction" yaml:"tunnel2StartupAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel_bandwidth AwsVpnConnection#tunnel_bandwidth}.
	// Experimental.
	TunnelBandwidth *string `field:"optional" json:"tunnelBandwidth" yaml:"tunnelBandwidth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#tunnel_inside_ip_version AwsVpnConnection#tunnel_inside_ip_version}.
	// Experimental.
	TunnelInsideIpVersion *string `field:"optional" json:"tunnelInsideIpVersion" yaml:"tunnelInsideIpVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#vpn_concentrator_id AwsVpnConnection#vpn_concentrator_id}.
	// Experimental.
	VpnConcentratorId *string `field:"optional" json:"vpnConcentratorId" yaml:"vpnConcentratorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpn_connection#vpn_gateway_id AwsVpnConnection#vpn_gateway_id}.
	// Experimental.
	VpnGatewayId *string `field:"optional" json:"vpnGatewayId" yaml:"vpnGatewayId"`
}

