package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_HlsAkamaiSettingsPropertyOutputReference interface {
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
	// Experimental.
	ConnectionRetryInterval() *float64
	// Experimental.
	SetConnectionRetryInterval(val *float64)
	// Experimental.
	ConnectionRetryIntervalInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	FilecacheDuration() *float64
	// Experimental.
	SetFilecacheDuration(val *float64)
	// Experimental.
	FilecacheDurationInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	HttpTransferMode() *string
	// Experimental.
	SetHttpTransferMode(val *string)
	// Experimental.
	HttpTransferModeInput() *string
	// Experimental.
	InternalValue() *TfChannel_HlsAkamaiSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_HlsAkamaiSettingsProperty)
	// Experimental.
	NumRetries() *float64
	// Experimental.
	SetNumRetries(val *float64)
	// Experimental.
	NumRetriesInput() *float64
	// Experimental.
	RestartDelay() *float64
	// Experimental.
	SetRestartDelay(val *float64)
	// Experimental.
	RestartDelayInput() *float64
	// Experimental.
	Salt() *string
	// Experimental.
	SetSalt(val *string)
	// Experimental.
	SaltInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Token() *string
	// Experimental.
	SetToken(val *string)
	// Experimental.
	TokenInput() *string
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
	ResetConnectionRetryInterval()
	// Experimental.
	ResetFilecacheDuration()
	// Experimental.
	ResetHttpTransferMode()
	// Experimental.
	ResetNumRetries()
	// Experimental.
	ResetRestartDelay()
	// Experimental.
	ResetSalt()
	// Experimental.
	ResetToken()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_HlsAkamaiSettingsPropertyOutputReference
type jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) ConnectionRetryInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionRetryInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) ConnectionRetryIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionRetryIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) FilecacheDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filecacheDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) FilecacheDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filecacheDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) HttpTransferMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpTransferMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) HttpTransferModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpTransferModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) InternalValue() *TfChannel_HlsAkamaiSettingsProperty {
	var returns *TfChannel_HlsAkamaiSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) NumRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) NumRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) RestartDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restartDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) RestartDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restartDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) Salt() *string {
	var returns *string
	_jsii_.Get(
		j,
		"salt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) SaltInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"saltInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_HlsAkamaiSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_HlsAkamaiSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_HlsAkamaiSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.HlsAkamaiSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_HlsAkamaiSettingsPropertyOutputReference_Override(t TfChannel_HlsAkamaiSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.HlsAkamaiSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference)SetConnectionRetryInterval(val *float64) {
	if err := j.validateSetConnectionRetryIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionRetryInterval",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference)SetFilecacheDuration(val *float64) {
	if err := j.validateSetFilecacheDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filecacheDuration",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference)SetHttpTransferMode(val *string) {
	if err := j.validateSetHttpTransferModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpTransferMode",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_HlsAkamaiSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference)SetNumRetries(val *float64) {
	if err := j.validateSetNumRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numRetries",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference)SetRestartDelay(val *float64) {
	if err := j.validateSetRestartDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartDelay",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference)SetSalt(val *string) {
	if err := j.validateSetSaltParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"salt",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference)SetToken(val *string) {
	if err := j.validateSetTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) ResetConnectionRetryInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetConnectionRetryInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) ResetFilecacheDuration() {
	_jsii_.InvokeVoid(
		t,
		"resetFilecacheDuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) ResetHttpTransferMode() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpTransferMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) ResetNumRetries() {
	_jsii_.InvokeVoid(
		t,
		"resetNumRetries",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) ResetRestartDelay() {
	_jsii_.InvokeVoid(
		t,
		"resetRestartDelay",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) ResetSalt() {
	_jsii_.InvokeVoid(
		t,
		"resetSalt",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) ResetToken() {
	_jsii_.InvokeVoid(
		t,
		"resetToken",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_HlsAkamaiSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

