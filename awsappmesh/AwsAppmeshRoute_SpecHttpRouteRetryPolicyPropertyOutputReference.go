package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference interface {
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
	HttpRetryEvents() *[]*string
	// Experimental.
	SetHttpRetryEvents(val *[]*string)
	// Experimental.
	HttpRetryEventsInput() *[]*string
	// Experimental.
	InternalValue() *AwsAppmeshRoute_SpecHttpRouteRetryPolicyProperty
	// Experimental.
	SetInternalValue(val *AwsAppmeshRoute_SpecHttpRouteRetryPolicyProperty)
	// Experimental.
	MaxRetries() *float64
	// Experimental.
	SetMaxRetries(val *float64)
	// Experimental.
	MaxRetriesInput() *float64
	// Experimental.
	PerRetryTimeout() AwsAppmeshRoute_SpecHttpRouteRetryPolicyPerRetryTimeoutPropertyOutputReference
	// Experimental.
	PerRetryTimeoutInput() *AwsAppmeshRoute_SpecHttpRouteRetryPolicyPerRetryTimeoutProperty
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
	PutPerRetryTimeout(value *AwsAppmeshRoute_SpecHttpRouteRetryPolicyPerRetryTimeoutProperty)
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

// The jsii proxy struct for AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference
type jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) HttpRetryEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpRetryEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) HttpRetryEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpRetryEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) InternalValue() *AwsAppmeshRoute_SpecHttpRouteRetryPolicyProperty {
	var returns *AwsAppmeshRoute_SpecHttpRouteRetryPolicyProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) PerRetryTimeout() AwsAppmeshRoute_SpecHttpRouteRetryPolicyPerRetryTimeoutPropertyOutputReference {
	var returns AwsAppmeshRoute_SpecHttpRouteRetryPolicyPerRetryTimeoutPropertyOutputReference
	_jsii_.Get(
		j,
		"perRetryTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) PerRetryTimeoutInput() *AwsAppmeshRoute_SpecHttpRouteRetryPolicyPerRetryTimeoutProperty {
	var returns *AwsAppmeshRoute_SpecHttpRouteRetryPolicyPerRetryTimeoutProperty
	_jsii_.Get(
		j,
		"perRetryTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) TcpRetryEvents() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tcpRetryEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) TcpRetryEventsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tcpRetryEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsAppmeshRoute.SpecHttpRouteRetryPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference_Override(a AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsAppmeshRoute.SpecHttpRouteRetryPolicyPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference)SetHttpRetryEvents(val *[]*string) {
	if err := j.validateSetHttpRetryEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"httpRetryEvents",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference)SetInternalValue(val *AwsAppmeshRoute_SpecHttpRouteRetryPolicyProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference)SetMaxRetries(val *float64) {
	if err := j.validateSetMaxRetriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference)SetTcpRetryEvents(val *[]*string) {
	if err := j.validateSetTcpRetryEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tcpRetryEvents",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) PutPerRetryTimeout(value *AwsAppmeshRoute_SpecHttpRouteRetryPolicyPerRetryTimeoutProperty) {
	if err := a.validatePutPerRetryTimeoutParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPerRetryTimeout",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) ResetHttpRetryEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetHttpRetryEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) ResetTcpRetryEvents() {
	_jsii_.InvokeVoid(
		a,
		"resetTcpRetryEvents",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppmeshRoute_SpecHttpRouteRetryPolicyPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

