package cloudfront


// Experimental.
type AwsDistribution_OriginProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#domain_name AwsDistribution#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#origin_id AwsDistribution#origin_id}.
	// Experimental.
	OriginId *string `field:"required" json:"originId" yaml:"originId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#connection_attempts AwsDistribution#connection_attempts}.
	// Experimental.
	ConnectionAttempts *float64 `field:"optional" json:"connectionAttempts" yaml:"connectionAttempts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#connection_timeout AwsDistribution#connection_timeout}.
	// Experimental.
	ConnectionTimeout *float64 `field:"optional" json:"connectionTimeout" yaml:"connectionTimeout"`
	// custom_header block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#custom_header AwsDistribution#custom_header}
	// Experimental.
	CustomHeader interface{} `field:"optional" json:"customHeader" yaml:"customHeader"`
	// custom_origin_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#custom_origin_config AwsDistribution#custom_origin_config}
	// Experimental.
	CustomOriginConfig *AwsDistribution_CustomOriginConfigProperty `field:"optional" json:"customOriginConfig" yaml:"customOriginConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#origin_access_control_id AwsDistribution#origin_access_control_id}.
	// Experimental.
	OriginAccessControlId *string `field:"optional" json:"originAccessControlId" yaml:"originAccessControlId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#origin_path AwsDistribution#origin_path}.
	// Experimental.
	OriginPath *string `field:"optional" json:"originPath" yaml:"originPath"`
	// origin_shield block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#origin_shield AwsDistribution#origin_shield}
	// Experimental.
	OriginShield *AwsDistribution_OriginShieldProperty `field:"optional" json:"originShield" yaml:"originShield"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#response_completion_timeout AwsDistribution#response_completion_timeout}.
	// Experimental.
	ResponseCompletionTimeout *float64 `field:"optional" json:"responseCompletionTimeout" yaml:"responseCompletionTimeout"`
	// s3_origin_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#s3_origin_config AwsDistribution#s3_origin_config}
	// Experimental.
	S3OriginConfig *AwsDistribution_S3OriginConfigProperty `field:"optional" json:"s3OriginConfig" yaml:"s3OriginConfig"`
	// vpc_origin_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_distribution#vpc_origin_config AwsDistribution#vpc_origin_config}
	// Experimental.
	VpcOriginConfig *AwsDistribution_VpcOriginConfigProperty `field:"optional" json:"vpcOriginConfig" yaml:"vpcOriginConfig"`
}

