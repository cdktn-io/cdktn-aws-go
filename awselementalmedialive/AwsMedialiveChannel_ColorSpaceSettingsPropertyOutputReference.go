package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ColorSpacePassthroughSettings() AwsMedialiveChannel_ColorSpacePassthroughSettingsPropertyOutputReference
	// Experimental.
	ColorSpacePassthroughSettingsInput() *AwsMedialiveChannel_ColorSpacePassthroughSettingsProperty
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
	DolbyVision81Settings() AwsMedialiveChannel_DolbyVision81SettingsPropertyOutputReference
	// Experimental.
	DolbyVision81SettingsInput() *AwsMedialiveChannel_DolbyVision81SettingsProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	Hdr10Settings() AwsMedialiveChannel_Hdr10SettingsPropertyOutputReference
	// Experimental.
	Hdr10SettingsInput() *AwsMedialiveChannel_Hdr10SettingsProperty
	// Experimental.
	InternalValue() *AwsMedialiveChannel_ColorSpaceSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsMedialiveChannel_ColorSpaceSettingsProperty)
	// Experimental.
	Rec601Settings() AwsMedialiveChannel_Rec601SettingsPropertyOutputReference
	// Experimental.
	Rec601SettingsInput() *AwsMedialiveChannel_Rec601SettingsProperty
	// Experimental.
	Rec709Settings() AwsMedialiveChannel_Rec709SettingsPropertyOutputReference
	// Experimental.
	Rec709SettingsInput() *AwsMedialiveChannel_Rec709SettingsProperty
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
	PutColorSpacePassthroughSettings(value *AwsMedialiveChannel_ColorSpacePassthroughSettingsProperty)
	// Experimental.
	PutDolbyVision81Settings(value *AwsMedialiveChannel_DolbyVision81SettingsProperty)
	// Experimental.
	PutHdr10Settings(value *AwsMedialiveChannel_Hdr10SettingsProperty)
	// Experimental.
	PutRec601Settings(value *AwsMedialiveChannel_Rec601SettingsProperty)
	// Experimental.
	PutRec709Settings(value *AwsMedialiveChannel_Rec709SettingsProperty)
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

// The jsii proxy struct for AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference
type jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) ColorSpacePassthroughSettings() AwsMedialiveChannel_ColorSpacePassthroughSettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_ColorSpacePassthroughSettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"colorSpacePassthroughSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) ColorSpacePassthroughSettingsInput() *AwsMedialiveChannel_ColorSpacePassthroughSettingsProperty {
	var returns *AwsMedialiveChannel_ColorSpacePassthroughSettingsProperty
	_jsii_.Get(
		j,
		"colorSpacePassthroughSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) DolbyVision81Settings() AwsMedialiveChannel_DolbyVision81SettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_DolbyVision81SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"dolbyVision81Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) DolbyVision81SettingsInput() *AwsMedialiveChannel_DolbyVision81SettingsProperty {
	var returns *AwsMedialiveChannel_DolbyVision81SettingsProperty
	_jsii_.Get(
		j,
		"dolbyVision81SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) Hdr10Settings() AwsMedialiveChannel_Hdr10SettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_Hdr10SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"hdr10Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) Hdr10SettingsInput() *AwsMedialiveChannel_Hdr10SettingsProperty {
	var returns *AwsMedialiveChannel_Hdr10SettingsProperty
	_jsii_.Get(
		j,
		"hdr10SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) InternalValue() *AwsMedialiveChannel_ColorSpaceSettingsProperty {
	var returns *AwsMedialiveChannel_ColorSpaceSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) Rec601Settings() AwsMedialiveChannel_Rec601SettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_Rec601SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rec601Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) Rec601SettingsInput() *AwsMedialiveChannel_Rec601SettingsProperty {
	var returns *AwsMedialiveChannel_Rec601SettingsProperty
	_jsii_.Get(
		j,
		"rec601SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) Rec709Settings() AwsMedialiveChannel_Rec709SettingsPropertyOutputReference {
	var returns AwsMedialiveChannel_Rec709SettingsPropertyOutputReference
	_jsii_.Get(
		j,
		"rec709Settings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) Rec709SettingsInput() *AwsMedialiveChannel_Rec709SettingsProperty {
	var returns *AwsMedialiveChannel_Rec709SettingsProperty
	_jsii_.Get(
		j,
		"rec709SettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.ColorSpaceSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference_Override(a AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsMedialiveChannel.ColorSpaceSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference)SetInternalValue(val *AwsMedialiveChannel_ColorSpaceSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) PutColorSpacePassthroughSettings(value *AwsMedialiveChannel_ColorSpacePassthroughSettingsProperty) {
	if err := a.validatePutColorSpacePassthroughSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putColorSpacePassthroughSettings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) PutDolbyVision81Settings(value *AwsMedialiveChannel_DolbyVision81SettingsProperty) {
	if err := a.validatePutDolbyVision81SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDolbyVision81Settings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) PutHdr10Settings(value *AwsMedialiveChannel_Hdr10SettingsProperty) {
	if err := a.validatePutHdr10SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHdr10Settings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) PutRec601Settings(value *AwsMedialiveChannel_Rec601SettingsProperty) {
	if err := a.validatePutRec601SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRec601Settings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) PutRec709Settings(value *AwsMedialiveChannel_Rec709SettingsProperty) {
	if err := a.validatePutRec709SettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRec709Settings",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) ResetColorSpacePassthroughSettings() {
	_jsii_.InvokeVoid(
		a,
		"resetColorSpacePassthroughSettings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) ResetDolbyVision81Settings() {
	_jsii_.InvokeVoid(
		a,
		"resetDolbyVision81Settings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) ResetHdr10Settings() {
	_jsii_.InvokeVoid(
		a,
		"resetHdr10Settings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) ResetRec601Settings() {
	_jsii_.InvokeVoid(
		a,
		"resetRec601Settings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) ResetRec709Settings() {
	_jsii_.InvokeVoid(
		a,
		"resetRec709Settings",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsMedialiveChannel_ColorSpaceSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

