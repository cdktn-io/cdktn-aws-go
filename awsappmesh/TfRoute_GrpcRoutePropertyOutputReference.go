package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRoute_GrpcRoutePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Action() TfRoute_SpecGrpcRouteActionPropertyOutputReference
	// Experimental.
	ActionInput() *TfRoute_SpecGrpcRouteActionProperty
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
	InternalValue() *TfRoute_GrpcRouteProperty
	// Experimental.
	SetInternalValue(val *TfRoute_GrpcRouteProperty)
	// Experimental.
	Match() TfRoute_SpecGrpcRouteMatchPropertyOutputReference
	// Experimental.
	MatchInput() *TfRoute_SpecGrpcRouteMatchProperty
	// Experimental.
	RetryPolicy() TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference
	// Experimental.
	RetryPolicyInput() *TfRoute_SpecGrpcRouteRetryPolicyProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Timeout() TfRoute_SpecGrpcRouteTimeoutPropertyOutputReference
	// Experimental.
	TimeoutInput() *TfRoute_SpecGrpcRouteTimeoutProperty
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
	PutAction(value *TfRoute_SpecGrpcRouteActionProperty)
	// Experimental.
	PutMatch(value *TfRoute_SpecGrpcRouteMatchProperty)
	// Experimental.
	PutRetryPolicy(value *TfRoute_SpecGrpcRouteRetryPolicyProperty)
	// Experimental.
	PutTimeout(value *TfRoute_SpecGrpcRouteTimeoutProperty)
	// Experimental.
	ResetMatch()
	// Experimental.
	ResetRetryPolicy()
	// Experimental.
	ResetTimeout()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfRoute_GrpcRoutePropertyOutputReference
type jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) Action() TfRoute_SpecGrpcRouteActionPropertyOutputReference {
	var returns TfRoute_SpecGrpcRouteActionPropertyOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) ActionInput() *TfRoute_SpecGrpcRouteActionProperty {
	var returns *TfRoute_SpecGrpcRouteActionProperty
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) InternalValue() *TfRoute_GrpcRouteProperty {
	var returns *TfRoute_GrpcRouteProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) Match() TfRoute_SpecGrpcRouteMatchPropertyOutputReference {
	var returns TfRoute_SpecGrpcRouteMatchPropertyOutputReference
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) MatchInput() *TfRoute_SpecGrpcRouteMatchProperty {
	var returns *TfRoute_SpecGrpcRouteMatchProperty
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) RetryPolicy() TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference {
	var returns TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) RetryPolicyInput() *TfRoute_SpecGrpcRouteRetryPolicyProperty {
	var returns *TfRoute_SpecGrpcRouteRetryPolicyProperty
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) Timeout() TfRoute_SpecGrpcRouteTimeoutPropertyOutputReference {
	var returns TfRoute_SpecGrpcRouteTimeoutPropertyOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) TimeoutInput() *TfRoute_SpecGrpcRouteTimeoutProperty {
	var returns *TfRoute_SpecGrpcRouteTimeoutProperty
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRoute_GrpcRoutePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfRoute_GrpcRoutePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRoute_GrpcRoutePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfRoute.GrpcRoutePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRoute_GrpcRoutePropertyOutputReference_Override(t TfRoute_GrpcRoutePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfRoute.GrpcRoutePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference)SetInternalValue(val *TfRoute_GrpcRouteProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) PutAction(value *TfRoute_SpecGrpcRouteActionProperty) {
	if err := t.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAction",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) PutMatch(value *TfRoute_SpecGrpcRouteMatchProperty) {
	if err := t.validatePutMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putMatch",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) PutRetryPolicy(value *TfRoute_SpecGrpcRouteRetryPolicyProperty) {
	if err := t.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) PutTimeout(value *TfRoute_SpecGrpcRouteTimeoutProperty) {
	if err := t.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putTimeout",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) ResetMatch() {
	_jsii_.InvokeVoid(
		t,
		"resetMatch",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		t,
		"resetTimeout",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRoute_GrpcRoutePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

