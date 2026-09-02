package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTheme_UiColorPalettePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Accent() *string
	// Experimental.
	SetAccent(val *string)
	// Experimental.
	AccentForeground() *string
	// Experimental.
	SetAccentForeground(val *string)
	// Experimental.
	AccentForegroundInput() *string
	// Experimental.
	AccentInput() *string
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
	Danger() *string
	// Experimental.
	SetDanger(val *string)
	// Experimental.
	DangerForeground() *string
	// Experimental.
	SetDangerForeground(val *string)
	// Experimental.
	DangerForegroundInput() *string
	// Experimental.
	DangerInput() *string
	// Experimental.
	Dimension() *string
	// Experimental.
	SetDimension(val *string)
	// Experimental.
	DimensionForeground() *string
	// Experimental.
	SetDimensionForeground(val *string)
	// Experimental.
	DimensionForegroundInput() *string
	// Experimental.
	DimensionInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfTheme_UiColorPaletteProperty
	// Experimental.
	SetInternalValue(val *TfTheme_UiColorPaletteProperty)
	// Experimental.
	Measure() *string
	// Experimental.
	SetMeasure(val *string)
	// Experimental.
	MeasureForeground() *string
	// Experimental.
	SetMeasureForeground(val *string)
	// Experimental.
	MeasureForegroundInput() *string
	// Experimental.
	MeasureInput() *string
	// Experimental.
	PrimaryBackground() *string
	// Experimental.
	SetPrimaryBackground(val *string)
	// Experimental.
	PrimaryBackgroundInput() *string
	// Experimental.
	PrimaryForeground() *string
	// Experimental.
	SetPrimaryForeground(val *string)
	// Experimental.
	PrimaryForegroundInput() *string
	// Experimental.
	SecondaryBackground() *string
	// Experimental.
	SetSecondaryBackground(val *string)
	// Experimental.
	SecondaryBackgroundInput() *string
	// Experimental.
	SecondaryForeground() *string
	// Experimental.
	SetSecondaryForeground(val *string)
	// Experimental.
	SecondaryForegroundInput() *string
	// Experimental.
	Success() *string
	// Experimental.
	SetSuccess(val *string)
	// Experimental.
	SuccessForeground() *string
	// Experimental.
	SetSuccessForeground(val *string)
	// Experimental.
	SuccessForegroundInput() *string
	// Experimental.
	SuccessInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Warning() *string
	// Experimental.
	SetWarning(val *string)
	// Experimental.
	WarningForeground() *string
	// Experimental.
	SetWarningForeground(val *string)
	// Experimental.
	WarningForegroundInput() *string
	// Experimental.
	WarningInput() *string
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
	ResetAccent()
	// Experimental.
	ResetAccentForeground()
	// Experimental.
	ResetDanger()
	// Experimental.
	ResetDangerForeground()
	// Experimental.
	ResetDimension()
	// Experimental.
	ResetDimensionForeground()
	// Experimental.
	ResetMeasure()
	// Experimental.
	ResetMeasureForeground()
	// Experimental.
	ResetPrimaryBackground()
	// Experimental.
	ResetPrimaryForeground()
	// Experimental.
	ResetSecondaryBackground()
	// Experimental.
	ResetSecondaryForeground()
	// Experimental.
	ResetSuccess()
	// Experimental.
	ResetSuccessForeground()
	// Experimental.
	ResetWarning()
	// Experimental.
	ResetWarningForeground()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTheme_UiColorPalettePropertyOutputReference
type jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) Accent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) AccentForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accentForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) AccentForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accentForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) AccentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) Danger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"danger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) DangerForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dangerForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) DangerForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dangerForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) DangerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dangerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) Dimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) DimensionForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) DimensionForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) DimensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) InternalValue() *TfTheme_UiColorPaletteProperty {
	var returns *TfTheme_UiColorPaletteProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) Measure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) MeasureForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) MeasureForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) MeasureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) PrimaryBackground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) PrimaryBackgroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBackgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) PrimaryForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) PrimaryForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) SecondaryBackground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) SecondaryBackgroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBackgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) SecondaryForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) SecondaryForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) Success() *string {
	var returns *string
	_jsii_.Get(
		j,
		"success",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) SuccessForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) SuccessForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) SuccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) Warning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) WarningForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) WarningForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) WarningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTheme_UiColorPalettePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTheme_UiColorPalettePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTheme_UiColorPalettePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfTheme.UiColorPalettePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTheme_UiColorPalettePropertyOutputReference_Override(t TfTheme_UiColorPalettePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfTheme.UiColorPalettePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetAccent(val *string) {
	if err := j.validateSetAccentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accent",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetAccentForeground(val *string) {
	if err := j.validateSetAccentForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accentForeground",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetDanger(val *string) {
	if err := j.validateSetDangerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"danger",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetDangerForeground(val *string) {
	if err := j.validateSetDangerForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dangerForeground",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetDimension(val *string) {
	if err := j.validateSetDimensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimension",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetDimensionForeground(val *string) {
	if err := j.validateSetDimensionForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensionForeground",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetInternalValue(val *TfTheme_UiColorPaletteProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetMeasure(val *string) {
	if err := j.validateSetMeasureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measure",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetMeasureForeground(val *string) {
	if err := j.validateSetMeasureForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureForeground",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetPrimaryBackground(val *string) {
	if err := j.validateSetPrimaryBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryBackground",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetPrimaryForeground(val *string) {
	if err := j.validateSetPrimaryForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryForeground",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetSecondaryBackground(val *string) {
	if err := j.validateSetSecondaryBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryBackground",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetSecondaryForeground(val *string) {
	if err := j.validateSetSecondaryForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryForeground",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetSuccess(val *string) {
	if err := j.validateSetSuccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"success",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetSuccessForeground(val *string) {
	if err := j.validateSetSuccessForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successForeground",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetWarning(val *string) {
	if err := j.validateSetWarningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warning",
		val,
	)
}

func (j *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference)SetWarningForeground(val *string) {
	if err := j.validateSetWarningForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warningForeground",
		val,
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetAccent() {
	_jsii_.InvokeVoid(
		t,
		"resetAccent",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetAccentForeground() {
	_jsii_.InvokeVoid(
		t,
		"resetAccentForeground",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetDanger() {
	_jsii_.InvokeVoid(
		t,
		"resetDanger",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetDangerForeground() {
	_jsii_.InvokeVoid(
		t,
		"resetDangerForeground",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		t,
		"resetDimension",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetDimensionForeground() {
	_jsii_.InvokeVoid(
		t,
		"resetDimensionForeground",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetMeasure() {
	_jsii_.InvokeVoid(
		t,
		"resetMeasure",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetMeasureForeground() {
	_jsii_.InvokeVoid(
		t,
		"resetMeasureForeground",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetPrimaryBackground() {
	_jsii_.InvokeVoid(
		t,
		"resetPrimaryBackground",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetPrimaryForeground() {
	_jsii_.InvokeVoid(
		t,
		"resetPrimaryForeground",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetSecondaryBackground() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondaryBackground",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetSecondaryForeground() {
	_jsii_.InvokeVoid(
		t,
		"resetSecondaryForeground",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetSuccess() {
	_jsii_.InvokeVoid(
		t,
		"resetSuccess",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetSuccessForeground() {
	_jsii_.InvokeVoid(
		t,
		"resetSuccessForeground",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetWarning() {
	_jsii_.InvokeVoid(
		t,
		"resetWarning",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ResetWarningForeground() {
	_jsii_.InvokeVoid(
		t,
		"resetWarningForeground",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTheme_UiColorPalettePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

