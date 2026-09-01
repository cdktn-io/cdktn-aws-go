package awscloudfront

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCloudfrontMultitenantDistributionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#comment AwsCloudfrontMultitenantDistribution#comment}.
	// Experimental.
	Comment *string `field:"required" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#enabled AwsCloudfrontMultitenantDistribution#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// active_trusted_key_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#active_trusted_key_groups AwsCloudfrontMultitenantDistribution#active_trusted_key_groups}
	// Experimental.
	ActiveTrustedKeyGroups interface{} `field:"optional" json:"activeTrustedKeyGroups" yaml:"activeTrustedKeyGroups"`
	// cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#cache_behavior AwsCloudfrontMultitenantDistribution#cache_behavior}
	// Experimental.
	CacheBehavior interface{} `field:"optional" json:"cacheBehavior" yaml:"cacheBehavior"`
	// custom_error_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#custom_error_response AwsCloudfrontMultitenantDistribution#custom_error_response}
	// Experimental.
	CustomErrorResponse interface{} `field:"optional" json:"customErrorResponse" yaml:"customErrorResponse"`
	// default_cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#default_cache_behavior AwsCloudfrontMultitenantDistribution#default_cache_behavior}
	// Experimental.
	DefaultCacheBehavior interface{} `field:"optional" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#default_root_object AwsCloudfrontMultitenantDistribution#default_root_object}.
	// Experimental.
	DefaultRootObject *string `field:"optional" json:"defaultRootObject" yaml:"defaultRootObject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#http_version AwsCloudfrontMultitenantDistribution#http_version}.
	// Experimental.
	HttpVersion *string `field:"optional" json:"httpVersion" yaml:"httpVersion"`
	// origin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin AwsCloudfrontMultitenantDistribution#origin}
	// Experimental.
	Origin interface{} `field:"optional" json:"origin" yaml:"origin"`
	// origin_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_group AwsCloudfrontMultitenantDistribution#origin_group}
	// Experimental.
	OriginGroup interface{} `field:"optional" json:"originGroup" yaml:"originGroup"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#restrictions AwsCloudfrontMultitenantDistribution#restrictions}
	// Experimental.
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#tags AwsCloudfrontMultitenantDistribution#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// tenant_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#tenant_config AwsCloudfrontMultitenantDistribution#tenant_config}
	// Experimental.
	TenantConfig interface{} `field:"optional" json:"tenantConfig" yaml:"tenantConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#timeouts AwsCloudfrontMultitenantDistribution#timeouts}
	// Experimental.
	Timeouts *AwsCloudfrontMultitenantDistribution_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// viewer_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#viewer_certificate AwsCloudfrontMultitenantDistribution#viewer_certificate}
	// Experimental.
	ViewerCertificate interface{} `field:"optional" json:"viewerCertificate" yaml:"viewerCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#web_acl_id AwsCloudfrontMultitenantDistribution#web_acl_id}.
	// Experimental.
	WebAclId *string `field:"optional" json:"webAclId" yaml:"webAclId"`
}

