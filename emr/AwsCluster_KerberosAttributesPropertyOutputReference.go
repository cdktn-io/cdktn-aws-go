package emr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/emr/jsii"

	"github.com/cdktn-io/cdktn-aws-go/emr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsCluster_KerberosAttributesPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AdDomainJoinPassword() *string
	// Experimental.
	SetAdDomainJoinPassword(val *string)
	// Experimental.
	AdDomainJoinPasswordInput() *string
	// Experimental.
	AdDomainJoinUser() *string
	// Experimental.
	SetAdDomainJoinUser(val *string)
	// Experimental.
	AdDomainJoinUserInput() *string
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
	CrossRealmTrustPrincipalPassword() *string
	// Experimental.
	SetCrossRealmTrustPrincipalPassword(val *string)
	// Experimental.
	CrossRealmTrustPrincipalPasswordInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsCluster_KerberosAttributesProperty
	// Experimental.
	SetInternalValue(val *AwsCluster_KerberosAttributesProperty)
	// Experimental.
	KdcAdminPassword() *string
	// Experimental.
	SetKdcAdminPassword(val *string)
	// Experimental.
	KdcAdminPasswordInput() *string
	// Experimental.
	Realm() *string
	// Experimental.
	SetRealm(val *string)
	// Experimental.
	RealmInput() *string
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
	ResetAdDomainJoinPassword()
	// Experimental.
	ResetAdDomainJoinUser()
	// Experimental.
	ResetCrossRealmTrustPrincipalPassword()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsCluster_KerberosAttributesPropertyOutputReference
type jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) AdDomainJoinPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) AdDomainJoinPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) AdDomainJoinUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) AdDomainJoinUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) CrossRealmTrustPrincipalPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustPrincipalPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) CrossRealmTrustPrincipalPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustPrincipalPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) InternalValue() *AwsCluster_KerberosAttributesProperty {
	var returns *AwsCluster_KerberosAttributesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) KdcAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) KdcAdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcAdminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) RealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsCluster_KerberosAttributesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsCluster_KerberosAttributesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsCluster_KerberosAttributesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr.AwsCluster.KerberosAttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsCluster_KerberosAttributesPropertyOutputReference_Override(a AwsCluster_KerberosAttributesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.AwsCluster.KerberosAttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference)SetAdDomainJoinPassword(val *string) {
	if err := j.validateSetAdDomainJoinPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adDomainJoinPassword",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference)SetAdDomainJoinUser(val *string) {
	if err := j.validateSetAdDomainJoinUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adDomainJoinUser",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference)SetCrossRealmTrustPrincipalPassword(val *string) {
	if err := j.validateSetCrossRealmTrustPrincipalPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustPrincipalPassword",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference)SetInternalValue(val *AwsCluster_KerberosAttributesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference)SetKdcAdminPassword(val *string) {
	if err := j.validateSetKdcAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kdcAdminPassword",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference)SetRealm(val *string) {
	if err := j.validateSetRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) ResetAdDomainJoinPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetAdDomainJoinPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) ResetAdDomainJoinUser() {
	_jsii_.InvokeVoid(
		a,
		"resetAdDomainJoinUser",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) ResetCrossRealmTrustPrincipalPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetCrossRealmTrustPrincipalPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsCluster_KerberosAttributesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

