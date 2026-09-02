package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CatalogEncryptionMode() *string
	// Experimental.
	SetCatalogEncryptionMode(val *string)
	// Experimental.
	CatalogEncryptionModeInput() *string
	// Experimental.
	CatalogEncryptionServiceRole() *string
	// Experimental.
	SetCatalogEncryptionServiceRole(val *string)
	// Experimental.
	CatalogEncryptionServiceRoleInput() *string
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
	InternalValue() *TfDataCatalogEncryptionSettings_EncryptionAtRestProperty
	// Experimental.
	SetInternalValue(val *TfDataCatalogEncryptionSettings_EncryptionAtRestProperty)
	// Experimental.
	SseAwsKmsKeyId() *string
	// Experimental.
	SetSseAwsKmsKeyId(val *string)
	// Experimental.
	SseAwsKmsKeyIdInput() *string
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
	ResetCatalogEncryptionServiceRole()
	// Experimental.
	ResetSseAwsKmsKeyId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference
type jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) CatalogEncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogEncryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) CatalogEncryptionModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogEncryptionModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) CatalogEncryptionServiceRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogEncryptionServiceRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) CatalogEncryptionServiceRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"catalogEncryptionServiceRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) InternalValue() *TfDataCatalogEncryptionSettings_EncryptionAtRestProperty {
	var returns *TfDataCatalogEncryptionSettings_EncryptionAtRestProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) SseAwsKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sseAwsKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) SseAwsKmsKeyIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sseAwsKmsKeyIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfDataCatalogEncryptionSettings.EncryptionAtRestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference_Override(t TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfDataCatalogEncryptionSettings.EncryptionAtRestPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetCatalogEncryptionMode(val *string) {
	if err := j.validateSetCatalogEncryptionModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogEncryptionMode",
		val,
	)
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetCatalogEncryptionServiceRole(val *string) {
	if err := j.validateSetCatalogEncryptionServiceRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"catalogEncryptionServiceRole",
		val,
	)
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetInternalValue(val *TfDataCatalogEncryptionSettings_EncryptionAtRestProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetSseAwsKmsKeyId(val *string) {
	if err := j.validateSetSseAwsKmsKeyIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sseAwsKmsKeyId",
		val,
	)
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) ResetCatalogEncryptionServiceRole() {
	_jsii_.InvokeVoid(
		t,
		"resetCatalogEncryptionServiceRole",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) ResetSseAwsKmsKeyId() {
	_jsii_.InvokeVoid(
		t,
		"resetSseAwsKmsKeyId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

