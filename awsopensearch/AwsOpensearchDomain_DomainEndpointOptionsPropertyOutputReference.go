package awsopensearch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsopensearch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsopensearch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference interface {
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
	CustomEndpoint() *string
	// Experimental.
	SetCustomEndpoint(val *string)
	// Experimental.
	CustomEndpointCertificateArn() *string
	// Experimental.
	SetCustomEndpointCertificateArn(val *string)
	// Experimental.
	CustomEndpointCertificateArnInput() *string
	// Experimental.
	CustomEndpointEnabled() interface{}
	// Experimental.
	SetCustomEndpointEnabled(val interface{})
	// Experimental.
	CustomEndpointEnabledInput() interface{}
	// Experimental.
	CustomEndpointInput() *string
	// Experimental.
	EnforceHttps() interface{}
	// Experimental.
	SetEnforceHttps(val interface{})
	// Experimental.
	EnforceHttpsInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsOpensearchDomain_DomainEndpointOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsOpensearchDomain_DomainEndpointOptionsProperty)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TlsSecurityPolicy() *string
	// Experimental.
	SetTlsSecurityPolicy(val *string)
	// Experimental.
	TlsSecurityPolicyInput() *string
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
	ResetCustomEndpoint()
	// Experimental.
	ResetCustomEndpointCertificateArn()
	// Experimental.
	ResetCustomEndpointEnabled()
	// Experimental.
	ResetEnforceHttps()
	// Experimental.
	ResetTlsSecurityPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference
type jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) CustomEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) CustomEndpointCertificateArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEndpointCertificateArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) CustomEndpointCertificateArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEndpointCertificateArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) CustomEndpointEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customEndpointEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) CustomEndpointEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customEndpointEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) CustomEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) EnforceHttps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceHttps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) EnforceHttpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceHttpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) InternalValue() *AwsOpensearchDomain_DomainEndpointOptionsProperty {
	var returns *AwsOpensearchDomain_DomainEndpointOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) TlsSecurityPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsSecurityPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) TlsSecurityPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsSecurityPolicyInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-opensearch.AwsOpensearchDomain.DomainEndpointOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference_Override(a AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-opensearch.AwsOpensearchDomain.DomainEndpointOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference)SetCustomEndpoint(val *string) {
	if err := j.validateSetCustomEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference)SetCustomEndpointCertificateArn(val *string) {
	if err := j.validateSetCustomEndpointCertificateArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customEndpointCertificateArn",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference)SetCustomEndpointEnabled(val interface{}) {
	if err := j.validateSetCustomEndpointEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customEndpointEnabled",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference)SetEnforceHttps(val interface{}) {
	if err := j.validateSetEnforceHttpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceHttps",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference)SetInternalValue(val *AwsOpensearchDomain_DomainEndpointOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference)SetTlsSecurityPolicy(val *string) {
	if err := j.validateSetTlsSecurityPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsSecurityPolicy",
		val,
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) ResetCustomEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) ResetCustomEndpointCertificateArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomEndpointCertificateArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) ResetCustomEndpointEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetCustomEndpointEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) ResetEnforceHttps() {
	_jsii_.InvokeVoid(
		a,
		"resetEnforceHttps",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) ResetTlsSecurityPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsSecurityPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsOpensearchDomain_DomainEndpointOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

