package savingsplans

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsOfferingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#currencies DataAwsOfferings#currencies}.
	// Experimental.
	Currencies *[]*string `field:"optional" json:"currencies" yaml:"currencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#descriptions DataAwsOfferings#descriptions}.
	// Experimental.
	Descriptions *[]*string `field:"optional" json:"descriptions" yaml:"descriptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#durations DataAwsOfferings#durations}.
	// Experimental.
	Durations *[]*float64 `field:"optional" json:"durations" yaml:"durations"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#filter DataAwsOfferings#filter}
	// Experimental.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#offering_ids DataAwsOfferings#offering_ids}.
	// Experimental.
	OfferingIds *[]*string `field:"optional" json:"offeringIds" yaml:"offeringIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#operations DataAwsOfferings#operations}.
	// Experimental.
	Operations *[]*string `field:"optional" json:"operations" yaml:"operations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#payment_options DataAwsOfferings#payment_options}.
	// Experimental.
	PaymentOptions *[]*string `field:"optional" json:"paymentOptions" yaml:"paymentOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#plan_types DataAwsOfferings#plan_types}.
	// Experimental.
	PlanTypes *[]*string `field:"optional" json:"planTypes" yaml:"planTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#product_type DataAwsOfferings#product_type}.
	// Experimental.
	ProductType *string `field:"optional" json:"productType" yaml:"productType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#service_codes DataAwsOfferings#service_codes}.
	// Experimental.
	ServiceCodes *[]*string `field:"optional" json:"serviceCodes" yaml:"serviceCodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#usage_types DataAwsOfferings#usage_types}.
	// Experimental.
	UsageTypes *[]*string `field:"optional" json:"usageTypes" yaml:"usageTypes"`
}

