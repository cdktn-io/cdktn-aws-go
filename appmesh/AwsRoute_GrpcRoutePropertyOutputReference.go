package appmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRoute_GrpcRoutePropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	Action() AwsRoute_SpecGrpcRouteActionPropertyOutputReference
	// Experimental.
	ActionInput() *AwsRoute_SpecGrpcRouteActionProperty
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
	InternalValue() *AwsRoute_GrpcRouteProperty
	// Experimental.
	SetInternalValue(val *AwsRoute_GrpcRouteProperty)
	// Experimental.
	Match() AwsRoute_SpecGrpcRouteMatchPropertyOutputReference
	// Experimental.
	MatchInput() *AwsRoute_SpecGrpcRouteMatchProperty
	// Experimental.
	RetryPolicy() AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference
	// Experimental.
	RetryPolicyInput() *AwsRoute_SpecGrpcRouteRetryPolicyProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Timeout() AwsRoute_SpecGrpcRouteTimeoutPropertyOutputReference
	// Experimental.
	TimeoutInput() *AwsRoute_SpecGrpcRouteTimeoutProperty
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
	PutAction(value *AwsRoute_SpecGrpcRouteActionProperty)
	// Experimental.
	PutMatch(value *AwsRoute_SpecGrpcRouteMatchProperty)
	// Experimental.
	PutRetryPolicy(value *AwsRoute_SpecGrpcRouteRetryPolicyProperty)
	// Experimental.
	PutTimeout(value *AwsRoute_SpecGrpcRouteTimeoutProperty)
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

// The jsii proxy struct for AwsRoute_GrpcRoutePropertyOutputReference
type jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) Action() AwsRoute_SpecGrpcRouteActionPropertyOutputReference {
	var returns AwsRoute_SpecGrpcRouteActionPropertyOutputReference
	_jsii_.Get(
		j,
		"action",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) ActionInput() *AwsRoute_SpecGrpcRouteActionProperty {
	var returns *AwsRoute_SpecGrpcRouteActionProperty
	_jsii_.Get(
		j,
		"actionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) InternalValue() *AwsRoute_GrpcRouteProperty {
	var returns *AwsRoute_GrpcRouteProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) Match() AwsRoute_SpecGrpcRouteMatchPropertyOutputReference {
	var returns AwsRoute_SpecGrpcRouteMatchPropertyOutputReference
	_jsii_.Get(
		j,
		"match",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) MatchInput() *AwsRoute_SpecGrpcRouteMatchProperty {
	var returns *AwsRoute_SpecGrpcRouteMatchProperty
	_jsii_.Get(
		j,
		"matchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) RetryPolicy() AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference {
	var returns AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference
	_jsii_.Get(
		j,
		"retryPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) RetryPolicyInput() *AwsRoute_SpecGrpcRouteRetryPolicyProperty {
	var returns *AwsRoute_SpecGrpcRouteRetryPolicyProperty
	_jsii_.Get(
		j,
		"retryPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) Timeout() AwsRoute_SpecGrpcRouteTimeoutPropertyOutputReference {
	var returns AwsRoute_SpecGrpcRouteTimeoutPropertyOutputReference
	_jsii_.Get(
		j,
		"timeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) TimeoutInput() *AwsRoute_SpecGrpcRouteTimeoutProperty {
	var returns *AwsRoute_SpecGrpcRouteTimeoutProperty
	_jsii_.Get(
		j,
		"timeoutInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRoute_GrpcRoutePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRoute_GrpcRoutePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRoute_GrpcRoutePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsRoute.GrpcRoutePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRoute_GrpcRoutePropertyOutputReference_Override(a AwsRoute_GrpcRoutePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsRoute.GrpcRoutePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference)SetInternalValue(val *AwsRoute_GrpcRouteProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) PutAction(value *AwsRoute_SpecGrpcRouteActionProperty) {
	if err := a.validatePutActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAction",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) PutMatch(value *AwsRoute_SpecGrpcRouteMatchProperty) {
	if err := a.validatePutMatchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMatch",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) PutRetryPolicy(value *AwsRoute_SpecGrpcRouteRetryPolicyProperty) {
	if err := a.validatePutRetryPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetryPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) PutTimeout(value *AwsRoute_SpecGrpcRouteTimeoutProperty) {
	if err := a.validatePutTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTimeout",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) ResetMatch() {
	_jsii_.InvokeVoid(
		a,
		"resetMatch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) ResetRetryPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetRetryPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) ResetTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRoute_GrpcRoutePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

