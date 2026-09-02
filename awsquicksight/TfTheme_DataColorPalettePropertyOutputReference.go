package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTheme_DataColorPalettePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Colors() *[]*string
	// Experimental.
	SetColors(val *[]*string)
	// Experimental.
	ColorsInput() *[]*string
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
	EmptyFillColor() *string
	// Experimental.
	SetEmptyFillColor(val *string)
	// Experimental.
	EmptyFillColorInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfTheme_DataColorPaletteProperty
	// Experimental.
	SetInternalValue(val *TfTheme_DataColorPaletteProperty)
	// Experimental.
	MinMaxGradient() *[]*string
	// Experimental.
	SetMinMaxGradient(val *[]*string)
	// Experimental.
	MinMaxGradientInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
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
	ResetColors()
	// Experimental.
	ResetEmptyFillColor()
	// Experimental.
	ResetMinMaxGradient()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTheme_DataColorPalettePropertyOutputReference
type jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) Colors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"colors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) ColorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"colorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) EmptyFillColor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyFillColor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) EmptyFillColorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyFillColorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) InternalValue() *TfTheme_DataColorPaletteProperty {
	var returns *TfTheme_DataColorPaletteProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) MinMaxGradient() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"minMaxGradient",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) MinMaxGradientInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"minMaxGradientInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTheme_DataColorPalettePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTheme_DataColorPalettePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTheme_DataColorPalettePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfTheme.DataColorPalettePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTheme_DataColorPalettePropertyOutputReference_Override(t TfTheme_DataColorPalettePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfTheme.DataColorPalettePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference)SetColors(val *[]*string) {
	if err := j.validateSetColorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"colors",
		val,
	)
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference)SetEmptyFillColor(val *string) {
	if err := j.validateSetEmptyFillColorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emptyFillColor",
		val,
	)
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference)SetInternalValue(val *TfTheme_DataColorPaletteProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference)SetMinMaxGradient(val *[]*string) {
	if err := j.validateSetMinMaxGradientParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minMaxGradient",
		val,
	)
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) ResetColors() {
	_jsii_.InvokeVoid(
		t,
		"resetColors",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) ResetEmptyFillColor() {
	_jsii_.InvokeVoid(
		t,
		"resetEmptyFillColor",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) ResetMinMaxGradient() {
	_jsii_.InvokeVoid(
		t,
		"resetMinMaxGradient",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTheme_DataColorPalettePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

