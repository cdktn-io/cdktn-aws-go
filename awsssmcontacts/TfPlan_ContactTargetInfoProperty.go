package awsssmcontacts


// Experimental.
type TfPlan_ContactTargetInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#is_essential TfPlan#is_essential}.
	// Experimental.
	IsEssential interface{} `field:"required" json:"isEssential" yaml:"isEssential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#contact_id TfPlan#contact_id}.
	// Experimental.
	ContactId *string `field:"optional" json:"contactId" yaml:"contactId"`
}

