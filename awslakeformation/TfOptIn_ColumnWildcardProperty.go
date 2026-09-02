package awslakeformation


// Experimental.
type TfOptIn_ColumnWildcardProperty struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/lakeformation_opt_in#excluded_column_names TfOptIn#excluded_column_names}.
	// Experimental.
	ExcludedColumnNames *[]*string `field:"optional" json:"excludedColumnNames" yaml:"excludedColumnNames"`
}

