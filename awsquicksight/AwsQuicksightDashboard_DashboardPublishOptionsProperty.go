package awsquicksight


// Experimental.
type AwsQuicksightDashboard_DashboardPublishOptionsProperty struct {
	// ad_hoc_filtering_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#ad_hoc_filtering_option AwsQuicksightDashboard#ad_hoc_filtering_option}
	// Experimental.
	AdHocFilteringOption *AwsQuicksightDashboard_AdHocFilteringOptionProperty `field:"optional" json:"adHocFilteringOption" yaml:"adHocFilteringOption"`
	// data_point_drill_up_down_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#data_point_drill_up_down_option AwsQuicksightDashboard#data_point_drill_up_down_option}
	// Experimental.
	DataPointDrillUpDownOption *AwsQuicksightDashboard_DataPointDrillUpDownOptionProperty `field:"optional" json:"dataPointDrillUpDownOption" yaml:"dataPointDrillUpDownOption"`
	// data_point_menu_label_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#data_point_menu_label_option AwsQuicksightDashboard#data_point_menu_label_option}
	// Experimental.
	DataPointMenuLabelOption *AwsQuicksightDashboard_DataPointMenuLabelOptionProperty `field:"optional" json:"dataPointMenuLabelOption" yaml:"dataPointMenuLabelOption"`
	// data_point_tooltip_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#data_point_tooltip_option AwsQuicksightDashboard#data_point_tooltip_option}
	// Experimental.
	DataPointTooltipOption *AwsQuicksightDashboard_DataPointTooltipOptionProperty `field:"optional" json:"dataPointTooltipOption" yaml:"dataPointTooltipOption"`
	// export_to_csv_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#export_to_csv_option AwsQuicksightDashboard#export_to_csv_option}
	// Experimental.
	ExportToCsvOption *AwsQuicksightDashboard_ExportToCsvOptionProperty `field:"optional" json:"exportToCsvOption" yaml:"exportToCsvOption"`
	// export_with_hidden_fields_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#export_with_hidden_fields_option AwsQuicksightDashboard#export_with_hidden_fields_option}
	// Experimental.
	ExportWithHiddenFieldsOption *AwsQuicksightDashboard_ExportWithHiddenFieldsOptionProperty `field:"optional" json:"exportWithHiddenFieldsOption" yaml:"exportWithHiddenFieldsOption"`
	// sheet_controls_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#sheet_controls_option AwsQuicksightDashboard#sheet_controls_option}
	// Experimental.
	SheetControlsOption *AwsQuicksightDashboard_SheetControlsOptionProperty `field:"optional" json:"sheetControlsOption" yaml:"sheetControlsOption"`
	// sheet_layout_element_maximization_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#sheet_layout_element_maximization_option AwsQuicksightDashboard#sheet_layout_element_maximization_option}
	// Experimental.
	SheetLayoutElementMaximizationOption *AwsQuicksightDashboard_SheetLayoutElementMaximizationOptionProperty `field:"optional" json:"sheetLayoutElementMaximizationOption" yaml:"sheetLayoutElementMaximizationOption"`
	// visual_axis_sort_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#visual_axis_sort_option AwsQuicksightDashboard#visual_axis_sort_option}
	// Experimental.
	VisualAxisSortOption *AwsQuicksightDashboard_VisualAxisSortOptionProperty `field:"optional" json:"visualAxisSortOption" yaml:"visualAxisSortOption"`
	// visual_menu_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#visual_menu_option AwsQuicksightDashboard#visual_menu_option}
	// Experimental.
	VisualMenuOption *AwsQuicksightDashboard_VisualMenuOptionProperty `field:"optional" json:"visualMenuOption" yaml:"visualMenuOption"`
}

