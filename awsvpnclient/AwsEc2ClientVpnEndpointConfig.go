package awsvpnclient

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEc2ClientVpnEndpointConfig struct {
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
	// authentication_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#authentication_options AwsEc2ClientVpnEndpoint#authentication_options}
	// Experimental.
	AuthenticationOptions interface{} `field:"required" json:"authenticationOptions" yaml:"authenticationOptions"`
	// connection_log_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#connection_log_options AwsEc2ClientVpnEndpoint#connection_log_options}
	// Experimental.
	ConnectionLogOptions *AwsEc2ClientVpnEndpoint_ConnectionLogOptionsProperty `field:"required" json:"connectionLogOptions" yaml:"connectionLogOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#server_certificate_arn AwsEc2ClientVpnEndpoint#server_certificate_arn}.
	// Experimental.
	ServerCertificateArn *string `field:"required" json:"serverCertificateArn" yaml:"serverCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#client_cidr_block AwsEc2ClientVpnEndpoint#client_cidr_block}.
	// Experimental.
	ClientCidrBlock *string `field:"optional" json:"clientCidrBlock" yaml:"clientCidrBlock"`
	// client_connect_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#client_connect_options AwsEc2ClientVpnEndpoint#client_connect_options}
	// Experimental.
	ClientConnectOptions *AwsEc2ClientVpnEndpoint_ClientConnectOptionsProperty `field:"optional" json:"clientConnectOptions" yaml:"clientConnectOptions"`
	// client_login_banner_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#client_login_banner_options AwsEc2ClientVpnEndpoint#client_login_banner_options}
	// Experimental.
	ClientLoginBannerOptions *AwsEc2ClientVpnEndpoint_ClientLoginBannerOptionsProperty `field:"optional" json:"clientLoginBannerOptions" yaml:"clientLoginBannerOptions"`
	// client_route_enforcement_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#client_route_enforcement_options AwsEc2ClientVpnEndpoint#client_route_enforcement_options}
	// Experimental.
	ClientRouteEnforcementOptions *AwsEc2ClientVpnEndpoint_ClientRouteEnforcementOptionsProperty `field:"optional" json:"clientRouteEnforcementOptions" yaml:"clientRouteEnforcementOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#description AwsEc2ClientVpnEndpoint#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#disconnect_on_session_timeout AwsEc2ClientVpnEndpoint#disconnect_on_session_timeout}.
	// Experimental.
	DisconnectOnSessionTimeout interface{} `field:"optional" json:"disconnectOnSessionTimeout" yaml:"disconnectOnSessionTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#dns_servers AwsEc2ClientVpnEndpoint#dns_servers}.
	// Experimental.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#endpoint_ip_address_type AwsEc2ClientVpnEndpoint#endpoint_ip_address_type}.
	// Experimental.
	EndpointIpAddressType *string `field:"optional" json:"endpointIpAddressType" yaml:"endpointIpAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#id AwsEc2ClientVpnEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#region AwsEc2ClientVpnEndpoint#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#security_group_ids AwsEc2ClientVpnEndpoint#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#self_service_portal AwsEc2ClientVpnEndpoint#self_service_portal}.
	// Experimental.
	SelfServicePortal *string `field:"optional" json:"selfServicePortal" yaml:"selfServicePortal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#session_timeout_hours AwsEc2ClientVpnEndpoint#session_timeout_hours}.
	// Experimental.
	SessionTimeoutHours *float64 `field:"optional" json:"sessionTimeoutHours" yaml:"sessionTimeoutHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#split_tunnel AwsEc2ClientVpnEndpoint#split_tunnel}.
	// Experimental.
	SplitTunnel interface{} `field:"optional" json:"splitTunnel" yaml:"splitTunnel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#tags AwsEc2ClientVpnEndpoint#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#tags_all AwsEc2ClientVpnEndpoint#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#traffic_ip_address_type AwsEc2ClientVpnEndpoint#traffic_ip_address_type}.
	// Experimental.
	TrafficIpAddressType *string `field:"optional" json:"trafficIpAddressType" yaml:"trafficIpAddressType"`
	// transit_gateway_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#transit_gateway_configuration AwsEc2ClientVpnEndpoint#transit_gateway_configuration}
	// Experimental.
	TransitGatewayConfiguration *AwsEc2ClientVpnEndpoint_TransitGatewayConfigurationProperty `field:"optional" json:"transitGatewayConfiguration" yaml:"transitGatewayConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#transport_protocol AwsEc2ClientVpnEndpoint#transport_protocol}.
	// Experimental.
	TransportProtocol *string `field:"optional" json:"transportProtocol" yaml:"transportProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#vpc_id AwsEc2ClientVpnEndpoint#vpc_id}.
	// Experimental.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#vpn_port AwsEc2ClientVpnEndpoint#vpn_port}.
	// Experimental.
	VpnPort *float64 `field:"optional" json:"vpnPort" yaml:"vpnPort"`
}

