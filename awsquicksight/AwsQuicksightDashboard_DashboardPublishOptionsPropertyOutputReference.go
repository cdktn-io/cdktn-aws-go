package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdHocFilteringOption() AwsQuicksightDashboard_AdHocFilteringOptionPropertyOutputReference
	// Experimental.
	AdHocFilteringOptionInput() *AwsQuicksightDashboard_AdHocFilteringOptionProperty
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DataPointDrillUpDownOption() AwsQuicksightDashboard_DataPointDrillUpDownOptionPropertyOutputReference
	// Experimental.
	DataPointDrillUpDownOptionInput() *AwsQuicksightDashboard_DataPointDrillUpDownOptionProperty
	// Experimental.
	DataPointMenuLabelOption() AwsQuicksightDashboard_DataPointMenuLabelOptionPropertyOutputReference
	// Experimental.
	DataPointMenuLabelOptionInput() *AwsQuicksightDashboard_DataPointMenuLabelOptionProperty
	// Experimental.
	DataPointTooltipOption() AwsQuicksightDashboard_DataPointTooltipOptionPropertyOutputReference
	// Experimental.
	DataPointTooltipOptionInput() *AwsQuicksightDashboard_DataPointTooltipOptionProperty
	// Experimental.
	ExportToCsvOption() AwsQuicksightDashboard_ExportToCsvOptionPropertyOutputReference
	// Experimental.
	ExportToCsvOptionInput() *AwsQuicksightDashboard_ExportToCsvOptionProperty
	// Experimental.
	ExportWithHiddenFieldsOption() AwsQuicksightDashboard_ExportWithHiddenFieldsOptionPropertyOutputReference
	// Experimental.
	ExportWithHiddenFieldsOptionInput() *AwsQuicksightDashboard_ExportWithHiddenFieldsOptionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsQuicksightDashboard_DashboardPublishOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsQuicksightDashboard_DashboardPublishOptionsProperty)
	// Experimental.
	SheetControlsOption() AwsQuicksightDashboard_SheetControlsOptionPropertyOutputReference
	// Experimental.
	SheetControlsOptionInput() *AwsQuicksightDashboard_SheetControlsOptionProperty
	// Experimental.
	SheetLayoutElementMaximizationOption() AwsQuicksightDashboard_SheetLayoutElementMaximizationOptionPropertyOutputReference
	// Experimental.
	SheetLayoutElementMaximizationOptionInput() *AwsQuicksightDashboard_SheetLayoutElementMaximizationOptionProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VisualAxisSortOption() AwsQuicksightDashboard_VisualAxisSortOptionPropertyOutputReference
	// Experimental.
	VisualAxisSortOptionInput() *AwsQuicksightDashboard_VisualAxisSortOptionProperty
	// Experimental.
	VisualMenuOption() AwsQuicksightDashboard_VisualMenuOptionPropertyOutputReference
	// Experimental.
	VisualMenuOptionInput() *AwsQuicksightDashboard_VisualMenuOptionProperty
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
	PutAdHocFilteringOption(value *AwsQuicksightDashboard_AdHocFilteringOptionProperty)
	// Experimental.
	PutDataPointDrillUpDownOption(value *AwsQuicksightDashboard_DataPointDrillUpDownOptionProperty)
	// Experimental.
	PutDataPointMenuLabelOption(value *AwsQuicksightDashboard_DataPointMenuLabelOptionProperty)
	// Experimental.
	PutDataPointTooltipOption(value *AwsQuicksightDashboard_DataPointTooltipOptionProperty)
	// Experimental.
	PutExportToCsvOption(value *AwsQuicksightDashboard_ExportToCsvOptionProperty)
	// Experimental.
	PutExportWithHiddenFieldsOption(value *AwsQuicksightDashboard_ExportWithHiddenFieldsOptionProperty)
	// Experimental.
	PutSheetControlsOption(value *AwsQuicksightDashboard_SheetControlsOptionProperty)
	// Experimental.
	PutSheetLayoutElementMaximizationOption(value *AwsQuicksightDashboard_SheetLayoutElementMaximizationOptionProperty)
	// Experimental.
	PutVisualAxisSortOption(value *AwsQuicksightDashboard_VisualAxisSortOptionProperty)
	// Experimental.
	PutVisualMenuOption(value *AwsQuicksightDashboard_VisualMenuOptionProperty)
	// Experimental.
	ResetAdHocFilteringOption()
	// Experimental.
	ResetDataPointDrillUpDownOption()
	// Experimental.
	ResetDataPointMenuLabelOption()
	// Experimental.
	ResetDataPointTooltipOption()
	// Experimental.
	ResetExportToCsvOption()
	// Experimental.
	ResetExportWithHiddenFieldsOption()
	// Experimental.
	ResetSheetControlsOption()
	// Experimental.
	ResetSheetLayoutElementMaximizationOption()
	// Experimental.
	ResetVisualAxisSortOption()
	// Experimental.
	ResetVisualMenuOption()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference
type jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) AdHocFilteringOption() AwsQuicksightDashboard_AdHocFilteringOptionPropertyOutputReference {
	var returns AwsQuicksightDashboard_AdHocFilteringOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"adHocFilteringOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) AdHocFilteringOptionInput() *AwsQuicksightDashboard_AdHocFilteringOptionProperty {
	var returns *AwsQuicksightDashboard_AdHocFilteringOptionProperty
	_jsii_.Get(
		j,
		"adHocFilteringOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointDrillUpDownOption() AwsQuicksightDashboard_DataPointDrillUpDownOptionPropertyOutputReference {
	var returns AwsQuicksightDashboard_DataPointDrillUpDownOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"dataPointDrillUpDownOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointDrillUpDownOptionInput() *AwsQuicksightDashboard_DataPointDrillUpDownOptionProperty {
	var returns *AwsQuicksightDashboard_DataPointDrillUpDownOptionProperty
	_jsii_.Get(
		j,
		"dataPointDrillUpDownOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointMenuLabelOption() AwsQuicksightDashboard_DataPointMenuLabelOptionPropertyOutputReference {
	var returns AwsQuicksightDashboard_DataPointMenuLabelOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"dataPointMenuLabelOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointMenuLabelOptionInput() *AwsQuicksightDashboard_DataPointMenuLabelOptionProperty {
	var returns *AwsQuicksightDashboard_DataPointMenuLabelOptionProperty
	_jsii_.Get(
		j,
		"dataPointMenuLabelOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointTooltipOption() AwsQuicksightDashboard_DataPointTooltipOptionPropertyOutputReference {
	var returns AwsQuicksightDashboard_DataPointTooltipOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"dataPointTooltipOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointTooltipOptionInput() *AwsQuicksightDashboard_DataPointTooltipOptionProperty {
	var returns *AwsQuicksightDashboard_DataPointTooltipOptionProperty
	_jsii_.Get(
		j,
		"dataPointTooltipOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ExportToCsvOption() AwsQuicksightDashboard_ExportToCsvOptionPropertyOutputReference {
	var returns AwsQuicksightDashboard_ExportToCsvOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"exportToCsvOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ExportToCsvOptionInput() *AwsQuicksightDashboard_ExportToCsvOptionProperty {
	var returns *AwsQuicksightDashboard_ExportToCsvOptionProperty
	_jsii_.Get(
		j,
		"exportToCsvOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ExportWithHiddenFieldsOption() AwsQuicksightDashboard_ExportWithHiddenFieldsOptionPropertyOutputReference {
	var returns AwsQuicksightDashboard_ExportWithHiddenFieldsOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"exportWithHiddenFieldsOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ExportWithHiddenFieldsOptionInput() *AwsQuicksightDashboard_ExportWithHiddenFieldsOptionProperty {
	var returns *AwsQuicksightDashboard_ExportWithHiddenFieldsOptionProperty
	_jsii_.Get(
		j,
		"exportWithHiddenFieldsOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) InternalValue() *AwsQuicksightDashboard_DashboardPublishOptionsProperty {
	var returns *AwsQuicksightDashboard_DashboardPublishOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) SheetControlsOption() AwsQuicksightDashboard_SheetControlsOptionPropertyOutputReference {
	var returns AwsQuicksightDashboard_SheetControlsOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"sheetControlsOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) SheetControlsOptionInput() *AwsQuicksightDashboard_SheetControlsOptionProperty {
	var returns *AwsQuicksightDashboard_SheetControlsOptionProperty
	_jsii_.Get(
		j,
		"sheetControlsOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) SheetLayoutElementMaximizationOption() AwsQuicksightDashboard_SheetLayoutElementMaximizationOptionPropertyOutputReference {
	var returns AwsQuicksightDashboard_SheetLayoutElementMaximizationOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"sheetLayoutElementMaximizationOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) SheetLayoutElementMaximizationOptionInput() *AwsQuicksightDashboard_SheetLayoutElementMaximizationOptionProperty {
	var returns *AwsQuicksightDashboard_SheetLayoutElementMaximizationOptionProperty
	_jsii_.Get(
		j,
		"sheetLayoutElementMaximizationOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) VisualAxisSortOption() AwsQuicksightDashboard_VisualAxisSortOptionPropertyOutputReference {
	var returns AwsQuicksightDashboard_VisualAxisSortOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"visualAxisSortOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) VisualAxisSortOptionInput() *AwsQuicksightDashboard_VisualAxisSortOptionProperty {
	var returns *AwsQuicksightDashboard_VisualAxisSortOptionProperty
	_jsii_.Get(
		j,
		"visualAxisSortOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) VisualMenuOption() AwsQuicksightDashboard_VisualMenuOptionPropertyOutputReference {
	var returns AwsQuicksightDashboard_VisualMenuOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"visualMenuOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) VisualMenuOptionInput() *AwsQuicksightDashboard_VisualMenuOptionProperty {
	var returns *AwsQuicksightDashboard_VisualMenuOptionProperty
	_jsii_.Get(
		j,
		"visualMenuOptionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightDashboard.DashboardPublishOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference_Override(a AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightDashboard.DashboardPublishOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference)SetInternalValue(val *AwsQuicksightDashboard_DashboardPublishOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) PutAdHocFilteringOption(value *AwsQuicksightDashboard_AdHocFilteringOptionProperty) {
	if err := a.validatePutAdHocFilteringOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdHocFilteringOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) PutDataPointDrillUpDownOption(value *AwsQuicksightDashboard_DataPointDrillUpDownOptionProperty) {
	if err := a.validatePutDataPointDrillUpDownOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataPointDrillUpDownOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) PutDataPointMenuLabelOption(value *AwsQuicksightDashboard_DataPointMenuLabelOptionProperty) {
	if err := a.validatePutDataPointMenuLabelOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataPointMenuLabelOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) PutDataPointTooltipOption(value *AwsQuicksightDashboard_DataPointTooltipOptionProperty) {
	if err := a.validatePutDataPointTooltipOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataPointTooltipOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) PutExportToCsvOption(value *AwsQuicksightDashboard_ExportToCsvOptionProperty) {
	if err := a.validatePutExportToCsvOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExportToCsvOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) PutExportWithHiddenFieldsOption(value *AwsQuicksightDashboard_ExportWithHiddenFieldsOptionProperty) {
	if err := a.validatePutExportWithHiddenFieldsOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExportWithHiddenFieldsOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) PutSheetControlsOption(value *AwsQuicksightDashboard_SheetControlsOptionProperty) {
	if err := a.validatePutSheetControlsOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSheetControlsOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) PutSheetLayoutElementMaximizationOption(value *AwsQuicksightDashboard_SheetLayoutElementMaximizationOptionProperty) {
	if err := a.validatePutSheetLayoutElementMaximizationOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSheetLayoutElementMaximizationOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) PutVisualAxisSortOption(value *AwsQuicksightDashboard_VisualAxisSortOptionProperty) {
	if err := a.validatePutVisualAxisSortOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVisualAxisSortOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) PutVisualMenuOption(value *AwsQuicksightDashboard_VisualMenuOptionProperty) {
	if err := a.validatePutVisualMenuOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVisualMenuOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ResetAdHocFilteringOption() {
	_jsii_.InvokeVoid(
		a,
		"resetAdHocFilteringOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ResetDataPointDrillUpDownOption() {
	_jsii_.InvokeVoid(
		a,
		"resetDataPointDrillUpDownOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ResetDataPointMenuLabelOption() {
	_jsii_.InvokeVoid(
		a,
		"resetDataPointMenuLabelOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ResetDataPointTooltipOption() {
	_jsii_.InvokeVoid(
		a,
		"resetDataPointTooltipOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ResetExportToCsvOption() {
	_jsii_.InvokeVoid(
		a,
		"resetExportToCsvOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ResetExportWithHiddenFieldsOption() {
	_jsii_.InvokeVoid(
		a,
		"resetExportWithHiddenFieldsOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ResetSheetControlsOption() {
	_jsii_.InvokeVoid(
		a,
		"resetSheetControlsOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ResetSheetLayoutElementMaximizationOption() {
	_jsii_.InvokeVoid(
		a,
		"resetSheetLayoutElementMaximizationOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ResetVisualAxisSortOption() {
	_jsii_.InvokeVoid(
		a,
		"resetVisualAxisSortOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ResetVisualMenuOption() {
	_jsii_.InvokeVoid(
		a,
		"resetVisualMenuOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightDashboard_DashboardPublishOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

