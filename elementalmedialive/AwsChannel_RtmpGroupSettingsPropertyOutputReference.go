package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_RtmpGroupSettingsPropertyOutputReference interface {
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
	InternalValue() *AwsChannel_RtmpGroupSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_RtmpGroupSettingsProperty)
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

// The jsii proxy struct for AwsChannel_RtmpGroupSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) AdMarkers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adMarkers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) AdMarkersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"adMarkersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) AuthenticationScheme() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationScheme",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) AuthenticationSchemeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authenticationSchemeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) CacheFullBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheFullBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) CacheFullBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheFullBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) CacheLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) CacheLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cacheLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) CaptionData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"captionData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) CaptionDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"captionDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) InputLossAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) InputLossActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inputLossActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) InternalValue() *AwsChannel_RtmpGroupSettingsProperty {
	var returns *AwsChannel_RtmpGroupSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) RestartDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restartDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) RestartDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restartDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_RtmpGroupSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_RtmpGroupSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_RtmpGroupSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.RtmpGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_RtmpGroupSettingsPropertyOutputReference_Override(a AwsChannel_RtmpGroupSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.RtmpGroupSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference)SetAdMarkers(val *[]*string) {
	if err := j.validateSetAdMarkersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adMarkers",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference)SetAuthenticationScheme(val *string) {
	if err := j.validateSetAuthenticationSchemeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authenticationScheme",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference)SetCacheFullBehavior(val *string) {
	if err := j.validateSetCacheFullBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheFullBehavior",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference)SetCacheLength(val *float64) {
	if err := j.validateSetCacheLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheLength",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference)SetCaptionData(val *string) {
	if err := j.validateSetCaptionDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"captionData",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference)SetInputLossAction(val *string) {
	if err := j.validateSetInputLossActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputLossAction",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_RtmpGroupSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference)SetRestartDelay(val *float64) {
	if err := j.validateSetRestartDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartDelay",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) ResetAdMarkers() {
	_jsii_.InvokeVoid(
		a,
		"resetAdMarkers",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) ResetAuthenticationScheme() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthenticationScheme",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) ResetCacheFullBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheFullBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) ResetCacheLength() {
	_jsii_.InvokeVoid(
		a,
		"resetCacheLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) ResetCaptionData() {
	_jsii_.InvokeVoid(
		a,
		"resetCaptionData",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) ResetInputLossAction() {
	_jsii_.InvokeVoid(
		a,
		"resetInputLossAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) ResetRestartDelay() {
	_jsii_.InvokeVoid(
		a,
		"resetRestartDelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_RtmpGroupSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

