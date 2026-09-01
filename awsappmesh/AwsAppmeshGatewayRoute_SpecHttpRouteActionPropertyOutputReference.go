package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference interface {
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
	InternalValue() *AwsAppmeshGatewayRoute_SpecHttpRouteActionProperty
	// Experimental.
	SetInternalValue(val *AwsAppmeshGatewayRoute_SpecHttpRouteActionProperty)
	// Experimental.
	Rewrite() AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference
	// Experimental.
	RewriteInput() *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteProperty
	// Experimental.
	Target() AwsAppmeshGatewayRoute_SpecHttpRouteActionTargetPropertyOutputReference
	// Experimental.
	TargetInput() *AwsAppmeshGatewayRoute_SpecHttpRouteActionTargetProperty
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
	PutRewrite(value *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteProperty)
	// Experimental.
	PutTarget(value *AwsAppmeshGatewayRoute_SpecHttpRouteActionTargetProperty)
	// Experimental.
	ResetRewrite()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference
type jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) InternalValue() *AwsAppmeshGatewayRoute_SpecHttpRouteActionProperty {
	var returns *AwsAppmeshGatewayRoute_SpecHttpRouteActionProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) Rewrite() AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference {
	var returns AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference
	_jsii_.Get(
		j,
		"rewrite",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) RewriteInput() *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteProperty {
	var returns *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteProperty
	_jsii_.Get(
		j,
		"rewriteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) Target() AwsAppmeshGatewayRoute_SpecHttpRouteActionTargetPropertyOutputReference {
	var returns AwsAppmeshGatewayRoute_SpecHttpRouteActionTargetPropertyOutputReference
	_jsii_.Get(
		j,
		"target",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) TargetInput() *AwsAppmeshGatewayRoute_SpecHttpRouteActionTargetProperty {
	var returns *AwsAppmeshGatewayRoute_SpecHttpRouteActionTargetProperty
	_jsii_.Get(
		j,
		"targetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsAppmeshGatewayRoute.SpecHttpRouteActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference_Override(a AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsAppmeshGatewayRoute.SpecHttpRouteActionPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference)SetInternalValue(val *AwsAppmeshGatewayRoute_SpecHttpRouteActionProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) PutRewrite(value *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteProperty) {
	if err := a.validatePutRewriteParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRewrite",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) PutTarget(value *AwsAppmeshGatewayRoute_SpecHttpRouteActionTargetProperty) {
	if err := a.validatePutTargetParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTarget",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) ResetRewrite() {
	_jsii_.InvokeVoid(
		a,
		"resetRewrite",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

