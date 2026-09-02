package awsvpc

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEndpointServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint_service#acceptance_required TfEndpointService#acceptance_required}.
	// Experimental.
	AcceptanceRequired interface{} `field:"required" json:"acceptanceRequired" yaml:"acceptanceRequired"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint_service#allowed_principals TfEndpointService#allowed_principals}.
	// Experimental.
	AllowedPrincipals *[]*string `field:"optional" json:"allowedPrincipals" yaml:"allowedPrincipals"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint_service#gateway_load_balancer_arns TfEndpointService#gateway_load_balancer_arns}.
	// Experimental.
	GatewayLoadBalancerArns *[]*string `field:"optional" json:"gatewayLoadBalancerArns" yaml:"gatewayLoadBalancerArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint_service#id TfEndpointService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint_service#network_load_balancer_arns TfEndpointService#network_load_balancer_arns}.
	// Experimental.
	NetworkLoadBalancerArns *[]*string `field:"optional" json:"networkLoadBalancerArns" yaml:"networkLoadBalancerArns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint_service#private_dns_name TfEndpointService#private_dns_name}.
	// Experimental.
	PrivateDnsName *string `field:"optional" json:"privateDnsName" yaml:"privateDnsName"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint_service#region TfEndpointService#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint_service#supported_ip_address_types TfEndpointService#supported_ip_address_types}.
	// Experimental.
	SupportedIpAddressTypes *[]*string `field:"optional" json:"supportedIpAddressTypes" yaml:"supportedIpAddressTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint_service#supported_regions TfEndpointService#supported_regions}.
	// Experimental.
	SupportedRegions *[]*string `field:"optional" json:"supportedRegions" yaml:"supportedRegions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint_service#tags TfEndpointService#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint_service#tags_all TfEndpointService#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/vpc_endpoint_service#timeouts TfEndpointService#timeouts}
	// Experimental.
	Timeouts *TfEndpointService_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

