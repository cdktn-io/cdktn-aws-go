package awsappmesh

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsappmesh/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsappmesh/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference interface {
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
	Hostname() AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteHostnamePropertyOutputReference
	// Experimental.
	HostnameInput() *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteHostnameProperty
	// Experimental.
	InternalValue() *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteProperty
	// Experimental.
	SetInternalValue(val *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteProperty)
	// Experimental.
	Path() AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePathPropertyOutputReference
	// Experimental.
	PathInput() *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePathProperty
	// Experimental.
	Prefix() AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePrefixPropertyOutputReference
	// Experimental.
	PrefixInput() *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePrefixProperty
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
	PutHostname(value *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteHostnameProperty)
	// Experimental.
	PutPath(value *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePathProperty)
	// Experimental.
	PutPrefix(value *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePrefixProperty)
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

// The jsii proxy struct for AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference
type jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) Hostname() AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteHostnamePropertyOutputReference {
	var returns AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteHostnamePropertyOutputReference
	_jsii_.Get(
		j,
		"hostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) HostnameInput() *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteHostnameProperty {
	var returns *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteHostnameProperty
	_jsii_.Get(
		j,
		"hostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) InternalValue() *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteProperty {
	var returns *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) Path() AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePathPropertyOutputReference {
	var returns AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePathPropertyOutputReference
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) PathInput() *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePathProperty {
	var returns *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePathProperty
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) Prefix() AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePrefixPropertyOutputReference {
	var returns AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePrefixPropertyOutputReference
	_jsii_.Get(
		j,
		"prefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) PrefixInput() *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePrefixProperty {
	var returns *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePrefixProperty
	_jsii_.Get(
		j,
		"prefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsAppmeshGatewayRoute.SpecHttpRouteActionRewritePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference_Override(a AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-app-mesh.AwsAppmeshGatewayRoute.SpecHttpRouteActionRewritePropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference)SetInternalValue(val *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) PutHostname(value *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewriteHostnameProperty) {
	if err := a.validatePutHostnameParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putHostname",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) PutPath(value *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePathProperty) {
	if err := a.validatePutPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPath",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) PutPrefix(value *AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePrefixProperty) {
	if err := a.validatePutPrefixParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putPrefix",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ResetHostname() {
	_jsii_.InvokeVoid(
		a,
		"resetHostname",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ResetPath() {
	_jsii_.InvokeVoid(
		a,
		"resetPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ResetPrefix() {
	_jsii_.InvokeVoid(
		a,
		"resetPrefix",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsAppmeshGatewayRoute_SpecHttpRouteActionRewritePropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

