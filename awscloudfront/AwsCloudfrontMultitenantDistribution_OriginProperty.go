package awscloudfront


// Experimental.
type AwsCloudfrontMultitenantDistribution_OriginProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#domain_name AwsCloudfrontMultitenantDistribution#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#id AwsCloudfrontMultitenantDistribution#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#connection_attempts AwsCloudfrontMultitenantDistribution#connection_attempts}.
	// Experimental.
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#connection_timeout AwsCloudfrontMultitenantDistribution#connection_timeout}.
	// Experimental.
	ConnectionTimeout *float64 `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// custom_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#custom_header AwsCloudfrontMultitenantDistribution#custom_header}
	// Experimental.
	CustomHeader interface{} `field:"optional" json:"customHeader" yaml:"customHeader"`
	// custom_origin_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#custom_origin_config AwsCloudfrontMultitenantDistribution#custom_origin_config}
	// Experimental.
	CustomOriginConfig interface{} `field:"optional" json:"customOriginConfig" yaml:"customOriginConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_access_control_id AwsCloudfrontMultitenantDistribution#origin_access_control_id}.
	// Experimental.
	OriginAccessControlId *string `field:"optional" json:"originAccessControlId" yaml:"originAccessControlId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_path AwsCloudfrontMultitenantDistribution#origin_path}.
	// Experimental.
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// origin_shield block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#origin_shield AwsCloudfrontMultitenantDistribution#origin_shield}
	// Experimental.
	OriginShield interface{} `field:"optional" json:"originShield" yaml:"originShield"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#response_completion_timeout AwsCloudfrontMultitenantDistribution#response_completion_timeout}.
	// Experimental.
	ResponseCompletionTimeout *float64 `field:"optional" json:"responseCompletionTimeout" yaml:"responseCompletionTimeout"`
	// vpc_origin_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_multitenant_distribution#vpc_origin_config AwsCloudfrontMultitenantDistribution#vpc_origin_config}
	// Experimental.
	VpcOriginConfig interface{} `field:"optional" json:"vpcOriginConfig" yaml:"vpcOriginConfig"`
}

