package awsbcmdataexports


// Experimental.
type AwsBcmdataexportsExport_ExportProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#name AwsBcmdataexportsExport#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
	// data_query block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#data_query AwsBcmdataexportsExport#data_query}
	// Experimental.
	DataQuery interface{} `field:"optional" json:"dataQuery" yaml:"dataQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#description AwsBcmdataexportsExport#description}.
	// Experimental.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// destination_configurations block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#destination_configurations AwsBcmdataexportsExport#destination_configurations}
	// Experimental.
	DestinationConfigurations interface{} `field:"optional" json:"destinationConfigurations" yaml:"destinationConfigurations"`
	// refresh_cadence block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/bcmdataexports_export#refresh_cadence AwsBcmdataexportsExport#refresh_cadence}
	// Experimental.
	RefreshCadence interface{} `field:"optional" json:"refreshCadence" yaml:"refreshCadence"`
}

