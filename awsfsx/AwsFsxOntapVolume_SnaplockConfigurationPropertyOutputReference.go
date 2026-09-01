package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuditLogVolume() interface{}
	// Experimental.
	SetAuditLogVolume(val interface{})
	// Experimental.
	AuditLogVolumeInput() interface{}
	// Experimental.
	AutocommitPeriod() AwsFsxOntapVolume_AutocommitPeriodPropertyOutputReference
	// Experimental.
	AutocommitPeriodInput() *AwsFsxOntapVolume_AutocommitPeriodProperty
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
	InternalValue() *AwsFsxOntapVolume_SnaplockConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsFsxOntapVolume_SnaplockConfigurationProperty)
	// Experimental.
	PrivilegedDelete() *string
	// Experimental.
	SetPrivilegedDelete(val *string)
	// Experimental.
	PrivilegedDeleteInput() *string
	// Experimental.
	RetentionPeriod() AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference
	// Experimental.
	RetentionPeriodInput() *AwsFsxOntapVolume_RetentionPeriodProperty
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
	PutAutocommitPeriod(value *AwsFsxOntapVolume_AutocommitPeriodProperty)
	// Experimental.
	PutRetentionPeriod(value *AwsFsxOntapVolume_RetentionPeriodProperty)
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

// The jsii proxy struct for AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference
type jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) AuditLogVolume() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditLogVolume",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) AuditLogVolumeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"auditLogVolumeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) AutocommitPeriod() AwsFsxOntapVolume_AutocommitPeriodPropertyOutputReference {
	var returns AwsFsxOntapVolume_AutocommitPeriodPropertyOutputReference
	_jsii_.Get(
		j,
		"autocommitPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) AutocommitPeriodInput() *AwsFsxOntapVolume_AutocommitPeriodProperty {
	var returns *AwsFsxOntapVolume_AutocommitPeriodProperty
	_jsii_.Get(
		j,
		"autocommitPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) InternalValue() *AwsFsxOntapVolume_SnaplockConfigurationProperty {
	var returns *AwsFsxOntapVolume_SnaplockConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) PrivilegedDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegedDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) PrivilegedDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privilegedDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) RetentionPeriod() AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference {
	var returns AwsFsxOntapVolume_RetentionPeriodPropertyOutputReference
	_jsii_.Get(
		j,
		"retentionPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) RetentionPeriodInput() *AwsFsxOntapVolume_RetentionPeriodProperty {
	var returns *AwsFsxOntapVolume_RetentionPeriodProperty
	_jsii_.Get(
		j,
		"retentionPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) SnaplockType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snaplockType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) SnaplockTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"snaplockTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) VolumeAppendModeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeAppendModeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) VolumeAppendModeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeAppendModeEnabledInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxOntapVolume.SnaplockConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference_Override(a AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxOntapVolume.SnaplockConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference)SetAuditLogVolume(val interface{}) {
	if err := j.validateSetAuditLogVolumeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditLogVolume",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference)SetInternalValue(val *AwsFsxOntapVolume_SnaplockConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference)SetPrivilegedDelete(val *string) {
	if err := j.validateSetPrivilegedDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privilegedDelete",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference)SetSnaplockType(val *string) {
	if err := j.validateSetSnaplockTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"snaplockType",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference)SetVolumeAppendModeEnabled(val interface{}) {
	if err := j.validateSetVolumeAppendModeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeAppendModeEnabled",
		val,
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) PutAutocommitPeriod(value *AwsFsxOntapVolume_AutocommitPeriodProperty) {
	if err := a.validatePutAutocommitPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAutocommitPeriod",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) PutRetentionPeriod(value *AwsFsxOntapVolume_RetentionPeriodProperty) {
	if err := a.validatePutRetentionPeriodParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetentionPeriod",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetAuditLogVolume() {
	_jsii_.InvokeVoid(
		a,
		"resetAuditLogVolume",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetAutocommitPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetAutocommitPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetPrivilegedDelete() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivilegedDelete",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetRetentionPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetRetentionPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) ResetVolumeAppendModeEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetVolumeAppendModeEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFsxOntapVolume_SnaplockConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

