package awsinvoicing

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfInvoiceUnitConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/invoicing_invoice_unit#invoice_receiver TfInvoiceUnit#invoice_receiver}.
	// Experimental.
	InvoiceReceiver *string `field:"required" json:"invoiceReceiver" yaml:"invoiceReceiver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/invoicing_invoice_unit#name TfInvoiceUnit#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/invoicing_invoice_unit#description TfInvoiceUnit#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Region where this resource will be [managed](https://docs.aws.amazon.com/general/latest/gr/rande.html#regional-endpoints). Defaults to the Region set in the [provider configuration](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#aws-configuration-reference).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/invoicing_invoice_unit#region TfInvoiceUnit#region}
	// Experimental.
	Region *string `field:"optional" json:"region" yaml:"region"`
	// rule block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/invoicing_invoice_unit#rule TfInvoiceUnit#rule}
	// Experimental.
	Rule interface{} `field:"optional" json:"rule" yaml:"rule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/invoicing_invoice_unit#tags TfInvoiceUnit#tags}.
	// Experimental.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/invoicing_invoice_unit#tax_inheritance_disabled TfInvoiceUnit#tax_inheritance_disabled}.
	// Experimental.
	TaxInheritanceDisabled interface{} `field:"optional" json:"taxInheritanceDisabled" yaml:"taxInheritanceDisabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/invoicing_invoice_unit#timeouts TfInvoiceUnit#timeouts}
	// Experimental.
	Timeouts *TfInvoiceUnit_TimeoutsProperty `field:"optional" json:"timeouts" yaml:"timeouts"`
}

