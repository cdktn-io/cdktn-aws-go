package endusermessagingsms

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKeywordConfig struct {
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
	// Keyword to configure. 1-30 characters, upper-case, and cannot start or end with a space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_keyword#keyword AwsKeyword#keyword}
	// Experimental.
	Keyword *string `field:"required" json:"keyword" yaml:"keyword"`
	// Message to send when the keyword is received.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_keyword#keyword_message AwsKeyword#keyword_message}
	// Experimental.
	KeywordMessage *string `field:"required" json:"keywordMessage" yaml:"keywordMessage"`
	// ARN of the origination identity (phone number or pool) to attach the keyword to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_keyword#origination_identity_arn AwsKeyword#origination_identity_arn}
	// Experimental.
	OriginationIdentityArn *string `field:"required" json:"originationIdentityArn" yaml:"originationIdentityArn"`
	// Action to perform when the keyword is received.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_keyword#keyword_action AwsKeyword#keyword_action}
	// Experimental.
	KeywordAction *string `field:"optional" json:"keywordAction" yaml:"keywordAction"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/pinpointsmsvoicev2_keyword#region AwsKeyword#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

