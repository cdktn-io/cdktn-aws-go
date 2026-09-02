package awstimestreamquery


// Experimental.
type TfScheduledQuery_TimestreamConfigurationProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#database_name TfScheduledQuery#database_name}.
	// Experimental.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#table_name TfScheduledQuery#table_name}.
	// Experimental.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#time_column TfScheduledQuery#time_column}.
	// Experimental.
	TimeColumn *string `field:"required" json:"timeColumn" yaml:"timeColumn"`
	// dimension_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#dimension_mapping TfScheduledQuery#dimension_mapping}
	// Experimental.
	DimensionMapping interface{} `field:"optional" json:"dimensionMapping" yaml:"dimensionMapping"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#measure_name_column TfScheduledQuery#measure_name_column}.
	// Experimental.
	MeasureNameColumn *string `field:"optional" json:"measureNameColumn" yaml:"measureNameColumn"`
	// mixed_measure_mapping block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#mixed_measure_mapping TfScheduledQuery#mixed_measure_mapping}
	// Experimental.
	MixedMeasureMapping interface{} `field:"optional" json:"mixedMeasureMapping" yaml:"mixedMeasureMapping"`
	// multi_measure_mappings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#multi_measure_mappings TfScheduledQuery#multi_measure_mappings}
	// Experimental.
	MultiMeasureMappings interface{} `field:"optional" json:"multiMeasureMappings" yaml:"multiMeasureMappings"`
}

