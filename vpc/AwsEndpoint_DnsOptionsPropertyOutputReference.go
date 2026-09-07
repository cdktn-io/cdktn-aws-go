package vpc

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/vpc/jsii"

	"github.com/cdktn-io/cdktn-aws-go/vpc/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEndpoint_DnsOptionsPropertyOutputReference interface {
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
	DnsRecordIpType() *string
	// Experimental.
	SetDnsRecordIpType(val *string)
	// Experimental.
	DnsRecordIpTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsEndpoint_DnsOptionsProperty
	// Experimental.
	SetInternalValue(val *AwsEndpoint_DnsOptionsProperty)
	// Experimental.
	PrivateDnsOnlyForInboundResolverEndpoint() interface{}
	// Experimental.
	SetPrivateDnsOnlyForInboundResolverEndpoint(val interface{})
	// Experimental.
	PrivateDnsOnlyForInboundResolverEndpointInput() interface{}
	// Experimental.
	PrivateDnsPreference() *string
	// Experimental.
	SetPrivateDnsPreference(val *string)
	// Experimental.
	PrivateDnsPreferenceInput() *string
	// Experimental.
	PrivateDnsSpecifiedDomains() *[]*string
	// Experimental.
	SetPrivateDnsSpecifiedDomains(val *[]*string)
	// Experimental.
	PrivateDnsSpecifiedDomainsInput() *[]*string
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
	ResetDnsRecordIpType()
	// Experimental.
	ResetPrivateDnsOnlyForInboundResolverEndpoint()
	// Experimental.
	ResetPrivateDnsPreference()
	// Experimental.
	ResetPrivateDnsSpecifiedDomains()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsEndpoint_DnsOptionsPropertyOutputReference
type jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) DnsRecordIpType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsRecordIpType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) DnsRecordIpTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dnsRecordIpTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) InternalValue() *AwsEndpoint_DnsOptionsProperty {
	var returns *AwsEndpoint_DnsOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) PrivateDnsOnlyForInboundResolverEndpoint() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateDnsOnlyForInboundResolverEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) PrivateDnsOnlyForInboundResolverEndpointInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateDnsOnlyForInboundResolverEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) PrivateDnsPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) PrivateDnsPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) PrivateDnsSpecifiedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateDnsSpecifiedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) PrivateDnsSpecifiedDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateDnsSpecifiedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEndpoint_DnsOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEndpoint_DnsOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEndpoint_DnsOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEndpoint.DnsOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEndpoint_DnsOptionsPropertyOutputReference_Override(a AwsEndpoint_DnsOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc.AwsEndpoint.DnsOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference)SetDnsRecordIpType(val *string) {
	if err := j.validateSetDnsRecordIpTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsRecordIpType",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference)SetInternalValue(val *AwsEndpoint_DnsOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference)SetPrivateDnsOnlyForInboundResolverEndpoint(val interface{}) {
	if err := j.validateSetPrivateDnsOnlyForInboundResolverEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsOnlyForInboundResolverEndpoint",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference)SetPrivateDnsPreference(val *string) {
	if err := j.validateSetPrivateDnsPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsPreference",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference)SetPrivateDnsSpecifiedDomains(val *[]*string) {
	if err := j.validateSetPrivateDnsSpecifiedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsSpecifiedDomains",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) ResetDnsRecordIpType() {
	_jsii_.InvokeVoid(
		a,
		"resetDnsRecordIpType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) ResetPrivateDnsOnlyForInboundResolverEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateDnsOnlyForInboundResolverEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) ResetPrivateDnsPreference() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateDnsPreference",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) ResetPrivateDnsSpecifiedDomains() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivateDnsSpecifiedDomains",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEndpoint_DnsOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

