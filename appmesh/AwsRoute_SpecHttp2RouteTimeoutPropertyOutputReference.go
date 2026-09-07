package appmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference interface {
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
	Fqn() *string
	// Experimental.
	Idle() AwsRoute_SpecHttp2RouteTimeoutIdlePropertyOutputReference
	// Experimental.
	IdleInput() *AwsRoute_SpecHttp2RouteTimeoutIdleProperty
	// Experimental.
	InternalValue() *AwsRoute_SpecHttp2RouteTimeoutProperty
	// Experimental.
	SetInternalValue(val *AwsRoute_SpecHttp2RouteTimeoutProperty)
	// Experimental.
	PerRequest() AwsRoute_SpecHttp2RouteTimeoutPerRequestPropertyOutputReference
	// Experimental.
	PerRequestInput() *AwsRoute_SpecHttp2RouteTimeoutPerRequestProperty
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
	PutIdle(value *AwsRoute_SpecHttp2RouteTimeoutIdleProperty)
	// Experimental.
	PutPerRequest(value *AwsRoute_SpecHttp2RouteTimeoutPerRequestProperty)
	// Experimental.
	ResetIdle()
	// Experimental.
	ResetPerRequest()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference
type jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) Idle() AwsRoute_SpecHttp2RouteTimeoutIdlePropertyOutputReference {
	var returns AwsRoute_SpecHttp2RouteTimeoutIdlePropertyOutputReference
	_jsii_.Get(
		j,
		"idle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) IdleInput() *AwsRoute_SpecHttp2RouteTimeoutIdleProperty {
	var returns *AwsRoute_SpecHttp2RouteTimeoutIdleProperty
	_jsii_.Get(
		j,
		"idleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) InternalValue() *AwsRoute_SpecHttp2RouteTimeoutProperty {
	var returns *AwsRoute_SpecHttp2RouteTimeoutProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) PerRequest() AwsRoute_SpecHttp2RouteTimeoutPerRequestPropertyOutputReference {
	var returns AwsRoute_SpecHttp2RouteTimeoutPerRequestPropertyOutputReference
	_jsii_.Get(
		j,
		"perRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) PerRequestInput() *AwsRoute_SpecHttp2RouteTimeoutPerRequestProperty {
	var returns *AwsRoute_SpecHttp2RouteTimeoutPerRequestProperty
	_jsii_.Get(
		j,
		"perRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRoute_SpecHttp2RouteTimeoutPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsRoute.SpecHttp2RouteTimeoutPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference_Override(a AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsRoute.SpecHttp2RouteTimeoutPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference)SetInternalValue(val *AwsRoute_SpecHttp2RouteTimeoutProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) PutIdle(value *AwsRoute_SpecHttp2RouteTimeoutIdleProperty) {
	if err := a.validatePutIdleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putIdle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) PutPerRequest(value *AwsRoute_SpecHttp2RouteTimeoutPerRequestProperty) {
	if err := a.validatePutPerRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPerRequest",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) ResetIdle() {
	_jsii_.InvokeVoid(
		a,
		"resetIdle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) ResetPerRequest() {
	_jsii_.InvokeVoid(
		a,
		"resetPerRequest",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRoute_SpecHttp2RouteTimeoutPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

