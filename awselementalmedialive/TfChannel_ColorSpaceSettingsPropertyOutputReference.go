package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_ColorSpaceSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ColorSpacePassthroughSettings() TfChannel_ColorSpacePassthroughSettingsPropertyOutputReference
	// Experimental.
	ColorSpacePassthroughSettingsInput() *TfChannel_ColorSpacePassthroughSettingsProperty
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
	DolbyVision81Settings() TfChannel_DolbyVision81SettingsPropertyOutputReference
	// Experimental.
	DolbyVision81SettingsInput() *TfChannel_DolbyVision81SettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Hdr10Settings() TfChannel_Hdr10SettingsPropertyOutputReference
	// Experimental.
	Hdr10SettingsInput() *TfChannel_Hdr10SettingsProperty
	// Experimental.
	InternalValue() *TfChannel_ColorSpaceSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_ColorSpaceSettingsProperty)
	// Experimental.
	Rec601Settings() TfChannel_Rec601SettingsPropertyOutputReference
	// Experimental.
	Rec601SettingsInput() *TfChannel_Rec601SettingsProperty
	// Experimental.
	Rec709Settings() TfChannel_Rec709SettingsPropertyOutputReference
	// Experimental.
	Rec709SettingsInput() *TfChannel_Rec709SettingsProperty
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
	PutColorSpacePassthroughSettings(value *TfChannel_ColorSpacePassthroughSettingsProperty)
	// Experimental.
	PutDolbyVision81Settings(value *TfChannel_DolbyVision81SettingsProperty)
	// Experimental.
	PutHdr10Settings(value *TfChannel_Hdr10SettingsProperty)
	// Experimental.
	PutRec601Settings(value *TfChannel_Rec601SettingsProperty)
	// Experimental.
	PutRec709Settings(value *TfChannel_Rec709SettingsProperty)
	// Experimental.
	ResetColorSpacePassthroughSettings()
	// Experimental.
	ResetDolbyVision81Settings()
	// Experimental.
	ResetHdr10Settings()
	// Experimental.
	ResetRec601Settings()
	// Experimental.
	ResetRec709Settings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_ColorSpaceSettingsPropertyOutputReference
type jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) ColorSpacePassthroughSettings() TfChannel_ColorSpacePassthroughSettingsPropertyOutputReference {
	var returns TfChannel_ColorSpacePassthroughSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"colorSpacePassthroughSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) ColorSpacePassthroughSettingsInput() *TfChannel_ColorSpacePassthroughSettingsProperty {
	var returns *TfChannel_ColorSpacePassthroughSettingsProperty
	_jsii_.Get(
		j,
		"colorSpacePassthroughSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) DolbyVision81Settings() TfChannel_DolbyVision81SettingsPropertyOutputReference {
	var returns TfChannel_DolbyVision81SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"dolbyVision81Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) DolbyVision81SettingsInput() *TfChannel_DolbyVision81SettingsProperty {
	var returns *TfChannel_DolbyVision81SettingsProperty
	_jsii_.Get(
		j,
		"dolbyVision81SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) Hdr10Settings() TfChannel_Hdr10SettingsPropertyOutputReference {
	var returns TfChannel_Hdr10SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hdr10Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) Hdr10SettingsInput() *TfChannel_Hdr10SettingsProperty {
	var returns *TfChannel_Hdr10SettingsProperty
	_jsii_.Get(
		j,
		"hdr10SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) InternalValue() *TfChannel_ColorSpaceSettingsProperty {
	var returns *TfChannel_ColorSpaceSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) Rec601Settings() TfChannel_Rec601SettingsPropertyOutputReference {
	var returns TfChannel_Rec601SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rec601Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) Rec601SettingsInput() *TfChannel_Rec601SettingsProperty {
	var returns *TfChannel_Rec601SettingsProperty
	_jsii_.Get(
		j,
		"rec601SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) Rec709Settings() TfChannel_Rec709SettingsPropertyOutputReference {
	var returns TfChannel_Rec709SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rec709Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) Rec709SettingsInput() *TfChannel_Rec709SettingsProperty {
	var returns *TfChannel_Rec709SettingsProperty
	_jsii_.Get(
		j,
		"rec709SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_ColorSpaceSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_ColorSpaceSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_ColorSpaceSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.ColorSpaceSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_ColorSpaceSettingsPropertyOutputReference_Override(t TfChannel_ColorSpaceSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.ColorSpaceSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_ColorSpaceSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) PutColorSpacePassthroughSettings(value *TfChannel_ColorSpacePassthroughSettingsProperty) {
	if err := t.validatePutColorSpacePassthroughSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putColorSpacePassthroughSettings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) PutDolbyVision81Settings(value *TfChannel_DolbyVision81SettingsProperty) {
	if err := t.validatePutDolbyVision81SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDolbyVision81Settings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) PutHdr10Settings(value *TfChannel_Hdr10SettingsProperty) {
	if err := t.validatePutHdr10SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHdr10Settings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) PutRec601Settings(value *TfChannel_Rec601SettingsProperty) {
	if err := t.validatePutRec601SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRec601Settings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) PutRec709Settings(value *TfChannel_Rec709SettingsProperty) {
	if err := t.validatePutRec709SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRec709Settings",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) ResetColorSpacePassthroughSettings() {
	_jsii_.InvokeVoid(
		t,
		"resetColorSpacePassthroughSettings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) ResetDolbyVision81Settings() {
	_jsii_.InvokeVoid(
		t,
		"resetDolbyVision81Settings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) ResetHdr10Settings() {
	_jsii_.InvokeVoid(
		t,
		"resetHdr10Settings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) ResetRec601Settings() {
	_jsii_.InvokeVoid(
		t,
		"resetRec601Settings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) ResetRec709Settings() {
	_jsii_.InvokeVoid(
		t,
		"resetRec709Settings",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_ColorSpaceSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

