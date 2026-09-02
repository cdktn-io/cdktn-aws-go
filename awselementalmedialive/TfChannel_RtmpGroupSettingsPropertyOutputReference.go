package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_RtmpGroupSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdMarkers() *[]*string
	// Experimental.
	SetAdMarkers(val *[]*string)
	// Experimental.
	AdMarkersInput() *[]*string
	// Experimental.
	AuthenticationScheme() *string
	// Experimental.
	SetAuthenticationScheme(val *string)
	// Experimental.
	AuthenticationSchemeInput() *string
	// Experimental.
	CacheFullBehavior() *string
	// Experimental.
	SetCacheFullBehavior(val *string)
	// Experimental.
	CacheFullBehaviorInput() *string
	// Experimental.
	CacheLength() *float64
	// Experimental.
	SetCacheLength(val *float64)
	// Experimental.
	CacheLengthInput() *float64
	// Experimental.
	CaptionData() *string
	// Experimental.
	SetCaptionData(val *string)
	// Experimental.
	CaptionDataInput() *string
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
	InputLossAction() *string
	// Experimental.
	SetInputLossAction(val *string)
	// Experimental.
	InputLossActionInput() *string
	// Experimental.
	InternalValue() *TfChannel_RtmpGroupSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_RtmpGroupSettingsProperty)
	// Experimental.
	RestartDelay() *float64
	// Experimental.
	SetRestartDelay(val *float64)
	// Experimental.
	RestartDelayInput() *float64
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
	ResetAdMarkers()
	// Experimental.
	ResetAuthenticationScheme()
	// Experimental.
	ResetCacheFullBehavior()
	// Experimental.
	ResetCacheLength()
	// Experimental.
	ResetCaptionData()
	// Experimental.
	ResetInputLossAction()
	// Experimental.
	ResetRestartDelay()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_RtmpGroupSettingsPropertyOutputReference
type jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) AdMarkers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adMarkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) AdMarkersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adMarkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) AuthenticationScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) AuthenticationSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) CacheFullBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheFullBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) CacheFullBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheFullBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) CacheLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) CacheLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) CaptionData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"captionData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) CaptionDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"captionDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) InputLossAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) InputLossActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) InternalValue() *TfChannel_RtmpGroupSettingsProperty {
	var returns *TfChannel_RtmpGroupSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) RestartDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restartDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) RestartDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restartDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_RtmpGroupSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_RtmpGroupSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_RtmpGroupSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.RtmpGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_RtmpGroupSettingsPropertyOutputReference_Override(t TfChannel_RtmpGroupSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.RtmpGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference)SetAdMarkers(val *[]*string) {
	if err := j.validateSetAdMarkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adMarkers",
		val,
	)
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference)SetAuthenticationScheme(val *string) {
	if err := j.validateSetAuthenticationSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationScheme",
		val,
	)
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference)SetCacheFullBehavior(val *string) {
	if err := j.validateSetCacheFullBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheFullBehavior",
		val,
	)
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference)SetCacheLength(val *float64) {
	if err := j.validateSetCacheLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheLength",
		val,
	)
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference)SetCaptionData(val *string) {
	if err := j.validateSetCaptionDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captionData",
		val,
	)
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference)SetInputLossAction(val *string) {
	if err := j.validateSetInputLossActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputLossAction",
		val,
	)
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_RtmpGroupSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference)SetRestartDelay(val *float64) {
	if err := j.validateSetRestartDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartDelay",
		val,
	)
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) ResetAdMarkers() {
	_jsii_.InvokeVoid(
		t,
		"resetAdMarkers",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) ResetAuthenticationScheme() {
	_jsii_.InvokeVoid(
		t,
		"resetAuthenticationScheme",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) ResetCacheFullBehavior() {
	_jsii_.InvokeVoid(
		t,
		"resetCacheFullBehavior",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) ResetCacheLength() {
	_jsii_.InvokeVoid(
		t,
		"resetCacheLength",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) ResetCaptionData() {
	_jsii_.InvokeVoid(
		t,
		"resetCaptionData",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) ResetInputLossAction() {
	_jsii_.InvokeVoid(
		t,
		"resetInputLossAction",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) ResetRestartDelay() {
	_jsii_.InvokeVoid(
		t,
		"resetRestartDelay",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_RtmpGroupSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

