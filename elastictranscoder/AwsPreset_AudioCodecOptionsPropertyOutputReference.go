package elastictranscoder

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elastictranscoder/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elastictranscoder/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPreset_AudioCodecOptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	BitDepth() *string
	// Experimental.
	SetBitDepth(val *string)
	// Experimental.
	BitDepthInput() *string
	// Experimental.
	BitOrder() *string
	// Experimental.
	SetBitOrder(val *string)
	// Experimental.
	BitOrderInput() *string
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
	Fqn() *string
	// Experimental.
	InternalValue() *AwsPreset_AudioCodecOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsPreset_AudioCodecOptionsProperty)
	// Experimental.
	Profile() *string
	// Experimental.
	SetProfile(val *string)
	// Experimental.
	ProfileInput() *string
	// Experimental.
	Signed() *string
	// Experimental.
	SetSigned(val *string)
	// Experimental.
	SignedInput() *string
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
	ResetBitDepth()
	// Experimental.
	ResetBitOrder()
	// Experimental.
	ResetProfile()
	// Experimental.
	ResetSigned()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsPreset_AudioCodecOptionsPropertyOutputReference
type jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) BitDepth() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitDepth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) BitDepthInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitDepthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) BitOrder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitOrder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) BitOrderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bitOrderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) InternalValue() *AwsPreset_AudioCodecOptionsProperty {
	var returns *AwsPreset_AudioCodecOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) Profile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) ProfileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"profileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) Signed() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) SignedInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"signedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPreset_AudioCodecOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPreset_AudioCodecOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPreset_AudioCodecOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elastic-transcoder.AwsPreset.AudioCodecOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPreset_AudioCodecOptionsPropertyOutputReference_Override(a AwsPreset_AudioCodecOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elastic-transcoder.AwsPreset.AudioCodecOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference)SetBitDepth(val *string) {
	if err := j.validateSetBitDepthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitDepth",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference)SetBitOrder(val *string) {
	if err := j.validateSetBitOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bitOrder",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference)SetInternalValue(val *AwsPreset_AudioCodecOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference)SetProfile(val *string) {
	if err := j.validateSetProfileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"profile",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference)SetSigned(val *string) {
	if err := j.validateSetSignedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"signed",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) ResetBitDepth() {
	_jsii_.InvokeVoid(
		a,
		"resetBitDepth",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) ResetBitOrder() {
	_jsii_.InvokeVoid(
		a,
		"resetBitOrder",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) ResetProfile() {
	_jsii_.InvokeVoid(
		a,
		"resetProfile",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) ResetSigned() {
	_jsii_.InvokeVoid(
		a,
		"resetSigned",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPreset_AudioCodecOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

