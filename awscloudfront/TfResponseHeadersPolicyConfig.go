package awscloudfront

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfResponseHeadersPolicyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#name TfResponseHeadersPolicy#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#comment TfResponseHeadersPolicy#comment}.
	// Experimental.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
	// cors_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#cors_config TfResponseHeadersPolicy#cors_config}
	// Experimental.
	CorsConfig *TfResponseHeadersPolicy_CorsConfigProperty `field:"optional" json:"corsConfig" yaml:"corsConfig"`
	// custom_headers_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#custom_headers_config TfResponseHeadersPolicy#custom_headers_config}
	// Experimental.
	CustomHeadersConfig *TfResponseHeadersPolicy_CustomHeadersConfigProperty `field:"optional" json:"customHeadersConfig" yaml:"customHeadersConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#id TfResponseHeadersPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// remove_headers_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#remove_headers_config TfResponseHeadersPolicy#remove_headers_config}
	// Experimental.
	RemoveHeadersConfig *TfResponseHeadersPolicy_RemoveHeadersConfigProperty `field:"optional" json:"removeHeadersConfig" yaml:"removeHeadersConfig"`
	// security_headers_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#security_headers_config TfResponseHeadersPolicy#security_headers_config}
	// Experimental.
	SecurityHeadersConfig *TfResponseHeadersPolicy_SecurityHeadersConfigProperty `field:"optional" json:"securityHeadersConfig" yaml:"securityHeadersConfig"`
	// server_timing_headers_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_response_headers_policy#server_timing_headers_config TfResponseHeadersPolicy#server_timing_headers_config}
	// Experimental.
	ServerTimingHeadersConfig *TfResponseHeadersPolicy_ServerTimingHeadersConfigProperty `field:"optional" json:"serverTimingHeadersConfig" yaml:"serverTimingHeadersConfig"`
}

