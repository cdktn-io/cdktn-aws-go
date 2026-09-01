package awsemr

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsemr/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsemr/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsEmrCluster_KerberosAttributesPropertyOutputReference interface {
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
	InternalValue() *AwsEmrCluster_KerberosAttributesProperty
	// Experimental.
	SetInternalValue(val *AwsEmrCluster_KerberosAttributesProperty)
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

// The jsii proxy struct for AwsEmrCluster_KerberosAttributesPropertyOutputReference
type jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) AdDomainJoinPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) AdDomainJoinPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) AdDomainJoinUser() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) AdDomainJoinUserInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adDomainJoinUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) CrossRealmTrustPrincipalPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustPrincipalPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) CrossRealmTrustPrincipalPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"crossRealmTrustPrincipalPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) InternalValue() *AwsEmrCluster_KerberosAttributesProperty {
	var returns *AwsEmrCluster_KerberosAttributesProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) KdcAdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcAdminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) KdcAdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kdcAdminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) Realm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) RealmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"realmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsEmrCluster_KerberosAttributesPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsEmrCluster_KerberosAttributesPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsEmrCluster_KerberosAttributesPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrCluster.KerberosAttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsEmrCluster_KerberosAttributesPropertyOutputReference_Override(a AwsEmrCluster_KerberosAttributesPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-emr.AwsEmrCluster.KerberosAttributesPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference)SetAdDomainJoinPassword(val *string) {
	if err := j.validateSetAdDomainJoinPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adDomainJoinPassword",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference)SetAdDomainJoinUser(val *string) {
	if err := j.validateSetAdDomainJoinUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adDomainJoinUser",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference)SetCrossRealmTrustPrincipalPassword(val *string) {
	if err := j.validateSetCrossRealmTrustPrincipalPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossRealmTrustPrincipalPassword",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference)SetInternalValue(val *AwsEmrCluster_KerberosAttributesProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference)SetKdcAdminPassword(val *string) {
	if err := j.validateSetKdcAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kdcAdminPassword",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference)SetRealm(val *string) {
	if err := j.validateSetRealmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"realm",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) ResetAdDomainJoinPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetAdDomainJoinPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) ResetAdDomainJoinUser() {
	_jsii_.InvokeVoid(
		a,
		"resetAdDomainJoinUser",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) ResetCrossRealmTrustPrincipalPassword() {
	_jsii_.InvokeVoid(
		a,
		"resetCrossRealmTrustPrincipalPassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsEmrCluster_KerberosAttributesPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

