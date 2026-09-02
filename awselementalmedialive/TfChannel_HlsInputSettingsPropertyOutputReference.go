package awselementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awselementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfChannel_HlsInputSettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Bandwidth() *float64
	// Experimental.
	SetBandwidth(val *float64)
	// Experimental.
	BandwidthInput() *float64
	// Experimental.
	BufferSegments() *float64
	// Experimental.
	SetBufferSegments(val *float64)
	// Experimental.
	BufferSegmentsInput() *float64
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
	InternalValue() *TfChannel_HlsInputSettingsProperty
	// Experimental.
	SetInternalValue(val *TfChannel_HlsInputSettingsProperty)
	// Experimental.
	Retries() *float64
	// Experimental.
	SetRetries(val *float64)
	// Experimental.
	RetriesInput() *float64
	// Experimental.
	RetryInterval() *float64
	// Experimental.
	SetRetryInterval(val *float64)
	// Experimental.
	RetryIntervalInput() *float64
	// Experimental.
	Scte35Source() *string
	// Experimental.
	SetScte35Source(val *string)
	// Experimental.
	Scte35SourceInput() *string
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
	ResetBandwidth()
	// Experimental.
	ResetBufferSegments()
	// Experimental.
	ResetRetries()
	// Experimental.
	ResetRetryInterval()
	// Experimental.
	ResetScte35Source()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfChannel_HlsInputSettingsPropertyOutputReference
type jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) Bandwidth() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) BandwidthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bandwidthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) BufferSegments() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferSegments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) BufferSegmentsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bufferSegmentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) InternalValue() *TfChannel_HlsInputSettingsProperty {
	var returns *TfChannel_HlsInputSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) Retries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) RetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) RetryInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) RetryIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) Scte35Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte35Source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) Scte35SourceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scte35SourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfChannel_HlsInputSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfChannel_HlsInputSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfChannel_HlsInputSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.HlsInputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfChannel_HlsInputSettingsPropertyOutputReference_Override(t TfChannel_HlsInputSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.TfChannel.HlsInputSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference)SetBandwidth(val *float64) {
	if err := j.validateSetBandwidthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bandwidth",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference)SetBufferSegments(val *float64) {
	if err := j.validateSetBufferSegmentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bufferSegments",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference)SetInternalValue(val *TfChannel_HlsInputSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference)SetRetries(val *float64) {
	if err := j.validateSetRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retries",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference)SetRetryInterval(val *float64) {
	if err := j.validateSetRetryIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryInterval",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference)SetScte35Source(val *string) {
	if err := j.validateSetScte35SourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scte35Source",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) ResetBandwidth() {
	_jsii_.InvokeVoid(
		t,
		"resetBandwidth",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) ResetBufferSegments() {
	_jsii_.InvokeVoid(
		t,
		"resetBufferSegments",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) ResetRetries() {
	_jsii_.InvokeVoid(
		t,
		"resetRetries",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) ResetRetryInterval() {
	_jsii_.InvokeVoid(
		t,
		"resetRetryInterval",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) ResetScte35Source() {
	_jsii_.InvokeVoid(
		t,
		"resetScte35Source",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfChannel_HlsInputSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

