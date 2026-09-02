package awsnetworkmanager

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkmanager/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfVpcAttachment_OptionsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApplianceModeSupport() interface{}
	// Experimental.
	SetApplianceModeSupport(val interface{})
	// Experimental.
	ApplianceModeSupportInput() interface{}
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
	DnsSupport() interface{}
	// Experimental.
	SetDnsSupport(val interface{})
	// Experimental.
	DnsSupportInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfVpcAttachment_OptionsProperty
	// Experimental.
	SetInternalValue(val *TfVpcAttachment_OptionsProperty)
	// Experimental.
	Ipv6Support() interface{}
	// Experimental.
	SetIpv6Support(val interface{})
	// Experimental.
	Ipv6SupportInput() interface{}
	// Experimental.
	SecurityGroupReferencingSupport() interface{}
	// Experimental.
	SetSecurityGroupReferencingSupport(val interface{})
	// Experimental.
	SecurityGroupReferencingSupportInput() interface{}
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
	ResetApplianceModeSupport()
	// Experimental.
	ResetDnsSupport()
	// Experimental.
	ResetIpv6Support()
	// Experimental.
	ResetSecurityGroupReferencingSupport()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfVpcAttachment_OptionsPropertyOutputReference
type jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) ApplianceModeSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applianceModeSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) ApplianceModeSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"applianceModeSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) DnsSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) DnsSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dnsSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) InternalValue() *TfVpcAttachment_OptionsProperty {
	var returns *TfVpcAttachment_OptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) Ipv6Support() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6Support",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) Ipv6SupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipv6SupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) SecurityGroupReferencingSupport() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityGroupReferencingSupport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) SecurityGroupReferencingSupportInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"securityGroupReferencingSupportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfVpcAttachment_OptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfVpcAttachment_OptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfVpcAttachment_OptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-manager.TfVpcAttachment.OptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfVpcAttachment_OptionsPropertyOutputReference_Override(t TfVpcAttachment_OptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-manager.TfVpcAttachment.OptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference)SetApplianceModeSupport(val interface{}) {
	if err := j.validateSetApplianceModeSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applianceModeSupport",
		val,
	)
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference)SetDnsSupport(val interface{}) {
	if err := j.validateSetDnsSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dnsSupport",
		val,
	)
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference)SetInternalValue(val *TfVpcAttachment_OptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference)SetIpv6Support(val interface{}) {
	if err := j.validateSetIpv6SupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipv6Support",
		val,
	)
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference)SetSecurityGroupReferencingSupport(val interface{}) {
	if err := j.validateSetSecurityGroupReferencingSupportParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityGroupReferencingSupport",
		val,
	)
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) ResetApplianceModeSupport() {
	_jsii_.InvokeVoid(
		t,
		"resetApplianceModeSupport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) ResetDnsSupport() {
	_jsii_.InvokeVoid(
		t,
		"resetDnsSupport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) ResetIpv6Support() {
	_jsii_.InvokeVoid(
		t,
		"resetIpv6Support",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) ResetSecurityGroupReferencingSupport() {
	_jsii_.InvokeVoid(
		t,
		"resetSecurityGroupReferencingSupport",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfVpcAttachment_OptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

