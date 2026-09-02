package awssesv2

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfEmailIdentityMailFromAttributesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_email_identity_mail_from_attributes#email_identity TfEmailIdentityMailFromAttributes#email_identity}.
	// Experimental.
	EmailIdentity *string `field:"required" json:"emailIdentity" yaml:"emailIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_email_identity_mail_from_attributes#behavior_on_mx_failure TfEmailIdentityMailFromAttributes#behavior_on_mx_failure}.
	// Experimental.
	BehaviorOnMxFailure *string `field:"optional" json:"behaviorOnMxFailure" yaml:"behaviorOnMxFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_email_identity_mail_from_attributes#id TfEmailIdentityMailFromAttributes#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_email_identity_mail_from_attributes#mail_from_domain TfEmailIdentityMailFromAttributes#mail_from_domain}.
	// Experimental.
	MailFromDomain *string `field:"optional" json:"mailFromDomain" yaml:"mailFromDomain"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/sesv2_email_identity_mail_from_attributes#region TfEmailIdentityMailFromAttributes#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
}

