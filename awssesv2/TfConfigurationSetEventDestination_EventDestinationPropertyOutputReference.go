package awssesv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssesv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssesv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudWatchDestination() TfConfigurationSetEventDestination_CloudWatchDestinationPropertyOutputReference
	// Experimental.
	CloudWatchDestinationInput() *TfConfigurationSetEventDestination_CloudWatchDestinationProperty
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
	Enabled() interface{}
	// Experimental.
	SetEnabled(val interface{})
	// Experimental.
	EnabledInput() interface{}
	// Experimental.
	EventBridgeDestination() TfConfigurationSetEventDestination_EventBridgeDestinationPropertyOutputReference
	// Experimental.
	EventBridgeDestinationInput() *TfConfigurationSetEventDestination_EventBridgeDestinationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfConfigurationSetEventDestination_EventDestinationProperty
	// Experimental.
	SetInternalValue(val *TfConfigurationSetEventDestination_EventDestinationProperty)
	// Experimental.
	KinesisFirehoseDestination() TfConfigurationSetEventDestination_KinesisFirehoseDestinationPropertyOutputReference
	// Experimental.
	KinesisFirehoseDestinationInput() *TfConfigurationSetEventDestination_KinesisFirehoseDestinationProperty
	// Experimental.
	MatchingEventTypes() *[]*string
	// Experimental.
	SetMatchingEventTypes(val *[]*string)
	// Experimental.
	MatchingEventTypesInput() *[]*string
	// Experimental.
	PinpointDestination() TfConfigurationSetEventDestination_PinpointDestinationPropertyOutputReference
	// Experimental.
	PinpointDestinationInput() *TfConfigurationSetEventDestination_PinpointDestinationProperty
	// Experimental.
	SnsDestination() TfConfigurationSetEventDestination_SnsDestinationPropertyOutputReference
	// Experimental.
	SnsDestinationInput() *TfConfigurationSetEventDestination_SnsDestinationProperty
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
	PutCloudWatchDestination(value *TfConfigurationSetEventDestination_CloudWatchDestinationProperty)
	// Experimental.
	PutEventBridgeDestination(value *TfConfigurationSetEventDestination_EventBridgeDestinationProperty)
	// Experimental.
	PutKinesisFirehoseDestination(value *TfConfigurationSetEventDestination_KinesisFirehoseDestinationProperty)
	// Experimental.
	PutPinpointDestination(value *TfConfigurationSetEventDestination_PinpointDestinationProperty)
	// Experimental.
	PutSnsDestination(value *TfConfigurationSetEventDestination_SnsDestinationProperty)
	// Experimental.
	ResetCloudWatchDestination()
	// Experimental.
	ResetEnabled()
	// Experimental.
	ResetEventBridgeDestination()
	// Experimental.
	ResetKinesisFirehoseDestination()
	// Experimental.
	ResetPinpointDestination()
	// Experimental.
	ResetSnsDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference
type jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) CloudWatchDestination() TfConfigurationSetEventDestination_CloudWatchDestinationPropertyOutputReference {
	var returns TfConfigurationSetEventDestination_CloudWatchDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudWatchDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) CloudWatchDestinationInput() *TfConfigurationSetEventDestination_CloudWatchDestinationProperty {
	var returns *TfConfigurationSetEventDestination_CloudWatchDestinationProperty
	_jsii_.Get(
		j,
		"cloudWatchDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) EventBridgeDestination() TfConfigurationSetEventDestination_EventBridgeDestinationPropertyOutputReference {
	var returns TfConfigurationSetEventDestination_EventBridgeDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"eventBridgeDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) EventBridgeDestinationInput() *TfConfigurationSetEventDestination_EventBridgeDestinationProperty {
	var returns *TfConfigurationSetEventDestination_EventBridgeDestinationProperty
	_jsii_.Get(
		j,
		"eventBridgeDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) InternalValue() *TfConfigurationSetEventDestination_EventDestinationProperty {
	var returns *TfConfigurationSetEventDestination_EventDestinationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) KinesisFirehoseDestination() TfConfigurationSetEventDestination_KinesisFirehoseDestinationPropertyOutputReference {
	var returns TfConfigurationSetEventDestination_KinesisFirehoseDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"kinesisFirehoseDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) KinesisFirehoseDestinationInput() *TfConfigurationSetEventDestination_KinesisFirehoseDestinationProperty {
	var returns *TfConfigurationSetEventDestination_KinesisFirehoseDestinationProperty
	_jsii_.Get(
		j,
		"kinesisFirehoseDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) MatchingEventTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchingEventTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) MatchingEventTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"matchingEventTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PinpointDestination() TfConfigurationSetEventDestination_PinpointDestinationPropertyOutputReference {
	var returns TfConfigurationSetEventDestination_PinpointDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"pinpointDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PinpointDestinationInput() *TfConfigurationSetEventDestination_PinpointDestinationProperty {
	var returns *TfConfigurationSetEventDestination_PinpointDestinationProperty
	_jsii_.Get(
		j,
		"pinpointDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) SnsDestination() TfConfigurationSetEventDestination_SnsDestinationPropertyOutputReference {
	var returns TfConfigurationSetEventDestination_SnsDestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"snsDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) SnsDestinationInput() *TfConfigurationSetEventDestination_SnsDestinationProperty {
	var returns *TfConfigurationSetEventDestination_SnsDestinationProperty
	_jsii_.Get(
		j,
		"snsDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfConfigurationSetEventDestination_EventDestinationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfConfigurationSetEventDestination_EventDestinationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sesv2.TfConfigurationSetEventDestination.EventDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfConfigurationSetEventDestination_EventDestinationPropertyOutputReference_Override(t TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sesv2.TfConfigurationSetEventDestination.EventDestinationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetInternalValue(val *TfConfigurationSetEventDestination_EventDestinationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetMatchingEventTypes(val *[]*string) {
	if err := j.validateSetMatchingEventTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchingEventTypes",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PutCloudWatchDestination(value *TfConfigurationSetEventDestination_CloudWatchDestinationProperty) {
	if err := t.validatePutCloudWatchDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudWatchDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PutEventBridgeDestination(value *TfConfigurationSetEventDestination_EventBridgeDestinationProperty) {
	if err := t.validatePutEventBridgeDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEventBridgeDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PutKinesisFirehoseDestination(value *TfConfigurationSetEventDestination_KinesisFirehoseDestinationProperty) {
	if err := t.validatePutKinesisFirehoseDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putKinesisFirehoseDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PutPinpointDestination(value *TfConfigurationSetEventDestination_PinpointDestinationProperty) {
	if err := t.validatePutPinpointDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPinpointDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) PutSnsDestination(value *TfConfigurationSetEventDestination_SnsDestinationProperty) {
	if err := t.validatePutSnsDestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSnsDestination",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ResetCloudWatchDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudWatchDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ResetEventBridgeDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetEventBridgeDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ResetKinesisFirehoseDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetKinesisFirehoseDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ResetPinpointDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetPinpointDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ResetSnsDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetSnsDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfConfigurationSetEventDestination_EventDestinationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

