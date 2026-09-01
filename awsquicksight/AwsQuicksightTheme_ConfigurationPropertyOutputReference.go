package awsquicksight

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsquicksight/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsquicksight/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsQuicksightTheme_ConfigurationPropertyOutputReference interface {
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
	DataColorPalette() AwsQuicksightTheme_DataColorPalettePropertyOutputReference
	// Experimental.
	DataColorPaletteInput() *AwsQuicksightTheme_DataColorPaletteProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsQuicksightTheme_ConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsQuicksightTheme_ConfigurationProperty)
	// Experimental.
	Sheet() AwsQuicksightTheme_SheetPropertyOutputReference
	// Experimental.
	SheetInput() *AwsQuicksightTheme_SheetProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Typography() AwsQuicksightTheme_TypographyPropertyOutputReference
	// Experimental.
	TypographyInput() *AwsQuicksightTheme_TypographyProperty
	// Experimental.
	UiColorPalette() AwsQuicksightTheme_UiColorPalettePropertyOutputReference
	// Experimental.
	UiColorPaletteInput() *AwsQuicksightTheme_UiColorPaletteProperty
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
	PutDataColorPalette(value *AwsQuicksightTheme_DataColorPaletteProperty)
	// Experimental.
	PutSheet(value *AwsQuicksightTheme_SheetProperty)
	// Experimental.
	PutTypography(value *AwsQuicksightTheme_TypographyProperty)
	// Experimental.
	PutUiColorPalette(value *AwsQuicksightTheme_UiColorPaletteProperty)
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

// The jsii proxy struct for AwsQuicksightTheme_ConfigurationPropertyOutputReference
type jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) DataColorPalette() AwsQuicksightTheme_DataColorPalettePropertyOutputReference {
	var returns AwsQuicksightTheme_DataColorPalettePropertyOutputReference
	_jsii_.Get(
		j,
		"dataColorPalette",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) DataColorPaletteInput() *AwsQuicksightTheme_DataColorPaletteProperty {
	var returns *AwsQuicksightTheme_DataColorPaletteProperty
	_jsii_.Get(
		j,
		"dataColorPaletteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) InternalValue() *AwsQuicksightTheme_ConfigurationProperty {
	var returns *AwsQuicksightTheme_ConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) Sheet() AwsQuicksightTheme_SheetPropertyOutputReference {
	var returns AwsQuicksightTheme_SheetPropertyOutputReference
	_jsii_.Get(
		j,
		"sheet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) SheetInput() *AwsQuicksightTheme_SheetProperty {
	var returns *AwsQuicksightTheme_SheetProperty
	_jsii_.Get(
		j,
		"sheetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) Typography() AwsQuicksightTheme_TypographyPropertyOutputReference {
	var returns AwsQuicksightTheme_TypographyPropertyOutputReference
	_jsii_.Get(
		j,
		"typography",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) TypographyInput() *AwsQuicksightTheme_TypographyProperty {
	var returns *AwsQuicksightTheme_TypographyProperty
	_jsii_.Get(
		j,
		"typographyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) UiColorPalette() AwsQuicksightTheme_UiColorPalettePropertyOutputReference {
	var returns AwsQuicksightTheme_UiColorPalettePropertyOutputReference
	_jsii_.Get(
		j,
		"uiColorPalette",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) UiColorPaletteInput() *AwsQuicksightTheme_UiColorPaletteProperty {
	var returns *AwsQuicksightTheme_UiColorPaletteProperty
	_jsii_.Get(
		j,
		"uiColorPaletteInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsQuicksightTheme_ConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsQuicksightTheme_ConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsQuicksightTheme_ConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightTheme.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsQuicksightTheme_ConfigurationPropertyOutputReference_Override(a AwsQuicksightTheme_ConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-quicksight.AwsQuicksightTheme.ConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference)SetInternalValue(val *AwsQuicksightTheme_ConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) PutDataColorPalette(value *AwsQuicksightTheme_DataColorPaletteProperty) {
	if err := a.validatePutDataColorPaletteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataColorPalette",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) PutSheet(value *AwsQuicksightTheme_SheetProperty) {
	if err := a.validatePutSheetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSheet",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) PutTypography(value *AwsQuicksightTheme_TypographyProperty) {
	if err := a.validatePutTypographyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTypography",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) PutUiColorPalette(value *AwsQuicksightTheme_UiColorPaletteProperty) {
	if err := a.validatePutUiColorPaletteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUiColorPalette",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) ResetDataColorPalette() {
	_jsii_.InvokeVoid(
		a,
		"resetDataColorPalette",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) ResetSheet() {
	_jsii_.InvokeVoid(
		a,
		"resetSheet",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) ResetTypography() {
	_jsii_.InvokeVoid(
		a,
		"resetTypography",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) ResetUiColorPalette() {
	_jsii_.InvokeVoid(
		a,
		"resetUiColorPalette",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsQuicksightTheme_ConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

