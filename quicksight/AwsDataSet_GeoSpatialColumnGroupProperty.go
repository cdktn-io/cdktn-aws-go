package quicksight


// Experimental.
type AwsDataSet_GeoSpatialColumnGroupProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#columns AwsDataSet#columns}.
	// Experimental.
	Columns *[]*string `field:"required" json:"columns" yaml:"columns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#country_code AwsDataSet#country_code}.
	// Experimental.
	CountryCode *string `field:"required" json:"countryCode" yaml:"countryCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_data_set#name AwsDataSet#name}.
	// Experimental.
	Name *string `field:"required" json:"name" yaml:"name"`
}

