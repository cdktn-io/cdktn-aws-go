package route53domains

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRegisteredDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#domain_name AwsRegisteredDomain#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// admin_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#admin_contact AwsRegisteredDomain#admin_contact}
	// Experimental.
	AdminContact *AwsRegisteredDomain_AdminContactProperty `field:"optional" json:"adminContact" yaml:"adminContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#admin_privacy AwsRegisteredDomain#admin_privacy}.
	// Experimental.
	AdminPrivacy interface{} `field:"optional" json:"adminPrivacy" yaml:"adminPrivacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#auto_renew AwsRegisteredDomain#auto_renew}.
	// Experimental.
	AutoRenew interface{} `field:"optional" json:"autoRenew" yaml:"autoRenew"`
	// billing_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#billing_contact AwsRegisteredDomain#billing_contact}
	// Experimental.
	BillingContact *AwsRegisteredDomain_BillingContactProperty `field:"optional" json:"billingContact" yaml:"billingContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#billing_privacy AwsRegisteredDomain#billing_privacy}.
	// Experimental.
	BillingPrivacy interface{} `field:"optional" json:"billingPrivacy" yaml:"billingPrivacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#id AwsRegisteredDomain#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	// Experimental.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// name_server block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#name_server AwsRegisteredDomain#name_server}
	// Experimental.
	NameServer interface{} `field:"optional" json:"nameServer" yaml:"nameServer"`
	// registrant_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#registrant_contact AwsRegisteredDomain#registrant_contact}
	// Experimental.
	RegistrantContact *AwsRegisteredDomain_RegistrantContactProperty `field:"optional" json:"registrantContact" yaml:"registrantContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#registrant_privacy AwsRegisteredDomain#registrant_privacy}.
	// Experimental.
	RegistrantPrivacy interface{} `field:"optional" json:"registrantPrivacy" yaml:"registrantPrivacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#tags AwsRegisteredDomain#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#tags_all AwsRegisteredDomain#tags_all}.
	// Experimental.
	TagsAll *map[string]*string `field:"optional" json:"tagsAll" yaml:"tagsAll"`
	// tech_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#tech_contact AwsRegisteredDomain#tech_contact}
	// Experimental.
	TechContact *AwsRegisteredDomain_TechContactProperty `field:"optional" json:"techContact" yaml:"techContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#tech_privacy AwsRegisteredDomain#tech_privacy}.
	// Experimental.
	TechPrivacy interface{} `field:"optional" json:"techPrivacy" yaml:"techPrivacy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#timeouts AwsRegisteredDomain#timeouts}
	// Experimental.
	Timeouts *AwsRegisteredDomain_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_registered_domain#transfer_lock AwsRegisteredDomain#transfer_lock}.
	// Experimental.
	TransferLock interface{} `field:"optional" json:"transferLock" yaml:"transferLock"`
}

