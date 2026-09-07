package ses

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsReceiptRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#name AwsReceiptRule#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#rule_set_name AwsReceiptRule#rule_set_name}.
	// Experimental.
	RuleSetName *string `field:"required" json:"ruleSetName" yaml:"ruleSetName"`
	// add_header_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#add_header_action AwsReceiptRule#add_header_action}
	// Experimental.
	AddHeaderAction interface{} `field:"optional" json:"addHeaderAction" yaml:"addHeaderAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#after AwsReceiptRule#after}.
	// Experimental.
	After *string `field:"optional" json:"after" yaml:"after"`
	// bounce_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#bounce_action AwsReceiptRule#bounce_action}
	// Experimental.
	BounceAction interface{} `field:"optional" json:"bounceAction" yaml:"bounceAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#enabled AwsReceiptRule#enabled}.
	// Experimental.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#id AwsReceiptRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// lambda_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#lambda_action AwsReceiptRule#lambda_action}
	// Experimental.
	LambdaAction interface{} `field:"optional" json:"lambdaAction" yaml:"lambdaAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#recipients AwsReceiptRule#recipients}.
	// Experimental.
	Recipients *[]*string `field:"optional" json:"recipients" yaml:"recipients"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#region AwsReceiptRule#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// s3_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#s3_action AwsReceiptRule#s3_action}
	// Experimental.
	S3Action interface{} `field:"optional" json:"s3Action" yaml:"s3Action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#scan_enabled AwsReceiptRule#scan_enabled}.
	// Experimental.
	ScanEnabled interface{} `field:"optional" json:"scanEnabled" yaml:"scanEnabled"`
	// sns_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#sns_action AwsReceiptRule#sns_action}
	// Experimental.
	SnsAction interface{} `field:"optional" json:"snsAction" yaml:"snsAction"`
	// stop_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#stop_action AwsReceiptRule#stop_action}
	// Experimental.
	StopAction interface{} `field:"optional" json:"stopAction" yaml:"stopAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#tls_policy AwsReceiptRule#tls_policy}.
	// Experimental.
	TlsPolicy *string `field:"optional" json:"tlsPolicy" yaml:"tlsPolicy"`
	// workmail_action block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ses_receipt_rule#workmail_action AwsReceiptRule#workmail_action}
	// Experimental.
	WorkmailAction interface{} `field:"optional" json:"workmailAction" yaml:"workmailAction"`
}

