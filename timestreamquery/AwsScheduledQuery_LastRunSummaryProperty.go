package timestreamquery


// Experimental.
type AwsScheduledQuery_LastRunSummaryProperty struct {
	// error_report_location block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#error_report_location AwsScheduledQuery#error_report_location}
	// Experimental.
	ErrorReportLocation interface{} `field:"optional" json:"errorReportLocation" yaml:"errorReportLocation"`
	// execution_stats block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#execution_stats AwsScheduledQuery#execution_stats}
	// Experimental.
	ExecutionStats interface{} `field:"optional" json:"executionStats" yaml:"executionStats"`
	// query_insights_response block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/timestreamquery_scheduled_query#query_insights_response AwsScheduledQuery#query_insights_response}
	// Experimental.
	QueryInsightsResponse interface{} `field:"optional" json:"queryInsightsResponse" yaml:"queryInsightsResponse"`
}

