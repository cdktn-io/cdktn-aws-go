package awssavingsplans

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfOfferingsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#currencies DataTfOfferings#currencies}.
	// Experimental.
	Currencies *[]*string `field:"optional" json:"currencies" yaml:"currencies"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#descriptions DataTfOfferings#descriptions}.
	// Experimental.
	Descriptions *[]*string `field:"optional" json:"descriptions" yaml:"descriptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#durations DataTfOfferings#durations}.
	// Experimental.
	Durations *[]*float64 `field:"optional" json:"durations" yaml:"durations"`
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#filter DataTfOfferings#filter}
	// Experimental.
	Filter interface{} `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#offering_ids DataTfOfferings#offering_ids}.
	// Experimental.
	OfferingIds *[]*string `field:"optional" json:"offeringIds" yaml:"offeringIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#operations DataTfOfferings#operations}.
	// Experimental.
	Operations *[]*string `field:"optional" json:"operations" yaml:"operations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#payment_options DataTfOfferings#payment_options}.
	// Experimental.
	PaymentOptions *[]*string `field:"optional" json:"paymentOptions" yaml:"paymentOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#plan_types DataTfOfferings#plan_types}.
	// Experimental.
	PlanTypes *[]*string `field:"optional" json:"planTypes" yaml:"planTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#product_type DataTfOfferings#product_type}.
	// Experimental.
	ProductType *string `field:"optional" json:"productType" yaml:"productType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#service_codes DataTfOfferings#service_codes}.
	// Experimental.
	ServiceCodes *[]*string `field:"optional" json:"serviceCodes" yaml:"serviceCodes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/data-sources/savingsplans_offerings#usage_types DataTfOfferings#usage_types}.
	// Experimental.
	UsageTypes *[]*string `field:"optional" json:"usageTypes" yaml:"usageTypes"`
}

