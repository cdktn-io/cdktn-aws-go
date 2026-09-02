package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfTheme_ConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
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
	DataColorPalette() TfTheme_DataColorPalettePropertyOutputReference
	// Experimental.
	DataColorPaletteInput() *TfTheme_DataColorPaletteProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfTheme_ConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfTheme_ConfigurationProperty)
	// Experimental.
	Sheet() TfTheme_SheetPropertyOutputReference
	// Experimental.
	SheetInput() *TfTheme_SheetProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Typography() TfTheme_TypographyPropertyOutputReference
	// Experimental.
	TypographyInput() *TfTheme_TypographyProperty
	// Experimental.
	UiColorPalette() TfTheme_UiColorPalettePropertyOutputReference
	// Experimental.
	UiColorPaletteInput() *TfTheme_UiColorPaletteProperty
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
	PutDataColorPalette(value *TfTheme_DataColorPaletteProperty)
	// Experimental.
	PutSheet(value *TfTheme_SheetProperty)
	// Experimental.
	PutTypography(value *TfTheme_TypographyProperty)
	// Experimental.
	PutUiColorPalette(value *TfTheme_UiColorPaletteProperty)
	// Experimental.
	ResetDataColorPalette()
	// Experimental.
	ResetSheet()
	// Experimental.
	ResetTypography()
	// Experimental.
	ResetUiColorPalette()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfTheme_ConfigurationPropertyOutputReference
type jsiiProxy_TfTheme_ConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) DataColorPalette() TfTheme_DataColorPalettePropertyOutputReference {
	var returns TfTheme_DataColorPalettePropertyOutputReference
	_jsii_.Get(
		j,
		"dataColorPalette",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) DataColorPaletteInput() *TfTheme_DataColorPaletteProperty {
	var returns *TfTheme_DataColorPaletteProperty
	_jsii_.Get(
		j,
		"dataColorPaletteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) InternalValue() *TfTheme_ConfigurationProperty {
	var returns *TfTheme_ConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) Sheet() TfTheme_SheetPropertyOutputReference {
	var returns TfTheme_SheetPropertyOutputReference
	_jsii_.Get(
		j,
		"sheet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) SheetInput() *TfTheme_SheetProperty {
	var returns *TfTheme_SheetProperty
	_jsii_.Get(
		j,
		"sheetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) Typography() TfTheme_TypographyPropertyOutputReference {
	var returns TfTheme_TypographyPropertyOutputReference
	_jsii_.Get(
		j,
		"typography",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) TypographyInput() *TfTheme_TypographyProperty {
	var returns *TfTheme_TypographyProperty
	_jsii_.Get(
		j,
		"typographyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) UiColorPalette() TfTheme_UiColorPalettePropertyOutputReference {
	var returns TfTheme_UiColorPalettePropertyOutputReference
	_jsii_.Get(
		j,
		"uiColorPalette",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) UiColorPaletteInput() *TfTheme_UiColorPaletteProperty {
	var returns *TfTheme_UiColorPaletteProperty
	_jsii_.Get(
		j,
		"uiColorPaletteInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfTheme_ConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfTheme_ConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfTheme_ConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfTheme_ConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfTheme.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfTheme_ConfigurationPropertyOutputReference_Override(t TfTheme_ConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.TfTheme.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference)SetInternalValue(val *TfTheme_ConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) PutDataColorPalette(value *TfTheme_DataColorPaletteProperty) {
	if err := t.validatePutDataColorPaletteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataColorPalette",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) PutSheet(value *TfTheme_SheetProperty) {
	if err := t.validatePutSheetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSheet",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) PutTypography(value *TfTheme_TypographyProperty) {
	if err := t.validatePutTypographyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTypography",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) PutUiColorPalette(value *TfTheme_UiColorPaletteProperty) {
	if err := t.validatePutUiColorPaletteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUiColorPalette",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) ResetDataColorPalette() {
	_jsii_.InvokeVoid(
		t,
		"resetDataColorPalette",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) ResetSheet() {
	_jsii_.InvokeVoid(
		t,
		"resetSheet",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) ResetTypography() {
	_jsii_.InvokeVoid(
		t,
		"resetTypography",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) ResetUiColorPalette() {
	_jsii_.InvokeVoid(
		t,
		"resetUiColorPalette",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfTheme_ConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

