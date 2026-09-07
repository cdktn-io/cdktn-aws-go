package savingsplans

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSavingsPlanConfig struct {
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
	// The hourly commitment, in USD.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/savingsplans_savings_plan#commitment AwsSavingsPlan#commitment}
	// Experimental.
	Commitment *string `field:"required" json:"commitment" yaml:"commitment"`
	// The unique ID of a Savings Plan offering.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/savingsplans_savings_plan#savings_plan_offering_id AwsSavingsPlan#savings_plan_offering_id}
	// Experimental.
	SavingsPlanOfferingId *string `field:"required" json:"savingsPlanOfferingId" yaml:"savingsPlanOfferingId"`
	// The time at which to purchase the Savings Plan, in UTC format (YYYY-MM-DDTHH:MM:SSZ).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/savingsplans_savings_plan#purchase_time AwsSavingsPlan#purchase_time}
	// Experimental.
	PurchaseTime *string `field:"optional" json:"purchaseTime" yaml:"purchaseTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/savingsplans_savings_plan#tags AwsSavingsPlan#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/savingsplans_savings_plan#timeouts AwsSavingsPlan#timeouts}
	// Experimental.
	Timeouts *AwsSavingsPlan_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
	// The up-front payment amount.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/savingsplans_savings_plan#upfront_payment_amount AwsSavingsPlan#upfront_payment_amount}
	// Experimental.
	UpfrontPaymentAmount *string `field:"optional" json:"upfrontPaymentAmount" yaml:"upfrontPaymentAmount"`
}

