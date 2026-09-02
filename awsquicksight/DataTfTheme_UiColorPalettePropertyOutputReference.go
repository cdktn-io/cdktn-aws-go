package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataTfTheme_UiColorPalettePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Accent() *string
	// Experimental.
	AccentForeground() *string
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
	DangerForeground() *string
	// Experimental.
	Dimension() *string
	// Experimental.
	DimensionForeground() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *DataTfTheme_UiColorPaletteProperty
	// Experimental.
	SetInternalValue(val *DataTfTheme_UiColorPaletteProperty)
	// Experimental.
	Measure() *string
	// Experimental.
	MeasureForeground() *string
	// Experimental.
	PrimaryBackground() *string
	// Experimental.
	PrimaryForeground() *string
	// Experimental.
	SecondaryBackground() *string
	// Experimental.
	SecondaryForeground() *string
	// Experimental.
	Success() *string
	// Experimental.
	SuccessForeground() *string
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
	WarningForeground() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataTfTheme_UiColorPalettePropertyOutputReference
type jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) Accent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) AccentForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accentForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) Danger() *string {
	var returns *string
	_jsii_.Get(
		j,
		"danger",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) DangerForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dangerForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) Dimension() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) DimensionForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dimensionForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) InternalValue() *DataTfTheme_UiColorPaletteProperty {
	var returns *DataTfTheme_UiColorPaletteProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) Measure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) MeasureForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"measureForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) PrimaryBackground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) PrimaryForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"primaryForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) SecondaryBackground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryBackground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) SecondaryForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secondaryForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) Success() *string {
	var returns *string
	_jsii_.Get(
		j,
		"success",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) SuccessForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"successForeground",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) Warning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) WarningForeground() *string {
	var returns *string
	_jsii_.Get(
		j,
		"warningForeground",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataTfTheme_UiColorPalettePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataTfTheme_UiColorPalettePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataTfTheme_UiColorPalettePropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.DataTfTheme.UiColorPalettePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataTfTheme_UiColorPalettePropertyOutputReference_Override(d DataTfTheme_UiColorPalettePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.DataTfTheme.UiColorPalettePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference)SetInternalValue(val *DataTfTheme_UiColorPaletteProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataTfTheme_UiColorPalettePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

