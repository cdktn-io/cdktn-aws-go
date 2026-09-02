package awsec2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsec2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsec2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference interface {
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
	EnableResourceNameDnsAaaaRecord() interface{}
	// Experimental.
	SetEnableResourceNameDnsAaaaRecord(val interface{})
	// Experimental.
	EnableResourceNameDnsAaaaRecordInput() interface{}
	// Experimental.
	EnableResourceNameDnsARecord() interface{}
	// Experimental.
	SetEnableResourceNameDnsARecord(val interface{})
	// Experimental.
	EnableResourceNameDnsARecordInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	HostnameType() *string
	// Experimental.
	SetHostnameType(val *string)
	// Experimental.
	HostnameTypeInput() *string
	// Experimental.
	InternalValue() *TfLaunchTemplate_PrivateDnsNameOptionsProperty
	// Experimental.
	SetInternalValue(val *TfLaunchTemplate_PrivateDnsNameOptionsProperty)
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
	ResetEnableResourceNameDnsAaaaRecord()
	// Experimental.
	ResetEnableResourceNameDnsARecord()
	// Experimental.
	ResetHostnameType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference
type jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) EnableResourceNameDnsAaaaRecord() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) EnableResourceNameDnsAaaaRecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsAaaaRecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) EnableResourceNameDnsARecord() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecord",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) EnableResourceNameDnsARecordInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableResourceNameDnsARecordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) HostnameType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) HostnameTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostnameTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) InternalValue() *TfLaunchTemplate_PrivateDnsNameOptionsProperty {
	var returns *TfLaunchTemplate_PrivateDnsNameOptionsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-ec2.TfLaunchTemplate.PrivateDnsNameOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference_Override(t TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-ec2.TfLaunchTemplate.PrivateDnsNameOptionsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference)SetEnableResourceNameDnsAaaaRecord(val interface{}) {
	if err := j.validateSetEnableResourceNameDnsAaaaRecordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsAaaaRecord",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference)SetEnableResourceNameDnsARecord(val interface{}) {
	if err := j.validateSetEnableResourceNameDnsARecordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableResourceNameDnsARecord",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference)SetHostnameType(val *string) {
	if err := j.validateSetHostnameTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostnameType",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference)SetInternalValue(val *TfLaunchTemplate_PrivateDnsNameOptionsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) ResetEnableResourceNameDnsAaaaRecord() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableResourceNameDnsAaaaRecord",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) ResetEnableResourceNameDnsARecord() {
	_jsii_.InvokeVoid(
		t,
		"resetEnableResourceNameDnsARecord",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) ResetHostnameType() {
	_jsii_.InvokeVoid(
		t,
		"resetHostnameType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfLaunchTemplate_PrivateDnsNameOptionsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

