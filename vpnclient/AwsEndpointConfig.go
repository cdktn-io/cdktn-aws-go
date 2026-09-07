package vpnclient

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#authentication_options AwsEndpoint#authentication_options}
	// Experimental.
	AuthenticationOptions interface{} `field:"required" json:"authenticationOptions" yaml:"authenticationOptions"`
	// connection_log_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#connection_log_options AwsEndpoint#connection_log_options}
	// Experimental.
	ConnectionLogOptions *AwsEndpoint_ConnectionLogOptionsProperty `field:"required" json:"connectionLogOptions" yaml:"connectionLogOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#server_certificate_arn AwsEndpoint#server_certificate_arn}.
	// Experimental.
	ServerCertificateArn *string `field:"required" json:"serverCertificateArn" yaml:"serverCertificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#client_cidr_block AwsEndpoint#client_cidr_block}.
	// Experimental.
	ClientCidrBlock *string `field:"optional" json:"clientCidrBlock" yaml:"clientCidrBlock"`
	// client_connect_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#client_connect_options AwsEndpoint#client_connect_options}
	// Experimental.
	ClientConnectOptions *AwsEndpoint_ClientConnectOptionsProperty `field:"optional" json:"clientConnectOptions" yaml:"clientConnectOptions"`
	// client_login_banner_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#client_login_banner_options AwsEndpoint#client_login_banner_options}
	// Experimental.
	ClientLoginBannerOptions *AwsEndpoint_ClientLoginBannerOptionsProperty `field:"optional" json:"clientLoginBannerOptions" yaml:"clientLoginBannerOptions"`
	// client_route_enforcement_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#client_route_enforcement_options AwsEndpoint#client_route_enforcement_options}
	// Experimental.
	ClientRouteEnforcementOptions *AwsEndpoint_ClientRouteEnforcementOptionsProperty `field:"optional" json:"clientRouteEnforcementOptions" yaml:"clientRouteEnforcementOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#description AwsEndpoint#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#disconnect_on_session_timeout AwsEndpoint#disconnect_on_session_timeout}.
	// Experimental.
	DisconnectOnSessionTimeout interface{} `field:"optional" json:"disconnectOnSessionTimeout" yaml:"disconnectOnSessionTimeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#dns_servers AwsEndpoint#dns_servers}.
	// Experimental.
	DnsServers *[]*string `field:"optional" json:"dnsServers" yaml:"dnsServers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#endpoint_ip_address_type AwsEndpoint#endpoint_ip_address_type}.
	// Experimental.
	EndpointIpAddressType *string `field:"optional" json:"endpointIpAddressType" yaml:"endpointIpAddressType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#id AwsEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#region AwsEndpoint#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#security_group_ids AwsEndpoint#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"optional" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#self_service_portal AwsEndpoint#self_service_portal}.
	// Experimental.
	SelfServicePortal *string `field:"optional" json:"selfServicePortal" yaml:"selfServicePortal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#session_timeout_hours AwsEndpoint#session_timeout_hours}.
	// Experimental.
	SessionTimeoutHours *float64 `field:"optional" json:"sessionTimeoutHours" yaml:"sessionTimeoutHours"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#split_tunnel AwsEndpoint#split_tunnel}.
	// Experimental.
	SplitTunnel interface{} `field:"optional" json:"splitTunnel" yaml:"splitTunnel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#tags AwsEndpoint#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#tags_all AwsEndpoint#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#traffic_ip_address_type AwsEndpoint#traffic_ip_address_type}.
	// Experimental.
	TrafficIpAddressType *string `field:"optional" json:"trafficIpAddressType" yaml:"trafficIpAddressType"`
	// transit_gateway_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#transit_gateway_configuration AwsEndpoint#transit_gateway_configuration}
	// Experimental.
	TransitGatewayConfiguration *AwsEndpoint_TransitGatewayConfigurationProperty `field:"optional" json:"transitGatewayConfiguration" yaml:"transitGatewayConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#transport_protocol AwsEndpoint#transport_protocol}.
	// Experimental.
	TransportProtocol *string `field:"optional" json:"transportProtocol" yaml:"transportProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#vpc_id AwsEndpoint#vpc_id}.
	// Experimental.
	VpcId *string `field:"optional" json:"vpcId" yaml:"vpcId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ec2_client_vpn_endpoint#vpn_port AwsEndpoint#vpn_port}.
	// Experimental.
	VpnPort *float64 `field:"optional" json:"vpnPort" yaml:"vpnPort"`
}

