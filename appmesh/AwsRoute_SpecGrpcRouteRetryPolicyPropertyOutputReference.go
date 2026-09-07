package appmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference interface {
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
	GrpcRetryEvents() *[]*string
	// Experimental.
	SetGrpcRetryEvents(val *[]*string)
	// Experimental.
	GrpcRetryEventsInput() *[]*string
	// Experimental.
	HttpRetryEvents() *[]*string
	// Experimental.
	SetHttpRetryEvents(val *[]*string)
	// Experimental.
	HttpRetryEventsInput() *[]*string
	// Experimental.
	InternalValue() *AwsRoute_SpecGrpcRouteRetryPolicyProperty
	// Experimental.
	SetInternalValue(val *AwsRoute_SpecGrpcRouteRetryPolicyProperty)
	// Experimental.
	MaxRetries() *float64
	// Experimental.
	SetMaxRetries(val *float64)
	// Experimental.
	MaxRetriesInput() *float64
	// Experimental.
	PerRetryTimeout() AwsRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutPropertyOutputReference
	// Experimental.
	PerRetryTimeoutInput() *AwsRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutProperty
	// Experimental.
	TcpRetryEvents() *[]*string
	// Experimental.
	SetTcpRetryEvents(val *[]*string)
	// Experimental.
	TcpRetryEventsInput() *[]*string
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
	PutPerRetryTimeout(value *AwsRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutProperty)
	// Experimental.
	ResetGrpcRetryEvents()
	// Experimental.
	ResetHttpRetryEvents()
	// Experimental.
	ResetTcpRetryEvents()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference
type jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GrpcRetryEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grpcRetryEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GrpcRetryEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grpcRetryEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) HttpRetryEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpRetryEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) HttpRetryEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpRetryEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) InternalValue() *AwsRoute_SpecGrpcRouteRetryPolicyProperty {
	var returns *AwsRoute_SpecGrpcRouteRetryPolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) PerRetryTimeout() AwsRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutPropertyOutputReference {
	var returns AwsRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutPropertyOutputReference
	_jsii_.Get(
		j,
		"perRetryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) PerRetryTimeoutInput() *AwsRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutProperty {
	var returns *AwsRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutProperty
	_jsii_.Get(
		j,
		"perRetryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) TcpRetryEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tcpRetryEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) TcpRetryEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tcpRetryEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsRoute.SpecGrpcRouteRetryPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference_Override(a AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsRoute.SpecGrpcRouteRetryPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetGrpcRetryEvents(val *[]*string) {
	if err := j.validateSetGrpcRetryEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grpcRetryEvents",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetHttpRetryEvents(val *[]*string) {
	if err := j.validateSetHttpRetryEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpRetryEvents",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetInternalValue(val *AwsRoute_SpecGrpcRouteRetryPolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetTcpRetryEvents(val *[]*string) {
	if err := j.validateSetTcpRetryEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpRetryEvents",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) PutPerRetryTimeout(value *AwsRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutProperty) {
	if err := a.validatePutPerRetryTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPerRetryTimeout",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ResetGrpcRetryEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetGrpcRetryEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ResetHttpRetryEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpRetryEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ResetTcpRetryEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetTcpRetryEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

