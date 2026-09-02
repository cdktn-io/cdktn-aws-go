package awsbcmdataexports


// Experimental.
type TfExport_S3OutputConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#compression TfExport#compression}.
	// Experimental.
	Compression *string `field:"required" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#format TfExport#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#output_type TfExport#output_type}.
	// Experimental.
	OutputType *string `field:"required" json:"outputType" yaml:"outputType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#overwrite TfExport#overwrite}.
	// Experimental.
	Overwrite *string `field:"required" json:"overwrite" yaml:"overwrite"`
}

