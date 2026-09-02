package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference interface {
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
	InternalValue() *TfReplicationInstance_KerberosAuthenticationSettingsProperty
	// Experimental.
	SetInternalValue(val *TfReplicationInstance_KerberosAuthenticationSettingsProperty)
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

// The jsii proxy struct for TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference
type jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) InternalValue() *TfReplicationInstance_KerberosAuthenticationSettingsProperty {
	var returns *TfReplicationInstance_KerberosAuthenticationSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) KeyCacheSecretIamArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyCacheSecretIamArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) KeyCacheSecretIamArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyCacheSecretIamArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) KeyCacheSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyCacheSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) KeyCacheSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyCacheSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) Krb5FileContents() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krb5FileContents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) Krb5FileContentsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"krb5FileContentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.TfReplicationInstance.KerberosAuthenticationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference_Override(t TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.TfReplicationInstance.KerberosAuthenticationSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetInternalValue(val *TfReplicationInstance_KerberosAuthenticationSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetKeyCacheSecretIamArn(val *string) {
	if err := j.validateSetKeyCacheSecretIamArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyCacheSecretIamArn",
		val,
	)
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetKeyCacheSecretId(val *string) {
	if err := j.validateSetKeyCacheSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyCacheSecretId",
		val,
	)
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetKrb5FileContents(val *string) {
	if err := j.validateSetKrb5FileContentsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"krb5FileContents",
		val,
	)
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfReplicationInstance_KerberosAuthenticationSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

