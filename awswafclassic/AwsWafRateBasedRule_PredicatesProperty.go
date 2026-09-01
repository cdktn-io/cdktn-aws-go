package awswafclassic


// Experimental.
type AwsWafRateBasedRule_PredicatesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rate_based_rule#data_id AwsWafRateBasedRule#data_id}.
	// Experimental.
	DataId *string `field:"required" json:"dataId" yaml:"dataId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rate_based_rule#negated AwsWafRateBasedRule#negated}.
	// Experimental.
	Negated interface{} `field:"required" json:"negated" yaml:"negated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rate_based_rule#type AwsWafRateBasedRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

