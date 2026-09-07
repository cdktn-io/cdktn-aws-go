package wafclassic


// Experimental.
type AwsRateBasedRule_PredicatesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rate_based_rule#data_id AwsRateBasedRule#data_id}.
	// Experimental.
	DataId *string `field:"required" json:"dataId" yaml:"dataId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rate_based_rule#negated AwsRateBasedRule#negated}.
	// Experimental.
	Negated interface{} `field:"required" json:"negated" yaml:"negated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rate_based_rule#type AwsRateBasedRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

