package appmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/appmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/appmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference interface {
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
	Hostname() AwsGatewayRoute_SpecHttp2RouteActionRewriteHostnamePropertyOutputReference
	// Experimental.
	HostnameInput() *AwsGatewayRoute_SpecHttp2RouteActionRewriteHostnameProperty
	// Experimental.
	InternalValue() *AwsGatewayRoute_SpecHttp2RouteActionRewriteProperty
	// Experimental.
	SetInternalValue(val *AwsGatewayRoute_SpecHttp2RouteActionRewriteProperty)
	// Experimental.
	Path() AwsGatewayRoute_SpecHttp2RouteActionRewritePathPropertyOutputReference
	// Experimental.
	PathInput() *AwsGatewayRoute_SpecHttp2RouteActionRewritePathProperty
	// Experimental.
	Prefix() AwsGatewayRoute_SpecHttp2RouteActionRewritePrefixPropertyOutputReference
	// Experimental.
	PrefixInput() *AwsGatewayRoute_SpecHttp2RouteActionRewritePrefixProperty
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
	PutHostname(value *AwsGatewayRoute_SpecHttp2RouteActionRewriteHostnameProperty)
	// Experimental.
	PutPath(value *AwsGatewayRoute_SpecHttp2RouteActionRewritePathProperty)
	// Experimental.
	PutPrefix(value *AwsGatewayRoute_SpecHttp2RouteActionRewritePrefixProperty)
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

// The jsii proxy struct for AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference
type jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) Hostname() AwsGatewayRoute_SpecHttp2RouteActionRewriteHostnamePropertyOutputReference {
	var returns AwsGatewayRoute_SpecHttp2RouteActionRewriteHostnamePropertyOutputReference
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) HostnameInput() *AwsGatewayRoute_SpecHttp2RouteActionRewriteHostnameProperty {
	var returns *AwsGatewayRoute_SpecHttp2RouteActionRewriteHostnameProperty
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) InternalValue() *AwsGatewayRoute_SpecHttp2RouteActionRewriteProperty {
	var returns *AwsGatewayRoute_SpecHttp2RouteActionRewriteProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) Path() AwsGatewayRoute_SpecHttp2RouteActionRewritePathPropertyOutputReference {
	var returns AwsGatewayRoute_SpecHttp2RouteActionRewritePathPropertyOutputReference
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) PathInput() *AwsGatewayRoute_SpecHttp2RouteActionRewritePathProperty {
	var returns *AwsGatewayRoute_SpecHttp2RouteActionRewritePathProperty
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) Prefix() AwsGatewayRoute_SpecHttp2RouteActionRewritePrefixPropertyOutputReference {
	var returns AwsGatewayRoute_SpecHttp2RouteActionRewritePrefixPropertyOutputReference
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) PrefixInput() *AwsGatewayRoute_SpecHttp2RouteActionRewritePrefixProperty {
	var returns *AwsGatewayRoute_SpecHttp2RouteActionRewritePrefixProperty
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsGatewayRoute.SpecHttp2RouteActionRewritePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference_Override(a AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsGatewayRoute.SpecHttp2RouteActionRewritePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference)SetInternalValue(val *AwsGatewayRoute_SpecHttp2RouteActionRewriteProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) PutHostname(value *AwsGatewayRoute_SpecHttp2RouteActionRewriteHostnameProperty) {
	if err := a.validatePutHostnameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHostname",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) PutPath(value *AwsGatewayRoute_SpecHttp2RouteActionRewritePathProperty) {
	if err := a.validatePutPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPath",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) PutPrefix(value *AwsGatewayRoute_SpecHttp2RouteActionRewritePrefixProperty) {
	if err := a.validatePutPrefixParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrefix",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		a,
		"resetHostname",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsGatewayRoute_SpecHttp2RouteActionRewritePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

