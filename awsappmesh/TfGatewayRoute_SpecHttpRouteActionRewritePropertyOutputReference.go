package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference interface {
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
	Hostname() TfGatewayRoute_SpecHttpRouteActionRewriteHostnamePropertyOutputReference
	// Experimental.
	HostnameInput() *TfGatewayRoute_SpecHttpRouteActionRewriteHostnameProperty
	// Experimental.
	InternalValue() *TfGatewayRoute_SpecHttpRouteActionRewriteProperty
	// Experimental.
	SetInternalValue(val *TfGatewayRoute_SpecHttpRouteActionRewriteProperty)
	// Experimental.
	Path() TfGatewayRoute_SpecHttpRouteActionRewritePathPropertyOutputReference
	// Experimental.
	PathInput() *TfGatewayRoute_SpecHttpRouteActionRewritePathProperty
	// Experimental.
	Prefix() TfGatewayRoute_SpecHttpRouteActionRewritePrefixPropertyOutputReference
	// Experimental.
	PrefixInput() *TfGatewayRoute_SpecHttpRouteActionRewritePrefixProperty
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
	PutHostname(value *TfGatewayRoute_SpecHttpRouteActionRewriteHostnameProperty)
	// Experimental.
	PutPath(value *TfGatewayRoute_SpecHttpRouteActionRewritePathProperty)
	// Experimental.
	PutPrefix(value *TfGatewayRoute_SpecHttpRouteActionRewritePrefixProperty)
	// Experimental.
	ResetHostname()
	// Experimental.
	ResetPath()
	// Experimental.
	ResetPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference
type jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) Hostname() TfGatewayRoute_SpecHttpRouteActionRewriteHostnamePropertyOutputReference {
	var returns TfGatewayRoute_SpecHttpRouteActionRewriteHostnamePropertyOutputReference
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) HostnameInput() *TfGatewayRoute_SpecHttpRouteActionRewriteHostnameProperty {
	var returns *TfGatewayRoute_SpecHttpRouteActionRewriteHostnameProperty
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) InternalValue() *TfGatewayRoute_SpecHttpRouteActionRewriteProperty {
	var returns *TfGatewayRoute_SpecHttpRouteActionRewriteProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) Path() TfGatewayRoute_SpecHttpRouteActionRewritePathPropertyOutputReference {
	var returns TfGatewayRoute_SpecHttpRouteActionRewritePathPropertyOutputReference
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) PathInput() *TfGatewayRoute_SpecHttpRouteActionRewritePathProperty {
	var returns *TfGatewayRoute_SpecHttpRouteActionRewritePathProperty
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) Prefix() TfGatewayRoute_SpecHttpRouteActionRewritePrefixPropertyOutputReference {
	var returns TfGatewayRoute_SpecHttpRouteActionRewritePrefixPropertyOutputReference
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) PrefixInput() *TfGatewayRoute_SpecHttpRouteActionRewritePrefixProperty {
	var returns *TfGatewayRoute_SpecHttpRouteActionRewritePrefixProperty
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfGatewayRoute.SpecHttpRouteActionRewritePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference_Override(t TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.TfGatewayRoute.SpecHttpRouteActionRewritePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference)SetInternalValue(val *TfGatewayRoute_SpecHttpRouteActionRewriteProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) PutHostname(value *TfGatewayRoute_SpecHttpRouteActionRewriteHostnameProperty) {
	if err := t.validatePutHostnameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putHostname",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) PutPath(value *TfGatewayRoute_SpecHttpRouteActionRewritePathProperty) {
	if err := t.validatePutPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPath",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) PutPrefix(value *TfGatewayRoute_SpecHttpRouteActionRewritePrefixProperty) {
	if err := t.validatePutPrefixParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putPrefix",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		t,
		"resetHostname",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		t,
		"resetPath",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		t,
		"resetPrefix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

