package elementalmedialive

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/elementalmedialive/jsii"

	"github.com/cdktn-io/cdktn-aws-go/elementalmedialive/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsChannel_HlsMediaStoreSettingsPropertyOutputReference interface {
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
	InternalValue() *AwsChannel_HlsMediaStoreSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsChannel_HlsMediaStoreSettingsProperty)
	// Experimental.
	MediaStoreStorageClass() *string
	// Experimental.
	SetMediaStoreStorageClass(val *string)
	// Experimental.
	MediaStoreStorageClassInput() *string
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
	ResetConnectionRetryInterval()
	// Experimental.
	ResetFilecacheDuration()
	// Experimental.
	ResetMediaStoreStorageClass()
	// Experimental.
	ResetNumRetries()
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

// The jsii proxy struct for AwsChannel_HlsMediaStoreSettingsPropertyOutputReference
type jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) ConnectionRetryInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionRetryInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) ConnectionRetryIntervalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"connectionRetryIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) FilecacheDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filecacheDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) FilecacheDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"filecacheDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) InternalValue() *AwsChannel_HlsMediaStoreSettingsProperty {
	var returns *AwsChannel_HlsMediaStoreSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) MediaStoreStorageClass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaStoreStorageClass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) MediaStoreStorageClassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mediaStoreStorageClassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) NumRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) NumRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) RestartDelay() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restartDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) RestartDelayInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"restartDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsChannel_HlsMediaStoreSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsChannel_HlsMediaStoreSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsChannel_HlsMediaStoreSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.HlsMediaStoreSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsChannel_HlsMediaStoreSettingsPropertyOutputReference_Override(a AwsChannel_HlsMediaStoreSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-elemental-medialive.AwsChannel.HlsMediaStoreSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference)SetConnectionRetryInterval(val *float64) {
	if err := j.validateSetConnectionRetryIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionRetryInterval",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference)SetFilecacheDuration(val *float64) {
	if err := j.validateSetFilecacheDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filecacheDuration",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference)SetInternalValue(val *AwsChannel_HlsMediaStoreSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference)SetMediaStoreStorageClass(val *string) {
	if err := j.validateSetMediaStoreStorageClassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mediaStoreStorageClass",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference)SetNumRetries(val *float64) {
	if err := j.validateSetNumRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"numRetries",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference)SetRestartDelay(val *float64) {
	if err := j.validateSetRestartDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restartDelay",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) ResetConnectionRetryInterval() {
	_jsii_.InvokeVoid(
		a,
		"resetConnectionRetryInterval",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) ResetFilecacheDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetFilecacheDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) ResetMediaStoreStorageClass() {
	_jsii_.InvokeVoid(
		a,
		"resetMediaStoreStorageClass",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) ResetNumRetries() {
	_jsii_.InvokeVoid(
		a,
		"resetNumRetries",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) ResetRestartDelay() {
	_jsii_.InvokeVoid(
		a,
		"resetRestartDelay",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsChannel_HlsMediaStoreSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

