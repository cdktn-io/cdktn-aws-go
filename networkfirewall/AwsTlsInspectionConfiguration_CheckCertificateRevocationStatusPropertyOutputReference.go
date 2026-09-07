package networkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/networkfirewall/jsii"

	"github.com/cdktn-io/cdktn-aws-go/networkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference interface {
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
	InternalValue() interface{}
	// Experimental.
	SetInternalValue(val interface{})
	// Experimental.
	RevokedStatusAction() *string
	// Experimental.
	SetRevokedStatusAction(val *string)
	// Experimental.
	RevokedStatusActionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UnknownStatusAction() *string
	// Experimental.
	SetUnknownStatusAction(val *string)
	// Experimental.
	UnknownStatusActionInput() *string
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
	ResetRevokedStatusAction()
	// Experimental.
	ResetUnknownStatusAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference
type jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) RevokedStatusAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revokedStatusAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) RevokedStatusActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revokedStatusActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) UnknownStatusAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unknownStatusAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) UnknownStatusActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unknownStatusActionInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsTlsInspectionConfiguration.CheckCertificateRevocationStatusPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference_Override(a AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsTlsInspectionConfiguration.CheckCertificateRevocationStatusPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference)SetRevokedStatusAction(val *string) {
	if err := j.validateSetRevokedStatusActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revokedStatusAction",
		val,
	)
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference)SetUnknownStatusAction(val *string) {
	if err := j.validateSetUnknownStatusActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unknownStatusAction",
		val,
	)
}

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) ResetRevokedStatusAction() {
	_jsii_.InvokeVoid(
		a,
		"resetRevokedStatusAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) ResetUnknownStatusAction() {
	_jsii_.InvokeVoid(
		a,
		"resetUnknownStatusAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

