package ssmcontacts


// Experimental.
type AwsPlan_ContactTargetInfoProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#is_essential AwsPlan#is_essential}.
	// Experimental.
	IsEssential interface{} `field:"required" json:"isEssential" yaml:"isEssential"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/ssmcontacts_plan#contact_id AwsPlan#contact_id}.
	// Experimental.
	ContactId *string `field:"optional" json:"contactId" yaml:"contactId"`
}

