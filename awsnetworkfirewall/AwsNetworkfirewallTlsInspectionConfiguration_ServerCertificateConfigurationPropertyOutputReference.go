package awsnetworkfirewall

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsnetworkfirewall/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CertificateAuthorityArn() *string
	// Experimental.
	SetCertificateAuthorityArn(val *string)
	// Experimental.
	CertificateAuthorityArnInput() *string
	// Experimental.
	CheckCertificateRevocationStatus() AwsNetworkfirewallTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyList
	// Experimental.
	CheckCertificateRevocationStatusInput() interface{}
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
	Scope() AwsNetworkfirewallTlsInspectionConfiguration_ScopePropertyList
	// Experimental.
	ScopeInput() interface{}
	// Experimental.
	ServerCertificate() AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificatePropertyList
	// Experimental.
	ServerCertificateInput() interface{}
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
	PutCheckCertificateRevocationStatus(value interface{})
	// Experimental.
	PutScope(value interface{})
	// Experimental.
	PutServerCertificate(value interface{})
	// Experimental.
	ResetCertificateAuthorityArn()
	// Experimental.
	ResetCheckCertificateRevocationStatus()
	// Experimental.
	ResetScope()
	// Experimental.
	ResetServerCertificate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference
type jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) CertificateAuthorityArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) CertificateAuthorityArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateAuthorityArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) CheckCertificateRevocationStatus() AwsNetworkfirewallTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyList {
	var returns AwsNetworkfirewallTlsInspectionConfiguration_CheckCertificateRevocationStatusPropertyList
	_jsii_.Get(
		j,
		"checkCertificateRevocationStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) CheckCertificateRevocationStatusInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"checkCertificateRevocationStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) Scope() AwsNetworkfirewallTlsInspectionConfiguration_ScopePropertyList {
	var returns AwsNetworkfirewallTlsInspectionConfiguration_ScopePropertyList
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) ScopeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) ServerCertificate() AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificatePropertyList {
	var returns AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificatePropertyList
	_jsii_.Get(
		j,
		"serverCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) ServerCertificateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallTlsInspectionConfiguration.ServerCertificateConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference_Override(a AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-network-firewall.AwsNetworkfirewallTlsInspectionConfiguration.ServerCertificateConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference)SetCertificateAuthorityArn(val *string) {
	if err := j.validateSetCertificateAuthorityArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificateAuthorityArn",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) PutCheckCertificateRevocationStatus(value interface{}) {
	if err := a.validatePutCheckCertificateRevocationStatusParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putCheckCertificateRevocationStatus",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) PutScope(value interface{}) {
	if err := a.validatePutScopeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putScope",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) PutServerCertificate(value interface{}) {
	if err := a.validatePutServerCertificateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putServerCertificate",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) ResetCertificateAuthorityArn() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificateAuthorityArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) ResetCheckCertificateRevocationStatus() {
	_jsii_.InvokeVoid(
		a,
		"resetCheckCertificateRevocationStatus",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) ResetScope() {
	_jsii_.InvokeVoid(
		a,
		"resetScope",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) ResetServerCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetServerCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsNetworkfirewallTlsInspectionConfiguration_ServerCertificateConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

