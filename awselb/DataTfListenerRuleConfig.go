package awselb

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfListenerRuleConfig struct {
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
	// action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#action DataTfListenerRule#action}
	// Experimental.
	Action interface{} `field:"optional" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#arn DataTfListenerRule#arn}.
	// Experimental.
	Arn *string `field:"optional" json:"arn" yaml:"arn"`
	// condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#condition DataTfListenerRule#condition}
	// Experimental.
	Condition interface{} `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#listener_arn DataTfListenerRule#listener_arn}.
	// Experimental.
	ListenerArn *string `field:"optional" json:"listenerArn" yaml:"listenerArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#priority DataTfListenerRule#priority}.
	// Experimental.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#region DataTfListenerRule#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// transform block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/lb_listener_rule#transform DataTfListenerRule#transform}
	// Experimental.
	Transform interface{} `field:"optional" json:"transform" yaml:"transform"`
}

