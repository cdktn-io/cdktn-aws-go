package awsglue

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsglue/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsglue/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference interface {
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
	// Experimental.
	ConnectionPasswordEncryption() TfDataCatalogEncryptionSettings_ConnectionPasswordEncryptionPropertyOutputReference
	// Experimental.
	ConnectionPasswordEncryptionInput() *TfDataCatalogEncryptionSettings_ConnectionPasswordEncryptionProperty
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	EncryptionAtRest() TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference
	// Experimental.
	EncryptionAtRestInput() *TfDataCatalogEncryptionSettings_EncryptionAtRestProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsProperty
	// Experimental.
	SetInternalValue(val *TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsProperty)
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
	PutConnectionPasswordEncryption(value *TfDataCatalogEncryptionSettings_ConnectionPasswordEncryptionProperty)
	// Experimental.
	PutEncryptionAtRest(value *TfDataCatalogEncryptionSettings_EncryptionAtRestProperty)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference
type jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) ConnectionPasswordEncryption() TfDataCatalogEncryptionSettings_ConnectionPasswordEncryptionPropertyOutputReference {
	var returns TfDataCatalogEncryptionSettings_ConnectionPasswordEncryptionPropertyOutputReference
	_jsii_.Get(
		j,
		"connectionPasswordEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) ConnectionPasswordEncryptionInput() *TfDataCatalogEncryptionSettings_ConnectionPasswordEncryptionProperty {
	var returns *TfDataCatalogEncryptionSettings_ConnectionPasswordEncryptionProperty
	_jsii_.Get(
		j,
		"connectionPasswordEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) EncryptionAtRest() TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference {
	var returns TfDataCatalogEncryptionSettings_EncryptionAtRestPropertyOutputReference
	_jsii_.Get(
		j,
		"encryptionAtRest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) EncryptionAtRestInput() *TfDataCatalogEncryptionSettings_EncryptionAtRestProperty {
	var returns *TfDataCatalogEncryptionSettings_EncryptionAtRestProperty
	_jsii_.Get(
		j,
		"encryptionAtRestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) InternalValue() *TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsProperty {
	var returns *TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-glue.TfDataCatalogEncryptionSettings.DataCatalogEncryptionSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference_Override(t TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-glue.TfDataCatalogEncryptionSettings.DataCatalogEncryptionSettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference)SetInternalValue(val *TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) PutConnectionPasswordEncryption(value *TfDataCatalogEncryptionSettings_ConnectionPasswordEncryptionProperty) {
	if err := t.validatePutConnectionPasswordEncryptionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putConnectionPasswordEncryption",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) PutEncryptionAtRest(value *TfDataCatalogEncryptionSettings_EncryptionAtRestProperty) {
	if err := t.validatePutEncryptionAtRestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putEncryptionAtRest",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDataCatalogEncryptionSettings_DataCatalogEncryptionSettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

