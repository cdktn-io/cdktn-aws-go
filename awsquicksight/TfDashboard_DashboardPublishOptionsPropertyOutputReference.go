package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDashboard_DashboardPublishOptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdHocFilteringOption() TfDashboard_AdHocFilteringOptionPropertyOutputReference
	// Experimental.
	AdHocFilteringOptionInput() *TfDashboard_AdHocFilteringOptionProperty
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
	DataPointDrillUpDownOption() TfDashboard_DataPointDrillUpDownOptionPropertyOutputReference
	// Experimental.
	DataPointDrillUpDownOptionInput() *TfDashboard_DataPointDrillUpDownOptionProperty
	// Experimental.
	DataPointMenuLabelOption() TfDashboard_DataPointMenuLabelOptionPropertyOutputReference
	// Experimental.
	DataPointMenuLabelOptionInput() *TfDashboard_DataPointMenuLabelOptionProperty
	// Experimental.
	DataPointTooltipOption() TfDashboard_DataPointTooltipOptionPropertyOutputReference
	// Experimental.
	DataPointTooltipOptionInput() *TfDashboard_DataPointTooltipOptionProperty
	// Experimental.
	ExportToCsvOption() TfDashboard_ExportToCsvOptionPropertyOutputReference
	// Experimental.
	ExportToCsvOptionInput() *TfDashboard_ExportToCsvOptionProperty
	// Experimental.
	ExportWithHiddenFieldsOption() TfDashboard_ExportWithHiddenFieldsOptionPropertyOutputReference
	// Experimental.
	ExportWithHiddenFieldsOptionInput() *TfDashboard_ExportWithHiddenFieldsOptionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDashboard_DashboardPublishOptionsProperty
	// Experimental.
	SetInternalValue(val *TfDashboard_DashboardPublishOptionsProperty)
	// Experimental.
	SheetControlsOption() TfDashboard_SheetControlsOptionPropertyOutputReference
	// Experimental.
	SheetControlsOptionInput() *TfDashboard_SheetControlsOptionProperty
	// Experimental.
	SheetLayoutElementMaximizationOption() TfDashboard_SheetLayoutElementMaximizationOptionPropertyOutputReference
	// Experimental.
	SheetLayoutElementMaximizationOptionInput() *TfDashboard_SheetLayoutElementMaximizationOptionProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VisualAxisSortOption() TfDashboard_VisualAxisSortOptionPropertyOutputReference
	// Experimental.
	VisualAxisSortOptionInput() *TfDashboard_VisualAxisSortOptionProperty
	// Experimental.
	VisualMenuOption() TfDashboard_VisualMenuOptionPropertyOutputReference
	// Experimental.
	VisualMenuOptionInput() *TfDashboard_VisualMenuOptionProperty
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
	PutAdHocFilteringOption(value *TfDashboard_AdHocFilteringOptionProperty)
	// Experimental.
	PutDataPointDrillUpDownOption(value *TfDashboard_DataPointDrillUpDownOptionProperty)
	// Experimental.
	PutDataPointMenuLabelOption(value *TfDashboard_DataPointMenuLabelOptionProperty)
	// Experimental.
	PutDataPointTooltipOption(value *TfDashboard_DataPointTooltipOptionProperty)
	// Experimental.
	PutExportToCsvOption(value *TfDashboard_ExportToCsvOptionProperty)
	// Experimental.
	PutExportWithHiddenFieldsOption(value *TfDashboard_ExportWithHiddenFieldsOptionProperty)
	// Experimental.
	PutSheetControlsOption(value *TfDashboard_SheetControlsOptionProperty)
	// Experimental.
	PutSheetLayoutElementMaximizationOption(value *TfDashboard_SheetLayoutElementMaximizationOptionProperty)
	// Experimental.
	PutVisualAxisSortOption(value *TfDashboard_VisualAxisSortOptionProperty)
	// Experimental.
	PutVisualMenuOption(value *TfDashboard_VisualMenuOptionProperty)
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

// The jsii proxy struct for TfDashboard_DashboardPublishOptionsPropertyOutputReference
type jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) AdHocFilteringOption() TfDashboard_AdHocFilteringOptionPropertyOutputReference {
	var returns TfDashboard_AdHocFilteringOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"adHocFilteringOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) AdHocFilteringOptionInput() *TfDashboard_AdHocFilteringOptionProperty {
	var returns *TfDashboard_AdHocFilteringOptionProperty
	_jsii_.Get(
		j,
		"adHocFilteringOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointDrillUpDownOption() TfDashboard_DataPointDrillUpDownOptionPropertyOutputReference {
	var returns TfDashboard_DataPointDrillUpDownOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"dataPointDrillUpDownOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointDrillUpDownOptionInput() *TfDashboard_DataPointDrillUpDownOptionProperty {
	var returns *TfDashboard_DataPointDrillUpDownOptionProperty
	_jsii_.Get(
		j,
		"dataPointDrillUpDownOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointMenuLabelOption() TfDashboard_DataPointMenuLabelOptionPropertyOutputReference {
	var returns TfDashboard_DataPointMenuLabelOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"dataPointMenuLabelOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointMenuLabelOptionInput() *TfDashboard_DataPointMenuLabelOptionProperty {
	var returns *TfDashboard_DataPointMenuLabelOptionProperty
	_jsii_.Get(
		j,
		"dataPointMenuLabelOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointTooltipOption() TfDashboard_DataPointTooltipOptionPropertyOutputReference {
	var returns TfDashboard_DataPointTooltipOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"dataPointTooltipOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointTooltipOptionInput() *TfDashboard_DataPointTooltipOptionProperty {
	var returns *TfDashboard_DataPointTooltipOptionProperty
	_jsii_.Get(
		j,
		"dataPointTooltipOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ExportToCsvOption() TfDashboard_ExportToCsvOptionPropertyOutputReference {
	var returns TfDashboard_ExportToCsvOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"exportToCsvOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ExportToCsvOptionInput() *TfDashboard_ExportToCsvOptionProperty {
	var returns *TfDashboard_ExportToCsvOptionProperty
	_jsii_.Get(
		j,
		"exportToCsvOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ExportWithHiddenFieldsOption() TfDashboard_ExportWithHiddenFieldsOptionPropertyOutputReference {
	var returns TfDashboard_ExportWithHiddenFieldsOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"exportWithHiddenFieldsOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ExportWithHiddenFieldsOptionInput() *TfDashboard_ExportWithHiddenFieldsOptionProperty {
	var returns *TfDashboard_ExportWithHiddenFieldsOptionProperty
	_jsii_.Get(
		j,
		"exportWithHiddenFieldsOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) InternalValue() *TfDashboard_DashboardPublishOptionsProperty {
	var returns *TfDashboard_DashboardPublishOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) SheetControlsOption() TfDashboard_SheetControlsOptionPropertyOutputReference {
	var returns TfDashboard_SheetControlsOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"sheetControlsOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) SheetControlsOptionInput() *TfDashboard_SheetControlsOptionProperty {
	var returns *TfDashboard_SheetControlsOptionProperty
	_jsii_.Get(
		j,
		"sheetControlsOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) SheetLayoutElementMaximizationOption() TfDashboard_SheetLayoutElementMaximizationOptionPropertyOutputReference {
	var returns TfDashboard_SheetLayoutElementMaximizationOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"sheetLayoutElementMaximizationOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) SheetLayoutElementMaximizationOptionInput() *TfDashboard_SheetLayoutElementMaximizationOptionProperty {
	var returns *TfDashboard_SheetLayoutElementMaximizationOptionProperty
	_jsii_.Get(
		j,
		"sheetLayoutElementMaximizationOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) VisualAxisSortOption() TfDashboard_VisualAxisSortOptionPropertyOutputReference {
	var returns TfDashboard_VisualAxisSortOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"visualAxisSortOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) VisualAxisSortOptionInput() *TfDashboard_VisualAxisSortOptionProperty {
	var returns *TfDashboard_VisualAxisSortOptionProperty
	_jsii_.Get(
		j,
		"visualAxisSortOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) VisualMenuOption() TfDashboard_VisualMenuOptionPropertyOutputReference {
	var returns TfDashboard_VisualMenuOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"visualMenuOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) VisualMenuOptionInput() *TfDashboard_VisualMenuOptionProperty {
	var returns *TfDashboard_VisualMenuOptionProperty
	_jsii_.Get(
		j,
		"visualMenuOptionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDashboard_DashboardPublishOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDashboard_DashboardPublishOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDashboard_DashboardPublishOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfDashboard.DashboardPublishOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDashboard_DashboardPublishOptionsPropertyOutputReference_Override(t TfDashboard_DashboardPublishOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfDashboard.DashboardPublishOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference)SetInternalValue(val *TfDashboard_DashboardPublishOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) PutAdHocFilteringOption(value *TfDashboard_AdHocFilteringOptionProperty) {
	if err := t.validatePutAdHocFilteringOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAdHocFilteringOption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) PutDataPointDrillUpDownOption(value *TfDashboard_DataPointDrillUpDownOptionProperty) {
	if err := t.validatePutDataPointDrillUpDownOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataPointDrillUpDownOption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) PutDataPointMenuLabelOption(value *TfDashboard_DataPointMenuLabelOptionProperty) {
	if err := t.validatePutDataPointMenuLabelOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataPointMenuLabelOption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) PutDataPointTooltipOption(value *TfDashboard_DataPointTooltipOptionProperty) {
	if err := t.validatePutDataPointTooltipOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataPointTooltipOption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) PutExportToCsvOption(value *TfDashboard_ExportToCsvOptionProperty) {
	if err := t.validatePutExportToCsvOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExportToCsvOption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) PutExportWithHiddenFieldsOption(value *TfDashboard_ExportWithHiddenFieldsOptionProperty) {
	if err := t.validatePutExportWithHiddenFieldsOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putExportWithHiddenFieldsOption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) PutSheetControlsOption(value *TfDashboard_SheetControlsOptionProperty) {
	if err := t.validatePutSheetControlsOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSheetControlsOption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) PutSheetLayoutElementMaximizationOption(value *TfDashboard_SheetLayoutElementMaximizationOptionProperty) {
	if err := t.validatePutSheetLayoutElementMaximizationOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSheetLayoutElementMaximizationOption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) PutVisualAxisSortOption(value *TfDashboard_VisualAxisSortOptionProperty) {
	if err := t.validatePutVisualAxisSortOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVisualAxisSortOption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) PutVisualMenuOption(value *TfDashboard_VisualMenuOptionProperty) {
	if err := t.validatePutVisualMenuOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putVisualMenuOption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ResetAdHocFilteringOption() {
	_jsii_.InvokeVoid(
		t,
		"resetAdHocFilteringOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ResetDataPointDrillUpDownOption() {
	_jsii_.InvokeVoid(
		t,
		"resetDataPointDrillUpDownOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ResetDataPointMenuLabelOption() {
	_jsii_.InvokeVoid(
		t,
		"resetDataPointMenuLabelOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ResetDataPointTooltipOption() {
	_jsii_.InvokeVoid(
		t,
		"resetDataPointTooltipOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ResetExportToCsvOption() {
	_jsii_.InvokeVoid(
		t,
		"resetExportToCsvOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ResetExportWithHiddenFieldsOption() {
	_jsii_.InvokeVoid(
		t,
		"resetExportWithHiddenFieldsOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ResetSheetControlsOption() {
	_jsii_.InvokeVoid(
		t,
		"resetSheetControlsOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ResetSheetLayoutElementMaximizationOption() {
	_jsii_.InvokeVoid(
		t,
		"resetSheetLayoutElementMaximizationOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ResetVisualAxisSortOption() {
	_jsii_.InvokeVoid(
		t,
		"resetVisualAxisSortOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ResetVisualMenuOption() {
	_jsii_.InvokeVoid(
		t,
		"resetVisualMenuOption",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDashboard_DashboardPublishOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

