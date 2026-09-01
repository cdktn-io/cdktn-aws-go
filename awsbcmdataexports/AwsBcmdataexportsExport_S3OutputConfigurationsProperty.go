package awsbcmdataexports


// Experimental.
type AwsBcmdataexportsExport_S3OutputConfigurationsProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#compression AwsBcmdataexportsExport#compression}.
	// Experimental.
	Compression *string `field:"required" json:"compression" yaml:"compression"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#format AwsBcmdataexportsExport#format}.
	// Experimental.
	Format *string `field:"required" json:"format" yaml:"format"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#output_type AwsBcmdataexportsExport#output_type}.
	// Experimental.
	OutputType *string `field:"required" json:"outputType" yaml:"outputType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#overwrite AwsBcmdataexportsExport#overwrite}.
	// Experimental.
	Overwrite *string `field:"required" json:"overwrite" yaml:"overwrite"`
}

