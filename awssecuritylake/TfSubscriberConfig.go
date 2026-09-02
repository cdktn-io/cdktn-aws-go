package awssecuritylake

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfSubscriberConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#access_type TfSubscriber#access_type}.
	// Experimental.
	AccessType *string `field:"optional" json:"accessType" yaml:"accessType"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#region TfSubscriber#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#source TfSubscriber#source}
	// Experimental.
	Source interface{} `field:"optional" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#subscriber_description TfSubscriber#subscriber_description}.
	// Experimental.
	SubscriberDescription *string `field:"optional" json:"subscriberDescription" yaml:"subscriberDescription"`
	// subscriber_identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#subscriber_identity TfSubscriber#subscriber_identity}
	// Experimental.
	SubscriberIdentity interface{} `field:"optional" json:"subscriberIdentity" yaml:"subscriberIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#subscriber_name TfSubscriber#subscriber_name}.
	// Experimental.
	SubscriberName *string `field:"optional" json:"subscriberName" yaml:"subscriberName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#tags TfSubscriber#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/securitylake_subscriber#timeouts TfSubscriber#timeouts}
	// Experimental.
	Timeouts *TfSubscriber_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

