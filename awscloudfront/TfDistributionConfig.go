package awscloudfront

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDistributionConfig struct {
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
	// default_cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#default_cache_behavior TfDistribution#default_cache_behavior}
	// Experimental.
	DefaultCacheBehavior *TfDistribution_DefaultCacheBehaviorProperty `field:"required" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#enabled TfDistribution#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// origin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#origin TfDistribution#origin}
	// Experimental.
	Origin interface{} `field:"required" json:"origin" yaml:"origin"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#restrictions TfDistribution#restrictions}
	// Experimental.
	Restrictions *TfDistribution_RestrictionsProperty `field:"required" json:"restrictions" yaml:"restrictions"`
	// viewer_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#viewer_certificate TfDistribution#viewer_certificate}
	// Experimental.
	ViewerCertificate *TfDistribution_ViewerCertificateProperty `field:"required" json:"viewerCertificate" yaml:"viewerCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#aliases TfDistribution#aliases}.
	// Experimental.
	Aliases *[]*string `field:"optional" json:"aliases" yaml:"aliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#anycast_ip_list_id TfDistribution#anycast_ip_list_id}.
	// Experimental.
	AnycastIpListId *string `field:"optional" json:"anycastIpListId" yaml:"anycastIpListId"`
	// cache_tag_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#cache_tag_config TfDistribution#cache_tag_config}
	// Experimental.
	CacheTagConfig *TfDistribution_CacheTagConfigProperty `field:"optional" json:"cacheTagConfig" yaml:"cacheTagConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#comment TfDistribution#comment}.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// connection_function_association block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#connection_function_association TfDistribution#connection_function_association}
	// Experimental.
	ConnectionFunctionAssociation *TfDistribution_ConnectionFunctionAssociationProperty `field:"optional" json:"connectionFunctionAssociation" yaml:"connectionFunctionAssociation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#continuous_deployment_policy_id TfDistribution#continuous_deployment_policy_id}.
	// Experimental.
	ContinuousDeploymentPolicyId *string `field:"optional" json:"continuousDeploymentPolicyId" yaml:"continuousDeploymentPolicyId"`
	// custom_error_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#custom_error_response TfDistribution#custom_error_response}
	// Experimental.
	CustomErrorResponse interface{} `field:"optional" json:"customErrorResponse" yaml:"customErrorResponse"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#default_root_object TfDistribution#default_root_object}.
	// Experimental.
	DefaultRootObject *string `field:"optional" json:"defaultRootObject" yaml:"defaultRootObject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#http_version TfDistribution#http_version}.
	// Experimental.
	HttpVersion *string `field:"optional" json:"httpVersion" yaml:"httpVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#id TfDistribution#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#is_ipv6_enabled TfDistribution#is_ipv6_enabled}.
	// Experimental.
	IsIpv6Enabled interface{} `field:"optional" json:"isIpv6Enabled" yaml:"isIpv6Enabled"`
	// logging_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#logging_config TfDistribution#logging_config}
	// Experimental.
	LoggingConfig *TfDistribution_LoggingConfigProperty `field:"optional" json:"loggingConfig" yaml:"loggingConfig"`
	// ordered_cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#ordered_cache_behavior TfDistribution#ordered_cache_behavior}
	// Experimental.
	OrderedCacheBehavior interface{} `field:"optional" json:"orderedCacheBehavior" yaml:"orderedCacheBehavior"`
	// origin_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#origin_group TfDistribution#origin_group}
	// Experimental.
	OriginGroup interface{} `field:"optional" json:"originGroup" yaml:"originGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#price_class TfDistribution#price_class}.
	// Experimental.
	PriceClass *string `field:"optional" json:"priceClass" yaml:"priceClass"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#retain_on_delete TfDistribution#retain_on_delete}.
	// Experimental.
	RetainOnDelete interface{} `field:"optional" json:"retainOnDelete" yaml:"retainOnDelete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#staging TfDistribution#staging}.
	// Experimental.
	Staging interface{} `field:"optional" json:"staging" yaml:"staging"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#tags TfDistribution#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#tags_all TfDistribution#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// viewer_mtls_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#viewer_mtls_config TfDistribution#viewer_mtls_config}
	// Experimental.
	ViewerMtlsConfig *TfDistribution_ViewerMtlsConfigProperty `field:"optional" json:"viewerMtlsConfig" yaml:"viewerMtlsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#wait_for_deployment TfDistribution#wait_for_deployment}.
	// Experimental.
	WaitForDeployment interface{} `field:"optional" json:"waitForDeployment" yaml:"waitForDeployment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#web_acl_id TfDistribution#web_acl_id}.
	// Experimental.
	WebAclId *string `field:"optional" json:"webAclId" yaml:"webAclId"`
}

