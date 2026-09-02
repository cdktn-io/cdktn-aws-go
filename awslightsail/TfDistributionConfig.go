package awslightsail

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
	// The bundle ID to use for the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#bundle_id TfDistribution#bundle_id}
	// Experimental.
	BundleId *string `field:"required" json:"bundleId" yaml:"bundleId"`
	// default_cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#default_cache_behavior TfDistribution#default_cache_behavior}
	// Experimental.
	DefaultCacheBehavior *TfDistribution_DefaultCacheBehaviorProperty `field:"required" json:"defaultCacheBehavior" yaml:"defaultCacheBehavior"`
	// The name of the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#name TfDistribution#name}
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// origin block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#origin TfDistribution#origin}
	// Experimental.
	Origin *TfDistribution_OriginProperty `field:"required" json:"origin" yaml:"origin"`
	// cache_behavior block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#cache_behavior TfDistribution#cache_behavior}
	// Experimental.
	CacheBehavior interface{} `field:"optional" json:"cacheBehavior" yaml:"cacheBehavior"`
	// cache_behavior_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#cache_behavior_settings TfDistribution#cache_behavior_settings}
	// Experimental.
	CacheBehaviorSettings *TfDistribution_CacheBehaviorSettingsProperty `field:"optional" json:"cacheBehaviorSettings" yaml:"cacheBehaviorSettings"`
	// The name of the SSL/TLS certificate attached to the distribution, if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#certificate_name TfDistribution#certificate_name}
	// Experimental.
	CertificateName *string `field:"optional" json:"certificateName" yaml:"certificateName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#id TfDistribution#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The IP address type of the distribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#ip_address_type TfDistribution#ip_address_type}
	// Experimental.
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
	// Indicates whether the distribution is enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#is_enabled TfDistribution#is_enabled}
	// Experimental.
	IsEnabled interface{} `field:"optional" json:"isEnabled" yaml:"isEnabled"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#region TfDistribution#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#tags TfDistribution#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#tags_all TfDistribution#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#timeouts TfDistribution#timeouts}
	// Experimental.
	Timeouts *TfDistribution_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

