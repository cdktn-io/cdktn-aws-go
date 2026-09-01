package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsQuicksightTheme_UiColorPalettePropertyOutputReference interface {
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
	InternalValue() *AwsQuicksightTheme_UiColorPaletteProperty
	// Experimental.
	SetInternalValue(val *AwsQuicksightTheme_UiColorPaletteProperty)
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

// The jsii proxy struct for AwsQuicksightTheme_UiColorPalettePropertyOutputReference
type jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) Accent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) AccentForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accentForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) AccentForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accentForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) AccentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) Danger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"danger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) DangerForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dangerForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) DangerForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dangerForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) DangerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dangerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) Dimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) DimensionForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) DimensionForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) DimensionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) InternalValue() *AwsQuicksightTheme_UiColorPaletteProperty {
	var returns *AwsQuicksightTheme_UiColorPaletteProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) Measure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) MeasureForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) MeasureForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) MeasureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) PrimaryBackground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) PrimaryBackgroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBackgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) PrimaryForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) PrimaryForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) SecondaryBackground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) SecondaryBackgroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBackgroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) SecondaryForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) SecondaryForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) Success() *string {
	var returns *string
	_jsii_.Get(
		j,
		"success",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) SuccessForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) SuccessForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) SuccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) Warning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) WarningForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) WarningForegroundInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningForegroundInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) WarningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsQuicksightTheme_UiColorPalettePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsQuicksightTheme_UiColorPalettePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsQuicksightTheme_UiColorPalettePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightTheme.UiColorPalettePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsQuicksightTheme_UiColorPalettePropertyOutputReference_Override(a AwsQuicksightTheme_UiColorPalettePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightTheme.UiColorPalettePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetAccent(val *string) {
	if err := j.validateSetAccentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accent",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetAccentForeground(val *string) {
	if err := j.validateSetAccentForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accentForeground",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetDanger(val *string) {
	if err := j.validateSetDangerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"danger",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetDangerForeground(val *string) {
	if err := j.validateSetDangerForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dangerForeground",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetDimension(val *string) {
	if err := j.validateSetDimensionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimension",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetDimensionForeground(val *string) {
	if err := j.validateSetDimensionForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dimensionForeground",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetInternalValue(val *AwsQuicksightTheme_UiColorPaletteProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetMeasure(val *string) {
	if err := j.validateSetMeasureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measure",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetMeasureForeground(val *string) {
	if err := j.validateSetMeasureForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"measureForeground",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetPrimaryBackground(val *string) {
	if err := j.validateSetPrimaryBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryBackground",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetPrimaryForeground(val *string) {
	if err := j.validateSetPrimaryForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primaryForeground",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetSecondaryBackground(val *string) {
	if err := j.validateSetSecondaryBackgroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryBackground",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetSecondaryForeground(val *string) {
	if err := j.validateSetSecondaryForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secondaryForeground",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetSuccess(val *string) {
	if err := j.validateSetSuccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"success",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetSuccessForeground(val *string) {
	if err := j.validateSetSuccessForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"successForeground",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetWarning(val *string) {
	if err := j.validateSetWarningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warning",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference)SetWarningForeground(val *string) {
	if err := j.validateSetWarningForegroundParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"warningForeground",
		val,
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetAccent() {
	_jsii_.InvokeVoid(
		a,
		"resetAccent",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetAccentForeground() {
	_jsii_.InvokeVoid(
		a,
		"resetAccentForeground",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetDanger() {
	_jsii_.InvokeVoid(
		a,
		"resetDanger",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetDangerForeground() {
	_jsii_.InvokeVoid(
		a,
		"resetDangerForeground",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetDimension() {
	_jsii_.InvokeVoid(
		a,
		"resetDimension",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetDimensionForeground() {
	_jsii_.InvokeVoid(
		a,
		"resetDimensionForeground",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetMeasure() {
	_jsii_.InvokeVoid(
		a,
		"resetMeasure",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetMeasureForeground() {
	_jsii_.InvokeVoid(
		a,
		"resetMeasureForeground",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetPrimaryBackground() {
	_jsii_.InvokeVoid(
		a,
		"resetPrimaryBackground",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetPrimaryForeground() {
	_jsii_.InvokeVoid(
		a,
		"resetPrimaryForeground",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetSecondaryBackground() {
	_jsii_.InvokeVoid(
		a,
		"resetSecondaryBackground",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetSecondaryForeground() {
	_jsii_.InvokeVoid(
		a,
		"resetSecondaryForeground",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetSuccess() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccess",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetSuccessForeground() {
	_jsii_.InvokeVoid(
		a,
		"resetSuccessForeground",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetWarning() {
	_jsii_.InvokeVoid(
		a,
		"resetWarning",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ResetWarningForeground() {
	_jsii_.InvokeVoid(
		a,
		"resetWarningForeground",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsQuicksightTheme_UiColorPalettePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

