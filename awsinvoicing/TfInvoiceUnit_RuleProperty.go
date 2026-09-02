package awsinvoicing


// Experimental.
type TfInvoiceUnit_RuleProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/invoicing_invoice_unit#linked_accounts TfInvoiceUnit#linked_accounts}.
	// Experimental.
	LinkedAccounts *[]*string `field:"required" json:"linkedAccounts" yaml:"linkedAccounts"`
}

