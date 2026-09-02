package awsquicksight


// Experimental.
type TfDashboard_DashboardPublishOptionsProperty struct {
	// ad_hoc_filtering_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#ad_hoc_filtering_option TfDashboard#ad_hoc_filtering_option}
	// Experimental.
	AdHocFilteringOption *TfDashboard_AdHocFilteringOptionProperty `field:"optional" json:"adHocFilteringOption" yaml:"adHocFilteringOption"`
	// data_point_drill_up_down_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#data_point_drill_up_down_option TfDashboard#data_point_drill_up_down_option}
	// Experimental.
	DataPointDrillUpDownOption *TfDashboard_DataPointDrillUpDownOptionProperty `field:"optional" json:"dataPointDrillUpDownOption" yaml:"dataPointDrillUpDownOption"`
	// data_point_menu_label_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#data_point_menu_label_option TfDashboard#data_point_menu_label_option}
	// Experimental.
	DataPointMenuLabelOption *TfDashboard_DataPointMenuLabelOptionProperty `field:"optional" json:"dataPointMenuLabelOption" yaml:"dataPointMenuLabelOption"`
	// data_point_tooltip_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#data_point_tooltip_option TfDashboard#data_point_tooltip_option}
	// Experimental.
	DataPointTooltipOption *TfDashboard_DataPointTooltipOptionProperty `field:"optional" json:"dataPointTooltipOption" yaml:"dataPointTooltipOption"`
	// export_to_csv_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#export_to_csv_option TfDashboard#export_to_csv_option}
	// Experimental.
	ExportToCsvOption *TfDashboard_ExportToCsvOptionProperty `field:"optional" json:"exportToCsvOption" yaml:"exportToCsvOption"`
	// export_with_hidden_fields_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#export_with_hidden_fields_option TfDashboard#export_with_hidden_fields_option}
	// Experimental.
	ExportWithHiddenFieldsOption *TfDashboard_ExportWithHiddenFieldsOptionProperty `field:"optional" json:"exportWithHiddenFieldsOption" yaml:"exportWithHiddenFieldsOption"`
	// sheet_controls_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#sheet_controls_option TfDashboard#sheet_controls_option}
	// Experimental.
	SheetControlsOption *TfDashboard_SheetControlsOptionProperty `field:"optional" json:"sheetControlsOption" yaml:"sheetControlsOption"`
	// sheet_layout_element_maximization_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#sheet_layout_element_maximization_option TfDashboard#sheet_layout_element_maximization_option}
	// Experimental.
	SheetLayoutElementMaximizationOption *TfDashboard_SheetLayoutElementMaximizationOptionProperty `field:"optional" json:"sheetLayoutElementMaximizationOption" yaml:"sheetLayoutElementMaximizationOption"`
	// visual_axis_sort_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#visual_axis_sort_option TfDashboard#visual_axis_sort_option}
	// Experimental.
	VisualAxisSortOption *TfDashboard_VisualAxisSortOptionProperty `field:"optional" json:"visualAxisSortOption" yaml:"visualAxisSortOption"`
	// visual_menu_option block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/6.62.0/docs/resources/quicksight_dashboard#visual_menu_option TfDashboard#visual_menu_option}
	// Experimental.
	VisualMenuOption *TfDashboard_VisualMenuOptionProperty `field:"optional" json:"visualMenuOption" yaml:"visualMenuOption"`
}

