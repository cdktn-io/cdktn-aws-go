package awsvpclattice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsvpclattice/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsvpclattice/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference interface {
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
	InternalValue() *TfServiceNetworkVpcAssociation_DnsOptionsProperty
	// Experimental.
	SetInternalValue(val *TfServiceNetworkVpcAssociation_DnsOptionsProperty)
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

// The jsii proxy struct for TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference
type jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) InternalValue() *TfServiceNetworkVpcAssociation_DnsOptionsProperty {
	var returns *TfServiceNetworkVpcAssociation_DnsOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) PrivateDnsPreference() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsPreference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) PrivateDnsPreferenceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateDnsPreferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) PrivateDnsSpecifiedDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateDnsSpecifiedDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) PrivateDnsSpecifiedDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateDnsSpecifiedDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-vpc-lattice.TfServiceNetworkVpcAssociation.DnsOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference_Override(t TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-vpc-lattice.TfServiceNetworkVpcAssociation.DnsOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference)SetInternalValue(val *TfServiceNetworkVpcAssociation_DnsOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference)SetPrivateDnsPreference(val *string) {
	if err := j.validateSetPrivateDnsPreferenceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsPreference",
		val,
	)
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference)SetPrivateDnsSpecifiedDomains(val *[]*string) {
	if err := j.validateSetPrivateDnsSpecifiedDomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateDnsSpecifiedDomains",
		val,
	)
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) ResetPrivateDnsPreference() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivateDnsPreference",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) ResetPrivateDnsSpecifiedDomains() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivateDnsSpecifiedDomains",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfServiceNetworkVpcAssociation_DnsOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

