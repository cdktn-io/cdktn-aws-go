package dms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/dms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/dms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference interface {
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
	InternalValue() *AwsReplicationInstance_KerberosAuthenticationSettingsProperty
	// Experimental.
	SetInternalValue(val *AwsReplicationInstance_KerberosAuthenticationSettingsProperty)
	// Experimental.
	KeyCacheSecretIamArn() *string
	// Experimental.
	SetKeyCacheSecretIamArn(val *string)
	// Experimental.
	KeyCacheSecretIamArnInput() *string
	// Experimental.
	KeyCacheSecretId() *string
	// Experimental.
	SetKeyCacheSecretId(val *string)
	// Experimental.
	KeyCacheSecretIdInput() *string
	// Experimental.
	Krb5FileContents() *string
	// Experimental.
	SetKrb5FileContents(val *string)
	// Experimental.
	Krb5FileContentsInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference
type jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) InternalValue() *AwsReplicationInstance_KerberosAuthenticationSettingsProperty {
	var returns *AwsReplicationInstance_KerberosAuthenticationSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) KeyCacheSecretIamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyCacheSecretIamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) KeyCacheSecretIamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyCacheSecretIamArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) KeyCacheSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyCacheSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) KeyCacheSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyCacheSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) Krb5FileContents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krb5FileContents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) Krb5FileContentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krb5FileContentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.AwsReplicationInstance.KerberosAuthenticationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference_Override(a AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.AwsReplicationInstance.KerberosAuthenticationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetInternalValue(val *AwsReplicationInstance_KerberosAuthenticationSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetKeyCacheSecretIamArn(val *string) {
	if err := j.validateSetKeyCacheSecretIamArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyCacheSecretIamArn",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetKeyCacheSecretId(val *string) {
	if err := j.validateSetKeyCacheSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyCacheSecretId",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetKrb5FileContents(val *string) {
	if err := j.validateSetKrb5FileContentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"krb5FileContents",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

