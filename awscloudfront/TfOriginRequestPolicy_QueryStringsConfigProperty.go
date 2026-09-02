package awscloudfront


// Experimental.
type TfOriginRequestPolicy_QueryStringsConfigProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_origin_request_policy#query_string_behavior TfOriginRequestPolicy#query_string_behavior}.
	// Experimental.
	QueryStringBehavior *string `field:"required" json:"queryStringBehavior" yaml:"queryStringBehavior"`
	// query_strings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/cloudfront_origin_request_policy#query_strings TfOriginRequestPolicy#query_strings}
	// Experimental.
	QueryStrings *TfOriginRequestPolicy_QueryStringsProperty `field:"optional" json:"queryStrings" yaml:"queryStrings"`
}

