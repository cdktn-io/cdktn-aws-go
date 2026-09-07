package wafclassic


// Experimental.
type AwsRule_PredicatesProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rule#data_id AwsRule#data_id}.
	// Experimental.
	DataId *string `field:"required" json:"dataId" yaml:"dataId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rule#negated AwsRule#negated}.
	// Experimental.
	Negated interface{} `field:"required" json:"negated" yaml:"negated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/waf_rule#type AwsRule#type}.
	// Experimental.
	Type *string `field:"required" json:"type" yaml:"type"`
}

