package awsroute53domains

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDomainConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#domain_name TfDomain#domain_name}.
	// Experimental.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// admin_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#admin_contact TfDomain#admin_contact}
	// Experimental.
	AdminContact interface{} `field:"optional" json:"adminContact" yaml:"adminContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#admin_privacy TfDomain#admin_privacy}.
	// Experimental.
	AdminPrivacy interface{} `field:"optional" json:"adminPrivacy" yaml:"adminPrivacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#auto_renew TfDomain#auto_renew}.
	// Experimental.
	AutoRenew interface{} `field:"optional" json:"autoRenew" yaml:"autoRenew"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#billing_contact TfDomain#billing_contact}.
	// Experimental.
	BillingContact interface{} `field:"optional" json:"billingContact" yaml:"billingContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#billing_privacy TfDomain#billing_privacy}.
	// Experimental.
	BillingPrivacy interface{} `field:"optional" json:"billingPrivacy" yaml:"billingPrivacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#duration_in_years TfDomain#duration_in_years}.
	// Experimental.
	DurationInYears *float64 `field:"optional" json:"durationInYears" yaml:"durationInYears"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#name_server TfDomain#name_server}.
	// Experimental.
	NameServer interface{} `field:"optional" json:"nameServer" yaml:"nameServer"`
	// registrant_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#registrant_contact TfDomain#registrant_contact}
	// Experimental.
	RegistrantContact interface{} `field:"optional" json:"registrantContact" yaml:"registrantContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#registrant_privacy TfDomain#registrant_privacy}.
	// Experimental.
	RegistrantPrivacy interface{} `field:"optional" json:"registrantPrivacy" yaml:"registrantPrivacy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#tags TfDomain#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// tech_contact block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#tech_contact TfDomain#tech_contact}
	// Experimental.
	TechContact interface{} `field:"optional" json:"techContact" yaml:"techContact"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#tech_privacy TfDomain#tech_privacy}.
	// Experimental.
	TechPrivacy interface{} `field:"optional" json:"techPrivacy" yaml:"techPrivacy"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#timeouts TfDomain#timeouts}
	// Experimental.
	Timeouts *TfDomain_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/route53domains_domain#transfer_lock TfDomain#transfer_lock}.
	// Experimental.
	TransferLock interface{} `field:"optional" json:"transferLock" yaml:"transferLock"`
}

