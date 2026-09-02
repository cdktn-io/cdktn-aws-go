package awsroute53resolver

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfResolverEndpointConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#direction TfResolverEndpoint#direction}.
	// Experimental.
	Direction *string `field:"required" json:"direction" yaml:"direction"`
	// ip_address block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#ip_address TfResolverEndpoint#ip_address}
	// Experimental.
	IpAddress interface{} `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#security_group_ids TfResolverEndpoint#security_group_ids}.
	// Experimental.
	SecurityGroupIds *[]*string `field:"required" json:"securityGroupIds" yaml:"securityGroupIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#id TfResolverEndpoint#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#name TfResolverEndpoint#name}.
	// Experimental.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#protocols TfResolverEndpoint#protocols}.
	// Experimental.
	Protocols *[]*string `field:"optional" json:"protocols" yaml:"protocols"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#region TfResolverEndpoint#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#resolver_endpoint_type TfResolverEndpoint#resolver_endpoint_type}.
	// Experimental.
	ResolverEndpointType *string `field:"optional" json:"resolverEndpointType" yaml:"resolverEndpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#rni_enhanced_metrics_enabled TfResolverEndpoint#rni_enhanced_metrics_enabled}.
	// Experimental.
	RniEnhancedMetricsEnabled interface{} `field:"optional" json:"rniEnhancedMetricsEnabled" yaml:"rniEnhancedMetricsEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#tags TfResolverEndpoint#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#tags_all TfResolverEndpoint#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#target_name_server_metrics_enabled TfResolverEndpoint#target_name_server_metrics_enabled}.
	// Experimental.
	TargetNameServerMetricsEnabled interface{} `field:"optional" json:"targetNameServerMetricsEnabled" yaml:"targetNameServerMetricsEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53_resolver_endpoint#timeouts TfResolverEndpoint#timeouts}
	// Experimental.
	Timeouts *TfResolverEndpoint_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

