package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference interface {
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
	InternalValue() *TfRoute_SpecGrpcRouteRetryPolicyProperty
	// Experimental.
	SetInternalValue(val *TfRoute_SpecGrpcRouteRetryPolicyProperty)
	// Experimental.
	MaxRetries() *float64
	// Experimental.
	SetMaxRetries(val *float64)
	// Experimental.
	MaxRetriesInput() *float64
	// Experimental.
	PerRetryTimeout() TfRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutPropertyOutputReference
	// Experimental.
	PerRetryTimeoutInput() *TfRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutProperty
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
	PutPerRetryTimeout(value *TfRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutProperty)
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

// The jsii proxy struct for TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference
type jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GrpcRetryEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grpcRetryEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GrpcRetryEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"grpcRetryEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) HttpRetryEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpRetryEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) HttpRetryEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpRetryEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) InternalValue() *TfRoute_SpecGrpcRouteRetryPolicyProperty {
	var returns *TfRoute_SpecGrpcRouteRetryPolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) PerRetryTimeout() TfRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutPropertyOutputReference {
	var returns TfRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutPropertyOutputReference
	_jsii_.Get(
		j,
		"perRetryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) PerRetryTimeoutInput() *TfRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutProperty {
	var returns *TfRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutProperty
	_jsii_.Get(
		j,
		"perRetryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) TcpRetryEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tcpRetryEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) TcpRetryEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tcpRetryEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfRoute.SpecGrpcRouteRetryPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference_Override(t TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfRoute.SpecGrpcRouteRetryPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetGrpcRetryEvents(val *[]*string) {
	if err := j.validateSetGrpcRetryEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"grpcRetryEvents",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetHttpRetryEvents(val *[]*string) {
	if err := j.validateSetHttpRetryEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpRetryEvents",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetInternalValue(val *TfRoute_SpecGrpcRouteRetryPolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetTcpRetryEvents(val *[]*string) {
	if err := j.validateSetTcpRetryEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpRetryEvents",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) PutPerRetryTimeout(value *TfRoute_SpecGrpcRouteRetryPolicyPerRetryTimeoutProperty) {
	if err := t.validatePutPerRetryTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPerRetryTimeout",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ResetGrpcRetryEvents() {
	_jsii_.InvokeVoid(
		t,
		"resetGrpcRetryEvents",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ResetHttpRetryEvents() {
	_jsii_.InvokeVoid(
		t,
		"resetHttpRetryEvents",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ResetTcpRetryEvents() {
	_jsii_.InvokeVoid(
		t,
		"resetTcpRetryEvents",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfRoute_SpecGrpcRouteRetryPolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

