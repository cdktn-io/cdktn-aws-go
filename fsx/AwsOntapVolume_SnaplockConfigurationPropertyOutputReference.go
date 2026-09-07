package fsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/fsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/fsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsOntapVolume_SnaplockConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuditLogVolume() interface{}
	// Experimental.
	SetAuditLogVolume(val interface{})
	// Experimental.
	AuditLogVolumeInput() interface{}
	// Experimental.
	AutocommitPeriod() AwsOntapVolume_AutocommitPeriodPropertyOutputReference
	// Experimental.
	AutocommitPeriodInput() *AwsOntapVolume_AutocommitPeriodProperty
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
	InternalValue() *AwsOntapVolume_SnaplockConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsOntapVolume_SnaplockConfigurationProperty)
	// Experimental.
	PrivilegedDelete() *string
	// Experimental.
	SetPrivilegedDelete(val *string)
	// Experimental.
	PrivilegedDeleteInput() *string
	// Experimental.
	RetentionPeriod() AwsOntapVolume_RetentionPeriodPropertyOutputReference
	// Experimental.
	RetentionPeriodInput() *AwsOntapVolume_RetentionPeriodProperty
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
	PutAutocommitPeriod(value *AwsOntapVolume_AutocommitPeriodProperty)
	// Experimental.
	PutRetentionPeriod(value *AwsOntapVolume_RetentionPeriodProperty)
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

// The jsii proxy struct for AwsOntapVolume_SnaplockConfigurationPropertyOutputReference
type jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) AuditLogVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditLogVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) AuditLogVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditLogVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) AutocommitPeriod() AwsOntapVolume_AutocommitPeriodPropertyOutputReference {
	var returns AwsOntapVolume_AutocommitPeriodPropertyOutputReference
	_jsii_.Get(
		j,
		"autocommitPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) AutocommitPeriodInput() *AwsOntapVolume_AutocommitPeriodProperty {
	var returns *AwsOntapVolume_AutocommitPeriodProperty
	_jsii_.Get(
		j,
		"autocommitPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) InternalValue() *AwsOntapVolume_SnaplockConfigurationProperty {
	var returns *AwsOntapVolume_SnaplockConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) PrivilegedDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegedDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) PrivilegedDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegedDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) RetentionPeriod() AwsOntapVolume_RetentionPeriodPropertyOutputReference {
	var returns AwsOntapVolume_RetentionPeriodPropertyOutputReference
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) RetentionPeriodInput() *AwsOntapVolume_RetentionPeriodProperty {
	var returns *AwsOntapVolume_RetentionPeriodProperty
	_jsii_.Get(
		j,
		"retentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) SnaplockType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snaplockType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) SnaplockTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snaplockTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) VolumeAppendModeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeAppendModeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) VolumeAppendModeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeAppendModeEnabledInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsOntapVolume_SnaplockConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsOntapVolume_SnaplockConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsOntapVolume_SnaplockConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsOntapVolume.SnaplockConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsOntapVolume_SnaplockConfigurationPropertyOutputReference_Override(a AwsOntapVolume_SnaplockConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsOntapVolume.SnaplockConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference)SetAuditLogVolume(val interface{}) {
	if err := j.validateSetAuditLogVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditLogVolume",
		val,
	)
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference)SetInternalValue(val *AwsOntapVolume_SnaplockConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference)SetPrivilegedDelete(val *string) {
	if err := j.validateSetPrivilegedDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedDelete",
		val,
	)
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference)SetSnaplockType(val *string) {
	if err := j.validateSetSnaplockTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snaplockType",
		val,
	)
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference)SetVolumeAppendModeEnabled(val interface{}) {
	if err := j.validateSetVolumeAppendModeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeAppendModeEnabled",
		val,
	)
}

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) PutAutocommitPeriod(value *AwsOntapVolume_AutocommitPeriodProperty) {
	if err := a.validatePutAutocommitPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutocommitPeriod",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) PutRetentionPeriod(value *AwsOntapVolume_RetentionPeriodProperty) {
	if err := a.validatePutRetentionPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetentionPeriod",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetAuditLogVolume() {
	_jsii_.InvokeVoid(
		a,
		"resetAuditLogVolume",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetAutocommitPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetAutocommitPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetPrivilegedDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivilegedDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetRetentionPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetRetentionPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetVolumeAppendModeEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetVolumeAppendModeEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsOntapVolume_SnaplockConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

