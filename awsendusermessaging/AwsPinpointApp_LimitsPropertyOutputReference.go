package awsendusermessaging

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsendusermessaging/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsendusermessaging/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsPinpointApp_LimitsPropertyOutputReference interface {
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
	InternalValue() *AwsPinpointApp_LimitsProperty
	// Experimental.
	SetInternalValue(val *AwsPinpointApp_LimitsProperty)
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

// The jsii proxy struct for AwsPinpointApp_LimitsPropertyOutputReference
type jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) Daily() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"daily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) DailyInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) InternalValue() *AwsPinpointApp_LimitsProperty {
	var returns *AwsPinpointApp_LimitsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) MaximumDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) MaximumDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maximumDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) MessagesPerSecond() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messagesPerSecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) MessagesPerSecondInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"messagesPerSecondInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) Total() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"total",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) TotalInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"totalInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsPinpointApp_LimitsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsPinpointApp_LimitsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsPinpointApp_LimitsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-end-user-messaging.AwsPinpointApp.LimitsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsPinpointApp_LimitsPropertyOutputReference_Override(a AwsPinpointApp_LimitsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-end-user-messaging.AwsPinpointApp.LimitsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference)SetDaily(val *float64) {
	if err := j.validateSetDailyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daily",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference)SetInternalValue(val *AwsPinpointApp_LimitsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference)SetMaximumDuration(val *float64) {
	if err := j.validateSetMaximumDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maximumDuration",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference)SetMessagesPerSecond(val *float64) {
	if err := j.validateSetMessagesPerSecondParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"messagesPerSecond",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference)SetTotal(val *float64) {
	if err := j.validateSetTotalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"total",
		val,
	)
}

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) ResetDaily() {
	_jsii_.InvokeVoid(
		a,
		"resetDaily",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) ResetMaximumDuration() {
	_jsii_.InvokeVoid(
		a,
		"resetMaximumDuration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) ResetMessagesPerSecond() {
	_jsii_.InvokeVoid(
		a,
		"resetMessagesPerSecond",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) ResetTotal() {
	_jsii_.InvokeVoid(
		a,
		"resetTotal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsPinpointApp_LimitsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

