package wafclassicregional


// Experimental.
type AwsRule_PredicateProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rule#data_id AwsRule#data_id}.
	// Experimental.
	DataId *string `field:"required" json:"dataId" yaml:"dataId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rule#negated AwsRule#negated}.
	// Experimental.
	Negated interface{} `field:"required" json:"negated" yaml:"negated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/wafregional_rule#type AwsRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

