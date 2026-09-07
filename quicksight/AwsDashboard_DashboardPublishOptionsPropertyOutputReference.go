package quicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/quicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/quicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsDashboard_DashboardPublishOptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdHocFilteringOption() AwsDashboard_AdHocFilteringOptionPropertyOutputReference
	// Experimental.
	AdHocFilteringOptionInput() *AwsDashboard_AdHocFilteringOptionProperty
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
	DataPointDrillUpDownOption() AwsDashboard_DataPointDrillUpDownOptionPropertyOutputReference
	// Experimental.
	DataPointDrillUpDownOptionInput() *AwsDashboard_DataPointDrillUpDownOptionProperty
	// Experimental.
	DataPointMenuLabelOption() AwsDashboard_DataPointMenuLabelOptionPropertyOutputReference
	// Experimental.
	DataPointMenuLabelOptionInput() *AwsDashboard_DataPointMenuLabelOptionProperty
	// Experimental.
	DataPointTooltipOption() AwsDashboard_DataPointTooltipOptionPropertyOutputReference
	// Experimental.
	DataPointTooltipOptionInput() *AwsDashboard_DataPointTooltipOptionProperty
	// Experimental.
	ExportToCsvOption() AwsDashboard_ExportToCsvOptionPropertyOutputReference
	// Experimental.
	ExportToCsvOptionInput() *AwsDashboard_ExportToCsvOptionProperty
	// Experimental.
	ExportWithHiddenFieldsOption() AwsDashboard_ExportWithHiddenFieldsOptionPropertyOutputReference
	// Experimental.
	ExportWithHiddenFieldsOptionInput() *AwsDashboard_ExportWithHiddenFieldsOptionProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsDashboard_DashboardPublishOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsDashboard_DashboardPublishOptionsProperty)
	// Experimental.
	SheetControlsOption() AwsDashboard_SheetControlsOptionPropertyOutputReference
	// Experimental.
	SheetControlsOptionInput() *AwsDashboard_SheetControlsOptionProperty
	// Experimental.
	SheetLayoutElementMaximizationOption() AwsDashboard_SheetLayoutElementMaximizationOptionPropertyOutputReference
	// Experimental.
	SheetLayoutElementMaximizationOptionInput() *AwsDashboard_SheetLayoutElementMaximizationOptionProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VisualAxisSortOption() AwsDashboard_VisualAxisSortOptionPropertyOutputReference
	// Experimental.
	VisualAxisSortOptionInput() *AwsDashboard_VisualAxisSortOptionProperty
	// Experimental.
	VisualMenuOption() AwsDashboard_VisualMenuOptionPropertyOutputReference
	// Experimental.
	VisualMenuOptionInput() *AwsDashboard_VisualMenuOptionProperty
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
	PutAdHocFilteringOption(value *AwsDashboard_AdHocFilteringOptionProperty)
	// Experimental.
	PutDataPointDrillUpDownOption(value *AwsDashboard_DataPointDrillUpDownOptionProperty)
	// Experimental.
	PutDataPointMenuLabelOption(value *AwsDashboard_DataPointMenuLabelOptionProperty)
	// Experimental.
	PutDataPointTooltipOption(value *AwsDashboard_DataPointTooltipOptionProperty)
	// Experimental.
	PutExportToCsvOption(value *AwsDashboard_ExportToCsvOptionProperty)
	// Experimental.
	PutExportWithHiddenFieldsOption(value *AwsDashboard_ExportWithHiddenFieldsOptionProperty)
	// Experimental.
	PutSheetControlsOption(value *AwsDashboard_SheetControlsOptionProperty)
	// Experimental.
	PutSheetLayoutElementMaximizationOption(value *AwsDashboard_SheetLayoutElementMaximizationOptionProperty)
	// Experimental.
	PutVisualAxisSortOption(value *AwsDashboard_VisualAxisSortOptionProperty)
	// Experimental.
	PutVisualMenuOption(value *AwsDashboard_VisualMenuOptionProperty)
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

// The jsii proxy struct for AwsDashboard_DashboardPublishOptionsPropertyOutputReference
type jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) AdHocFilteringOption() AwsDashboard_AdHocFilteringOptionPropertyOutputReference {
	var returns AwsDashboard_AdHocFilteringOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"adHocFilteringOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) AdHocFilteringOptionInput() *AwsDashboard_AdHocFilteringOptionProperty {
	var returns *AwsDashboard_AdHocFilteringOptionProperty
	_jsii_.Get(
		j,
		"adHocFilteringOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointDrillUpDownOption() AwsDashboard_DataPointDrillUpDownOptionPropertyOutputReference {
	var returns AwsDashboard_DataPointDrillUpDownOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"dataPointDrillUpDownOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointDrillUpDownOptionInput() *AwsDashboard_DataPointDrillUpDownOptionProperty {
	var returns *AwsDashboard_DataPointDrillUpDownOptionProperty
	_jsii_.Get(
		j,
		"dataPointDrillUpDownOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointMenuLabelOption() AwsDashboard_DataPointMenuLabelOptionPropertyOutputReference {
	var returns AwsDashboard_DataPointMenuLabelOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"dataPointMenuLabelOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointMenuLabelOptionInput() *AwsDashboard_DataPointMenuLabelOptionProperty {
	var returns *AwsDashboard_DataPointMenuLabelOptionProperty
	_jsii_.Get(
		j,
		"dataPointMenuLabelOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointTooltipOption() AwsDashboard_DataPointTooltipOptionPropertyOutputReference {
	var returns AwsDashboard_DataPointTooltipOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"dataPointTooltipOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) DataPointTooltipOptionInput() *AwsDashboard_DataPointTooltipOptionProperty {
	var returns *AwsDashboard_DataPointTooltipOptionProperty
	_jsii_.Get(
		j,
		"dataPointTooltipOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ExportToCsvOption() AwsDashboard_ExportToCsvOptionPropertyOutputReference {
	var returns AwsDashboard_ExportToCsvOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"exportToCsvOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ExportToCsvOptionInput() *AwsDashboard_ExportToCsvOptionProperty {
	var returns *AwsDashboard_ExportToCsvOptionProperty
	_jsii_.Get(
		j,
		"exportToCsvOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ExportWithHiddenFieldsOption() AwsDashboard_ExportWithHiddenFieldsOptionPropertyOutputReference {
	var returns AwsDashboard_ExportWithHiddenFieldsOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"exportWithHiddenFieldsOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ExportWithHiddenFieldsOptionInput() *AwsDashboard_ExportWithHiddenFieldsOptionProperty {
	var returns *AwsDashboard_ExportWithHiddenFieldsOptionProperty
	_jsii_.Get(
		j,
		"exportWithHiddenFieldsOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) InternalValue() *AwsDashboard_DashboardPublishOptionsProperty {
	var returns *AwsDashboard_DashboardPublishOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) SheetControlsOption() AwsDashboard_SheetControlsOptionPropertyOutputReference {
	var returns AwsDashboard_SheetControlsOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"sheetControlsOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) SheetControlsOptionInput() *AwsDashboard_SheetControlsOptionProperty {
	var returns *AwsDashboard_SheetControlsOptionProperty
	_jsii_.Get(
		j,
		"sheetControlsOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) SheetLayoutElementMaximizationOption() AwsDashboard_SheetLayoutElementMaximizationOptionPropertyOutputReference {
	var returns AwsDashboard_SheetLayoutElementMaximizationOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"sheetLayoutElementMaximizationOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) SheetLayoutElementMaximizationOptionInput() *AwsDashboard_SheetLayoutElementMaximizationOptionProperty {
	var returns *AwsDashboard_SheetLayoutElementMaximizationOptionProperty
	_jsii_.Get(
		j,
		"sheetLayoutElementMaximizationOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) VisualAxisSortOption() AwsDashboard_VisualAxisSortOptionPropertyOutputReference {
	var returns AwsDashboard_VisualAxisSortOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"visualAxisSortOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) VisualAxisSortOptionInput() *AwsDashboard_VisualAxisSortOptionProperty {
	var returns *AwsDashboard_VisualAxisSortOptionProperty
	_jsii_.Get(
		j,
		"visualAxisSortOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) VisualMenuOption() AwsDashboard_VisualMenuOptionPropertyOutputReference {
	var returns AwsDashboard_VisualMenuOptionPropertyOutputReference
	_jsii_.Get(
		j,
		"visualMenuOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) VisualMenuOptionInput() *AwsDashboard_VisualMenuOptionProperty {
	var returns *AwsDashboard_VisualMenuOptionProperty
	_jsii_.Get(
		j,
		"visualMenuOptionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsDashboard_DashboardPublishOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsDashboard_DashboardPublishOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsDashboard_DashboardPublishOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsDashboard.DashboardPublishOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsDashboard_DashboardPublishOptionsPropertyOutputReference_Override(a AwsDashboard_DashboardPublishOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsDashboard.DashboardPublishOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference)SetInternalValue(val *AwsDashboard_DashboardPublishOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) PutAdHocFilteringOption(value *AwsDashboard_AdHocFilteringOptionProperty) {
	if err := a.validatePutAdHocFilteringOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAdHocFilteringOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) PutDataPointDrillUpDownOption(value *AwsDashboard_DataPointDrillUpDownOptionProperty) {
	if err := a.validatePutDataPointDrillUpDownOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataPointDrillUpDownOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) PutDataPointMenuLabelOption(value *AwsDashboard_DataPointMenuLabelOptionProperty) {
	if err := a.validatePutDataPointMenuLabelOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataPointMenuLabelOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) PutDataPointTooltipOption(value *AwsDashboard_DataPointTooltipOptionProperty) {
	if err := a.validatePutDataPointTooltipOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataPointTooltipOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) PutExportToCsvOption(value *AwsDashboard_ExportToCsvOptionProperty) {
	if err := a.validatePutExportToCsvOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExportToCsvOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) PutExportWithHiddenFieldsOption(value *AwsDashboard_ExportWithHiddenFieldsOptionProperty) {
	if err := a.validatePutExportWithHiddenFieldsOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putExportWithHiddenFieldsOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) PutSheetControlsOption(value *AwsDashboard_SheetControlsOptionProperty) {
	if err := a.validatePutSheetControlsOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSheetControlsOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) PutSheetLayoutElementMaximizationOption(value *AwsDashboard_SheetLayoutElementMaximizationOptionProperty) {
	if err := a.validatePutSheetLayoutElementMaximizationOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSheetLayoutElementMaximizationOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) PutVisualAxisSortOption(value *AwsDashboard_VisualAxisSortOptionProperty) {
	if err := a.validatePutVisualAxisSortOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVisualAxisSortOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) PutVisualMenuOption(value *AwsDashboard_VisualMenuOptionProperty) {
	if err := a.validatePutVisualMenuOptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVisualMenuOption",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ResetAdHocFilteringOption() {
	_jsii_.InvokeVoid(
		a,
		"resetAdHocFilteringOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ResetDataPointDrillUpDownOption() {
	_jsii_.InvokeVoid(
		a,
		"resetDataPointDrillUpDownOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ResetDataPointMenuLabelOption() {
	_jsii_.InvokeVoid(
		a,
		"resetDataPointMenuLabelOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ResetDataPointTooltipOption() {
	_jsii_.InvokeVoid(
		a,
		"resetDataPointTooltipOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ResetExportToCsvOption() {
	_jsii_.InvokeVoid(
		a,
		"resetExportToCsvOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ResetExportWithHiddenFieldsOption() {
	_jsii_.InvokeVoid(
		a,
		"resetExportWithHiddenFieldsOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ResetSheetControlsOption() {
	_jsii_.InvokeVoid(
		a,
		"resetSheetControlsOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ResetSheetLayoutElementMaximizationOption() {
	_jsii_.InvokeVoid(
		a,
		"resetSheetLayoutElementMaximizationOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ResetVisualAxisSortOption() {
	_jsii_.InvokeVoid(
		a,
		"resetVisualAxisSortOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ResetVisualMenuOption() {
	_jsii_.InvokeVoid(
		a,
		"resetVisualMenuOption",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsDashboard_DashboardPublishOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

