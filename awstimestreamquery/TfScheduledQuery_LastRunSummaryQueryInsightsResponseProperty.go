package awstimestreamquery


// Experimental.
type TfScheduledQuery_LastRunSummaryQueryInsightsResponseProperty struct {
	// query_spatial_coverage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#query_spatial_coverage TfScheduledQuery#query_spatial_coverage}
	// Experimental.
	QuerySpatialCoverage interface{} `field:"optional" json:"querySpatialCoverage" yaml:"querySpatialCoverage"`
	// query_temporal_range block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#query_temporal_range TfScheduledQuery#query_temporal_range}
	// Experimental.
	QueryTemporalRange interface{} `field:"optional" json:"queryTemporalRange" yaml:"queryTemporalRange"`
}

