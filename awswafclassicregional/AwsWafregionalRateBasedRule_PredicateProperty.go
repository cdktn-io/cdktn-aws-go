package awswafclassicregional


// Experimental.
type AwsWafregionalRateBasedRule_PredicateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rate_based_rule#data_id AwsWafregionalRateBasedRule#data_id}.
	// Experimental.
	DataId *string `field:"required" json:"dataId" yaml:"dataId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rate_based_rule#negated AwsWafregionalRateBasedRule#negated}.
	// Experimental.
	Negated interface{} `field:"required" json:"negated" yaml:"negated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rate_based_rule#type AwsWafregionalRateBasedRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

