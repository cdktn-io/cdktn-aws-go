package elb

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAlbListenerConfig struct {
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
	// default_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#default_action AwsAlbListener#default_action}
	// Experimental.
	DefaultAction interface{} `field:"required" json:"defaultAction" yaml:"defaultAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#load_balancer_arn AwsAlbListener#load_balancer_arn}.
	// Experimental.
	LoadBalancerArn *string `field:"required" json:"loadBalancerArn" yaml:"loadBalancerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#alpn_policy AwsAlbListener#alpn_policy}.
	// Experimental.
	AlpnPolicy *string `field:"optional" json:"alpnPolicy" yaml:"alpnPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#certificate_arn AwsAlbListener#certificate_arn}.
	// Experimental.
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#id AwsAlbListener#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// mutual_authentication block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#mutual_authentication AwsAlbListener#mutual_authentication}
	// Experimental.
	MutualAuthentication *AwsAlbListener_MutualAuthenticationProperty `field:"optional" json:"mutualAuthentication" yaml:"mutualAuthentication"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#port AwsAlbListener#port}.
	// Experimental.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#protocol AwsAlbListener#protocol}.
	// Experimental.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#region AwsAlbListener#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_request_x_amzn_mtls_clientcert_header_name AwsAlbListener#routing_http_request_x_amzn_mtls_clientcert_header_name}.
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertHeaderName *string `field:"optional" json:"routingHttpRequestXAmznMtlsClientcertHeaderName" yaml:"routingHttpRequestXAmznMtlsClientcertHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_request_x_amzn_mtls_clientcert_issuer_header_name AwsAlbListener#routing_http_request_x_amzn_mtls_clientcert_issuer_header_name}.
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertIssuerHeaderName *string `field:"optional" json:"routingHttpRequestXAmznMtlsClientcertIssuerHeaderName" yaml:"routingHttpRequestXAmznMtlsClientcertIssuerHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_request_x_amzn_mtls_clientcert_leaf_header_name AwsAlbListener#routing_http_request_x_amzn_mtls_clientcert_leaf_header_name}.
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertLeafHeaderName *string `field:"optional" json:"routingHttpRequestXAmznMtlsClientcertLeafHeaderName" yaml:"routingHttpRequestXAmznMtlsClientcertLeafHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_request_x_amzn_mtls_clientcert_serial_number_header_name AwsAlbListener#routing_http_request_x_amzn_mtls_clientcert_serial_number_header_name}.
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName *string `field:"optional" json:"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName" yaml:"routingHttpRequestXAmznMtlsClientcertSerialNumberHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_request_x_amzn_mtls_clientcert_subject_header_name AwsAlbListener#routing_http_request_x_amzn_mtls_clientcert_subject_header_name}.
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertSubjectHeaderName *string `field:"optional" json:"routingHttpRequestXAmznMtlsClientcertSubjectHeaderName" yaml:"routingHttpRequestXAmznMtlsClientcertSubjectHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_request_x_amzn_mtls_clientcert_validity_header_name AwsAlbListener#routing_http_request_x_amzn_mtls_clientcert_validity_header_name}.
	// Experimental.
	RoutingHttpRequestXAmznMtlsClientcertValidityHeaderName *string `field:"optional" json:"routingHttpRequestXAmznMtlsClientcertValidityHeaderName" yaml:"routingHttpRequestXAmznMtlsClientcertValidityHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_request_x_amzn_tls_cipher_suite_header_name AwsAlbListener#routing_http_request_x_amzn_tls_cipher_suite_header_name}.
	// Experimental.
	RoutingHttpRequestXAmznTlsCipherSuiteHeaderName *string `field:"optional" json:"routingHttpRequestXAmznTlsCipherSuiteHeaderName" yaml:"routingHttpRequestXAmznTlsCipherSuiteHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_request_x_amzn_tls_version_header_name AwsAlbListener#routing_http_request_x_amzn_tls_version_header_name}.
	// Experimental.
	RoutingHttpRequestXAmznTlsVersionHeaderName *string `field:"optional" json:"routingHttpRequestXAmznTlsVersionHeaderName" yaml:"routingHttpRequestXAmznTlsVersionHeaderName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_response_access_control_allow_credentials_header_value AwsAlbListener#routing_http_response_access_control_allow_credentials_header_value}.
	// Experimental.
	RoutingHttpResponseAccessControlAllowCredentialsHeaderValue *string `field:"optional" json:"routingHttpResponseAccessControlAllowCredentialsHeaderValue" yaml:"routingHttpResponseAccessControlAllowCredentialsHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_response_access_control_allow_headers_header_value AwsAlbListener#routing_http_response_access_control_allow_headers_header_value}.
	// Experimental.
	RoutingHttpResponseAccessControlAllowHeadersHeaderValue *string `field:"optional" json:"routingHttpResponseAccessControlAllowHeadersHeaderValue" yaml:"routingHttpResponseAccessControlAllowHeadersHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_response_access_control_allow_methods_header_value AwsAlbListener#routing_http_response_access_control_allow_methods_header_value}.
	// Experimental.
	RoutingHttpResponseAccessControlAllowMethodsHeaderValue *string `field:"optional" json:"routingHttpResponseAccessControlAllowMethodsHeaderValue" yaml:"routingHttpResponseAccessControlAllowMethodsHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_response_access_control_allow_origin_header_value AwsAlbListener#routing_http_response_access_control_allow_origin_header_value}.
	// Experimental.
	RoutingHttpResponseAccessControlAllowOriginHeaderValue *string `field:"optional" json:"routingHttpResponseAccessControlAllowOriginHeaderValue" yaml:"routingHttpResponseAccessControlAllowOriginHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_response_access_control_expose_headers_header_value AwsAlbListener#routing_http_response_access_control_expose_headers_header_value}.
	// Experimental.
	RoutingHttpResponseAccessControlExposeHeadersHeaderValue *string `field:"optional" json:"routingHttpResponseAccessControlExposeHeadersHeaderValue" yaml:"routingHttpResponseAccessControlExposeHeadersHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_response_access_control_max_age_header_value AwsAlbListener#routing_http_response_access_control_max_age_header_value}.
	// Experimental.
	RoutingHttpResponseAccessControlMaxAgeHeaderValue *string `field:"optional" json:"routingHttpResponseAccessControlMaxAgeHeaderValue" yaml:"routingHttpResponseAccessControlMaxAgeHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_response_content_security_policy_header_value AwsAlbListener#routing_http_response_content_security_policy_header_value}.
	// Experimental.
	RoutingHttpResponseContentSecurityPolicyHeaderValue *string `field:"optional" json:"routingHttpResponseContentSecurityPolicyHeaderValue" yaml:"routingHttpResponseContentSecurityPolicyHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_response_server_enabled AwsAlbListener#routing_http_response_server_enabled}.
	// Experimental.
	RoutingHttpResponseServerEnabled interface{} `field:"optional" json:"routingHttpResponseServerEnabled" yaml:"routingHttpResponseServerEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_response_strict_transport_security_header_value AwsAlbListener#routing_http_response_strict_transport_security_header_value}.
	// Experimental.
	RoutingHttpResponseStrictTransportSecurityHeaderValue *string `field:"optional" json:"routingHttpResponseStrictTransportSecurityHeaderValue" yaml:"routingHttpResponseStrictTransportSecurityHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_response_x_content_type_options_header_value AwsAlbListener#routing_http_response_x_content_type_options_header_value}.
	// Experimental.
	RoutingHttpResponseXContentTypeOptionsHeaderValue *string `field:"optional" json:"routingHttpResponseXContentTypeOptionsHeaderValue" yaml:"routingHttpResponseXContentTypeOptionsHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#routing_http_response_x_frame_options_header_value AwsAlbListener#routing_http_response_x_frame_options_header_value}.
	// Experimental.
	RoutingHttpResponseXFrameOptionsHeaderValue *string `field:"optional" json:"routingHttpResponseXFrameOptionsHeaderValue" yaml:"routingHttpResponseXFrameOptionsHeaderValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#ssl_policy AwsAlbListener#ssl_policy}.
	// Experimental.
	SslPolicy *string `field:"optional" json:"sslPolicy" yaml:"sslPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#tags AwsAlbListener#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#tags_all AwsAlbListener#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#tcp_idle_timeout_seconds AwsAlbListener#tcp_idle_timeout_seconds}.
	// Experimental.
	TcpIdleTimeoutSeconds *float64 `field:"optional" json:"tcpIdleTimeoutSeconds" yaml:"tcpIdleTimeoutSeconds"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/alb_listener#timeouts AwsAlbListener#timeouts}
	// Experimental.
	Timeouts *AwsAlbListener_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

