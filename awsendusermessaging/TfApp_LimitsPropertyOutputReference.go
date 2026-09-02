package awsendusermessaging

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsendusermessaging/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsendusermessaging/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfApp_LimitsPropertyOutputReference interface {
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
	Daily() *float64
	// Experimental.
	SetDaily(val *float64)
	// Experimental.
	DailyInput() *float64
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfApp_LimitsProperty
	// Experimental.
	SetInternalValue(val *TfApp_LimitsProperty)
	// Experimental.
	MaximumDuration() *float64
	// Experimental.
	SetMaximumDuration(val *float64)
	// Experimental.
	MaximumDurationInput() *float64
	// Experimental.
	MessagesPerSecond() *float64
	// Experimental.
	SetMessagesPerSecond(val *float64)
	// Experimental.
	MessagesPerSecondInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Total() *float64
	// Experimental.
	SetTotal(val *float64)
	// Experimental.
	TotalInput() *float64
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
	ResetDaily()
	// Experimental.
	ResetMaximumDuration()
	// Experimental.
	ResetMessagesPerSecond()
	// Experimental.
	ResetTotal()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfApp_LimitsPropertyOutputReference
type jsiiProxy_TfApp_LimitsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) Daily() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) DailyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) InternalValue() *TfApp_LimitsProperty {
	var returns *TfApp_LimitsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) MaximumDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) MaximumDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) MessagesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messagesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) MessagesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messagesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) Total() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"total",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference) TotalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfApp_LimitsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfApp_LimitsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfApp_LimitsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfApp_LimitsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-end-user-messaging.TfApp.LimitsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfApp_LimitsPropertyOutputReference_Override(t TfApp_LimitsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-end-user-messaging.TfApp.LimitsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference)SetDaily(val *float64) {
	if err := j.validateSetDailyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daily",
		val,
	)
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference)SetInternalValue(val *TfApp_LimitsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference)SetMaximumDuration(val *float64) {
	if err := j.validateSetMaximumDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumDuration",
		val,
	)
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference)SetMessagesPerSecond(val *float64) {
	if err := j.validateSetMessagesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messagesPerSecond",
		val,
	)
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfApp_LimitsPropertyOutputReference)SetTotal(val *float64) {
	if err := j.validateSetTotalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"total",
		val,
	)
}

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) ResetDaily() {
	_jsii_.InvokeVoid(
		t,
		"resetDaily",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) ResetMaximumDuration() {
	_jsii_.InvokeVoid(
		t,
		"resetMaximumDuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) ResetMessagesPerSecond() {
	_jsii_.InvokeVoid(
		t,
		"resetMessagesPerSecond",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) ResetTotal() {
	_jsii_.InvokeVoid(
		t,
		"resetTotal",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfApp_LimitsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

