package cloudfront

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMultitenantDistributionConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#comment AwsMultitenantDistribution#comment}.
	// Experimental.
	Comment *string `field:"required" json:"comment" yaml:"comment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#enabled AwsMultitenantDistribution#enabled}.
	// Experimental.
	Enabled interface{} `field:"required" json:"enabled" yaml:"enabled"`
	// active_trusted_key_groups block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#active_trusted_key_groups AwsMultitenantDistribution#active_trusted_key_groups}
	// Experimental.
	ActiveTrustedKeyGroups interface{} `field:"optional" json:"activeTrustedKeyGroups" yaml:"activeTrustedKeyGroups"`
	// cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#cache_behavior AwsMultitenantDistribution#cache_behavior}
	// Experimental.
	CacheBehavior interface{} `field:"optional" json:"cacheBehavior" yaml:"cacheBehavior"`
	// custom_error_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#custom_error_response AwsMultitenantDistribution#custom_error_response}
	// Experimental.
	CustomErrorResponse interface{} `field:"optional" json:"customErrorResponse" yaml:"customErrorResponse"`
	// default_cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#default_cache_behavior AwsMultitenantDistribution#default_cache_behavior}
	// Experimental.
	DefaultCacheBehavior interface{} `field:"optional" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#default_root_object AwsMultitenantDistribution#default_root_object}.
	// Experimental.
	DefaultRootObject *string `field:"optional" json:"defaultRootObject" yaml:"defaultRootObject"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#http_version AwsMultitenantDistribution#http_version}.
	// Experimental.
	HttpVersion *string `field:"optional" json:"httpVersion" yaml:"httpVersion"`
	// origin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin AwsMultitenantDistribution#origin}
	// Experimental.
	Origin interface{} `field:"optional" json:"origin" yaml:"origin"`
	// origin_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_group AwsMultitenantDistribution#origin_group}
	// Experimental.
	OriginGroup interface{} `field:"optional" json:"originGroup" yaml:"originGroup"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#restrictions AwsMultitenantDistribution#restrictions}
	// Experimental.
	Restrictions interface{} `field:"optional" json:"restrictions" yaml:"restrictions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#tags AwsMultitenantDistribution#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// tenant_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#tenant_config AwsMultitenantDistribution#tenant_config}
	// Experimental.
	TenantConfig interface{} `field:"optional" json:"tenantConfig" yaml:"tenantConfig"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#timeouts AwsMultitenantDistribution#timeouts}
	// Experimental.
	Timeouts *AwsMultitenantDistribution_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// viewer_certificate block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#viewer_certificate AwsMultitenantDistribution#viewer_certificate}
	// Experimental.
	ViewerCertificate interface{} `field:"optional" json:"viewerCertificate" yaml:"viewerCertificate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#web_acl_id AwsMultitenantDistribution#web_acl_id}.
	// Experimental.
	WebAclId *string `field:"optional" json:"webAclId" yaml:"webAclId"`
}

