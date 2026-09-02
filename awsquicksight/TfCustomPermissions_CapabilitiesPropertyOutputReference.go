package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfCustomPermissions_CapabilitiesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AddOrRunAnomalyDetectionForAnalyses() *string
	// Experimental.
	SetAddOrRunAnomalyDetectionForAnalyses(val *string)
	// Experimental.
	AddOrRunAnomalyDetectionForAnalysesInput() *string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// Experimental.
	CreateAndUpdateDashboardEmailReports() *string
	// Experimental.
	SetCreateAndUpdateDashboardEmailReports(val *string)
	// Experimental.
	CreateAndUpdateDashboardEmailReportsInput() *string
	// Experimental.
	CreateAndUpdateDatasets() *string
	// Experimental.
	SetCreateAndUpdateDatasets(val *string)
	// Experimental.
	CreateAndUpdateDatasetsInput() *string
	// Experimental.
	CreateAndUpdateDataSources() *string
	// Experimental.
	SetCreateAndUpdateDataSources(val *string)
	// Experimental.
	CreateAndUpdateDataSourcesInput() *string
	// Experimental.
	CreateAndUpdateThemes() *string
	// Experimental.
	SetCreateAndUpdateThemes(val *string)
	// Experimental.
	CreateAndUpdateThemesInput() *string
	// Experimental.
	CreateAndUpdateThresholdAlerts() *string
	// Experimental.
	SetCreateAndUpdateThresholdAlerts(val *string)
	// Experimental.
	CreateAndUpdateThresholdAlertsInput() *string
	// Experimental.
	CreateSharedFolders() *string
	// Experimental.
	SetCreateSharedFolders(val *string)
	// Experimental.
	CreateSharedFoldersInput() *string
	// Experimental.
	CreateSpiceDataset() *string
	// Experimental.
	SetCreateSpiceDataset(val *string)
	// Experimental.
	CreateSpiceDatasetInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	ExportToCsv() *string
	// Experimental.
	SetExportToCsv(val *string)
	// Experimental.
	ExportToCsvInput() *string
	// Experimental.
	ExportToCsvInScheduledReports() *string
	// Experimental.
	SetExportToCsvInScheduledReports(val *string)
	// Experimental.
	ExportToCsvInScheduledReportsInput() *string
	// Experimental.
	ExportToExcel() *string
	// Experimental.
	SetExportToExcel(val *string)
	// Experimental.
	ExportToExcelInput() *string
	// Experimental.
	ExportToExcelInScheduledReports() *string
	// Experimental.
	SetExportToExcelInScheduledReports(val *string)
	// Experimental.
	ExportToExcelInScheduledReportsInput() *string
	// Experimental.
	ExportToPdf() *string
	// Experimental.
	SetExportToPdf(val *string)
	// Experimental.
	ExportToPdfInput() *string
	// Experimental.
	ExportToPdfInScheduledReports() *string
	// Experimental.
	SetExportToPdfInScheduledReports(val *string)
	// Experimental.
	ExportToPdfInScheduledReportsInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	IncludeContentInScheduledReportsEmail() *string
	// Experimental.
	SetIncludeContentInScheduledReportsEmail(val *string)
	// Experimental.
	IncludeContentInScheduledReportsEmailInput() *string
	// Experimental.
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	PrintReports() *string
	// Experimental.
	SetPrintReports(val *string)
	// Experimental.
	PrintReportsInput() *string
	// Experimental.
	RenameSharedFolders() *string
	// Experimental.
	SetRenameSharedFolders(val *string)
	// Experimental.
	RenameSharedFoldersInput() *string
	// Experimental.
	ShareAnalyses() *string
	// Experimental.
	SetShareAnalyses(val *string)
	// Experimental.
	ShareAnalysesInput() *string
	// Experimental.
	ShareDashboards() *string
	// Experimental.
	SetShareDashboards(val *string)
	// Experimental.
	ShareDashboardsInput() *string
	// Experimental.
	ShareDatasets() *string
	// Experimental.
	SetShareDatasets(val *string)
	// Experimental.
	ShareDatasetsInput() *string
	// Experimental.
	ShareDataSources() *string
	// Experimental.
	SetShareDataSources(val *string)
	// Experimental.
	ShareDataSourcesInput() *string
	// Experimental.
	SubscribeDashboardEmailReports() *string
	// Experimental.
	SetSubscribeDashboardEmailReports(val *string)
	// Experimental.
	SubscribeDashboardEmailReportsInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ViewAccountSpiceCapacity() *string
	// Experimental.
	SetViewAccountSpiceCapacity(val *string)
	// Experimental.
	ViewAccountSpiceCapacityInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	// Experimental.
	ResetAddOrRunAnomalyDetectionForAnalyses()
	// Experimental.
	ResetCreateAndUpdateDashboardEmailReports()
	// Experimental.
	ResetCreateAndUpdateDatasets()
	// Experimental.
	ResetCreateAndUpdateDataSources()
	// Experimental.
	ResetCreateAndUpdateThemes()
	// Experimental.
	ResetCreateAndUpdateThresholdAlerts()
	// Experimental.
	ResetCreateSharedFolders()
	// Experimental.
	ResetCreateSpiceDataset()
	// Experimental.
	ResetExportToCsv()
	// Experimental.
	ResetExportToCsvInScheduledReports()
	// Experimental.
	ResetExportToExcel()
	// Experimental.
	ResetExportToExcelInScheduledReports()
	// Experimental.
	ResetExportToPdf()
	// Experimental.
	ResetExportToPdfInScheduledReports()
	// Experimental.
	ResetIncludeContentInScheduledReportsEmail()
	// Experimental.
	ResetPrintReports()
	// Experimental.
	ResetRenameSharedFolders()
	// Experimental.
	ResetShareAnalyses()
	// Experimental.
	ResetShareDashboards()
	// Experimental.
	ResetShareDatasets()
	// Experimental.
	ResetShareDataSources()
	// Experimental.
	ResetSubscribeDashboardEmailReports()
	// Experimental.
	ResetViewAccountSpiceCapacity()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfCustomPermissions_CapabilitiesPropertyOutputReference
type jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) AddOrRunAnomalyDetectionForAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addOrRunAnomalyDetectionForAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) AddOrRunAnomalyDetectionForAnalysesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addOrRunAnomalyDetectionForAnalysesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateAndUpdateDashboardEmailReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDashboardEmailReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateAndUpdateDashboardEmailReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDashboardEmailReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateAndUpdateDatasets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateAndUpdateDatasetsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDatasetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateAndUpdateDataSources() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateAndUpdateDataSourcesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateDataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateAndUpdateThemes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThemes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateAndUpdateThemesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThemesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateAndUpdateThresholdAlerts() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThresholdAlerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateAndUpdateThresholdAlertsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createAndUpdateThresholdAlertsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateSharedFolders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSharedFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateSharedFoldersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSharedFoldersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateSpiceDataset() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSpiceDataset",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreateSpiceDatasetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createSpiceDatasetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ExportToCsv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ExportToCsvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ExportToCsvInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsvInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ExportToCsvInScheduledReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToCsvInScheduledReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ExportToExcel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ExportToExcelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ExportToExcelInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcelInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ExportToExcelInScheduledReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToExcelInScheduledReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ExportToPdf() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdf",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ExportToPdfInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdfInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ExportToPdfInScheduledReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdfInScheduledReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ExportToPdfInScheduledReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"exportToPdfInScheduledReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) IncludeContentInScheduledReportsEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeContentInScheduledReportsEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) IncludeContentInScheduledReportsEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"includeContentInScheduledReportsEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) PrintReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) PrintReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"printReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) RenameSharedFolders() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renameSharedFolders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) RenameSharedFoldersInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"renameSharedFoldersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ShareAnalyses() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAnalyses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ShareAnalysesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareAnalysesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ShareDashboards() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDashboards",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ShareDashboardsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDashboardsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ShareDatasets() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDatasets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ShareDatasetsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDatasetsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ShareDataSources() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ShareDataSourcesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"shareDataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) SubscribeDashboardEmailReports() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscribeDashboardEmailReports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) SubscribeDashboardEmailReportsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscribeDashboardEmailReportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ViewAccountSpiceCapacity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewAccountSpiceCapacity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ViewAccountSpiceCapacityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"viewAccountSpiceCapacityInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfCustomPermissions_CapabilitiesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) TfCustomPermissions_CapabilitiesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfCustomPermissions_CapabilitiesPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfCustomPermissions.CapabilitiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewTfCustomPermissions_CapabilitiesPropertyOutputReference_Override(t TfCustomPermissions_CapabilitiesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfCustomPermissions.CapabilitiesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		t,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetAddOrRunAnomalyDetectionForAnalyses(val *string) {
	if err := j.validateSetAddOrRunAnomalyDetectionForAnalysesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"addOrRunAnomalyDetectionForAnalyses",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetCreateAndUpdateDashboardEmailReports(val *string) {
	if err := j.validateSetCreateAndUpdateDashboardEmailReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateDashboardEmailReports",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetCreateAndUpdateDatasets(val *string) {
	if err := j.validateSetCreateAndUpdateDatasetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateDatasets",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetCreateAndUpdateDataSources(val *string) {
	if err := j.validateSetCreateAndUpdateDataSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateDataSources",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetCreateAndUpdateThemes(val *string) {
	if err := j.validateSetCreateAndUpdateThemesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateThemes",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetCreateAndUpdateThresholdAlerts(val *string) {
	if err := j.validateSetCreateAndUpdateThresholdAlertsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createAndUpdateThresholdAlerts",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetCreateSharedFolders(val *string) {
	if err := j.validateSetCreateSharedFoldersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSharedFolders",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetCreateSpiceDataset(val *string) {
	if err := j.validateSetCreateSpiceDatasetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createSpiceDataset",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetExportToCsv(val *string) {
	if err := j.validateSetExportToCsvParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToCsv",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetExportToCsvInScheduledReports(val *string) {
	if err := j.validateSetExportToCsvInScheduledReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToCsvInScheduledReports",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetExportToExcel(val *string) {
	if err := j.validateSetExportToExcelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToExcel",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetExportToExcelInScheduledReports(val *string) {
	if err := j.validateSetExportToExcelInScheduledReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToExcelInScheduledReports",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetExportToPdf(val *string) {
	if err := j.validateSetExportToPdfParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToPdf",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetExportToPdfInScheduledReports(val *string) {
	if err := j.validateSetExportToPdfInScheduledReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exportToPdfInScheduledReports",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetIncludeContentInScheduledReportsEmail(val *string) {
	if err := j.validateSetIncludeContentInScheduledReportsEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeContentInScheduledReportsEmail",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetPrintReports(val *string) {
	if err := j.validateSetPrintReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"printReports",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetRenameSharedFolders(val *string) {
	if err := j.validateSetRenameSharedFoldersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renameSharedFolders",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetShareAnalyses(val *string) {
	if err := j.validateSetShareAnalysesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareAnalyses",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetShareDashboards(val *string) {
	if err := j.validateSetShareDashboardsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDashboards",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetShareDatasets(val *string) {
	if err := j.validateSetShareDatasetsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDatasets",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetShareDataSources(val *string) {
	if err := j.validateSetShareDataSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"shareDataSources",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetSubscribeDashboardEmailReports(val *string) {
	if err := j.validateSetSubscribeDashboardEmailReportsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subscribeDashboardEmailReports",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference)SetViewAccountSpiceCapacity(val *string) {
	if err := j.validateSetViewAccountSpiceCapacityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"viewAccountSpiceCapacity",
		val,
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetAddOrRunAnomalyDetectionForAnalyses() {
	_jsii_.InvokeVoid(
		t,
		"resetAddOrRunAnomalyDetectionForAnalyses",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetCreateAndUpdateDashboardEmailReports() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateAndUpdateDashboardEmailReports",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetCreateAndUpdateDatasets() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateAndUpdateDatasets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetCreateAndUpdateDataSources() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateAndUpdateDataSources",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetCreateAndUpdateThemes() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateAndUpdateThemes",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetCreateAndUpdateThresholdAlerts() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateAndUpdateThresholdAlerts",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetCreateSharedFolders() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateSharedFolders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetCreateSpiceDataset() {
	_jsii_.InvokeVoid(
		t,
		"resetCreateSpiceDataset",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetExportToCsv() {
	_jsii_.InvokeVoid(
		t,
		"resetExportToCsv",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetExportToCsvInScheduledReports() {
	_jsii_.InvokeVoid(
		t,
		"resetExportToCsvInScheduledReports",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetExportToExcel() {
	_jsii_.InvokeVoid(
		t,
		"resetExportToExcel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetExportToExcelInScheduledReports() {
	_jsii_.InvokeVoid(
		t,
		"resetExportToExcelInScheduledReports",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetExportToPdf() {
	_jsii_.InvokeVoid(
		t,
		"resetExportToPdf",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetExportToPdfInScheduledReports() {
	_jsii_.InvokeVoid(
		t,
		"resetExportToPdfInScheduledReports",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetIncludeContentInScheduledReportsEmail() {
	_jsii_.InvokeVoid(
		t,
		"resetIncludeContentInScheduledReportsEmail",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetPrintReports() {
	_jsii_.InvokeVoid(
		t,
		"resetPrintReports",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetRenameSharedFolders() {
	_jsii_.InvokeVoid(
		t,
		"resetRenameSharedFolders",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetShareAnalyses() {
	_jsii_.InvokeVoid(
		t,
		"resetShareAnalyses",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetShareDashboards() {
	_jsii_.InvokeVoid(
		t,
		"resetShareDashboards",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetShareDatasets() {
	_jsii_.InvokeVoid(
		t,
		"resetShareDatasets",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetShareDataSources() {
	_jsii_.InvokeVoid(
		t,
		"resetShareDataSources",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetSubscribeDashboardEmailReports() {
	_jsii_.InvokeVoid(
		t,
		"resetSubscribeDashboardEmailReports",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ResetViewAccountSpiceCapacity() {
	_jsii_.InvokeVoid(
		t,
		"resetViewAccountSpiceCapacity",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := t.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfCustomPermissions_CapabilitiesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

