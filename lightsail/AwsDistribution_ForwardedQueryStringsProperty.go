package lightsail


// Experimental.
type AwsDistribution_ForwardedQueryStringsProperty struct {
	// Indicates whether the distribution forwards and caches based on query strings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#option AwsDistribution#option}
	// Experimental.
	Option interface{} `field:"optional" json:"option" yaml:"option"`
	// The specific query strings that the distribution forwards to the origin.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lightsail_distribution#query_strings_allowed_list AwsDistribution#query_strings_allowed_list}
	// Experimental.
	QueryStringsAllowedList *[]*string `field:"optional" json:"queryStringsAllowedList" yaml:"queryStringsAllowedList"`
}

