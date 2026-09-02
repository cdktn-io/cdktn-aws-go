package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOntapVolume_SnaplockConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuditLogVolume() interface{}
	// Experimental.
	SetAuditLogVolume(val interface{})
	// Experimental.
	AuditLogVolumeInput() interface{}
	// Experimental.
	AutocommitPeriod() TfOntapVolume_AutocommitPeriodPropertyOutputReference
	// Experimental.
	AutocommitPeriodInput() *TfOntapVolume_AutocommitPeriodProperty
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
	InternalValue() *TfOntapVolume_SnaplockConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfOntapVolume_SnaplockConfigurationProperty)
	// Experimental.
	PrivilegedDelete() *string
	// Experimental.
	SetPrivilegedDelete(val *string)
	// Experimental.
	PrivilegedDeleteInput() *string
	// Experimental.
	RetentionPeriod() TfOntapVolume_RetentionPeriodPropertyOutputReference
	// Experimental.
	RetentionPeriodInput() *TfOntapVolume_RetentionPeriodProperty
	// Experimental.
	SnaplockType() *string
	// Experimental.
	SetSnaplockType(val *string)
	// Experimental.
	SnaplockTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VolumeAppendModeEnabled() interface{}
	// Experimental.
	SetVolumeAppendModeEnabled(val interface{})
	// Experimental.
	VolumeAppendModeEnabledInput() interface{}
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
	PutAutocommitPeriod(value *TfOntapVolume_AutocommitPeriodProperty)
	// Experimental.
	PutRetentionPeriod(value *TfOntapVolume_RetentionPeriodProperty)
	// Experimental.
	ResetAuditLogVolume()
	// Experimental.
	ResetAutocommitPeriod()
	// Experimental.
	ResetPrivilegedDelete()
	// Experimental.
	ResetRetentionPeriod()
	// Experimental.
	ResetVolumeAppendModeEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfOntapVolume_SnaplockConfigurationPropertyOutputReference
type jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) AuditLogVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditLogVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) AuditLogVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditLogVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) AutocommitPeriod() TfOntapVolume_AutocommitPeriodPropertyOutputReference {
	var returns TfOntapVolume_AutocommitPeriodPropertyOutputReference
	_jsii_.Get(
		j,
		"autocommitPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) AutocommitPeriodInput() *TfOntapVolume_AutocommitPeriodProperty {
	var returns *TfOntapVolume_AutocommitPeriodProperty
	_jsii_.Get(
		j,
		"autocommitPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) InternalValue() *TfOntapVolume_SnaplockConfigurationProperty {
	var returns *TfOntapVolume_SnaplockConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) PrivilegedDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegedDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) PrivilegedDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegedDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) RetentionPeriod() TfOntapVolume_RetentionPeriodPropertyOutputReference {
	var returns TfOntapVolume_RetentionPeriodPropertyOutputReference
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) RetentionPeriodInput() *TfOntapVolume_RetentionPeriodProperty {
	var returns *TfOntapVolume_RetentionPeriodProperty
	_jsii_.Get(
		j,
		"retentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) SnaplockType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snaplockType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) SnaplockTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snaplockTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) VolumeAppendModeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeAppendModeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) VolumeAppendModeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeAppendModeEnabledInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfOntapVolume_SnaplockConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfOntapVolume_SnaplockConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfOntapVolume_SnaplockConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOntapVolume.SnaplockConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfOntapVolume_SnaplockConfigurationPropertyOutputReference_Override(t TfOntapVolume_SnaplockConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOntapVolume.SnaplockConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference)SetAuditLogVolume(val interface{}) {
	if err := j.validateSetAuditLogVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditLogVolume",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference)SetInternalValue(val *TfOntapVolume_SnaplockConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference)SetPrivilegedDelete(val *string) {
	if err := j.validateSetPrivilegedDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedDelete",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference)SetSnaplockType(val *string) {
	if err := j.validateSetSnaplockTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snaplockType",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference)SetVolumeAppendModeEnabled(val interface{}) {
	if err := j.validateSetVolumeAppendModeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeAppendModeEnabled",
		val,
	)
}

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) PutAutocommitPeriod(value *TfOntapVolume_AutocommitPeriodProperty) {
	if err := t.validatePutAutocommitPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putAutocommitPeriod",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) PutRetentionPeriod(value *TfOntapVolume_RetentionPeriodProperty) {
	if err := t.validatePutRetentionPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putRetentionPeriod",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetAuditLogVolume() {
	_jsii_.InvokeVoid(
		t,
		"resetAuditLogVolume",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetAutocommitPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetAutocommitPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetPrivilegedDelete() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivilegedDelete",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetRetentionPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetRetentionPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetVolumeAppendModeEnabled() {
	_jsii_.InvokeVoid(
		t,
		"resetVolumeAppendModeEnabled",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfOntapVolume_SnaplockConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

