package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AuditLogDestination() *string
	// Experimental.
	SetAuditLogDestination(val *string)
	// Experimental.
	AuditLogDestinationInput() *string
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
	FileAccessAuditLogLevel() *string
	// Experimental.
	SetFileAccessAuditLogLevel(val *string)
	// Experimental.
	FileAccessAuditLogLevelInput() *string
	// Experimental.
	FileShareAccessAuditLogLevel() *string
	// Experimental.
	SetFileShareAccessAuditLogLevel(val *string)
	// Experimental.
	FileShareAccessAuditLogLevelInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfWindowsFileSystem_AuditLogConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfWindowsFileSystem_AuditLogConfigurationProperty)
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
	ResetAuditLogDestination()
	// Experimental.
	ResetFileAccessAuditLogLevel()
	// Experimental.
	ResetFileShareAccessAuditLogLevel()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference
type jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) AuditLogDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) AuditLogDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) FileAccessAuditLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileAccessAuditLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) FileAccessAuditLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileAccessAuditLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) FileShareAccessAuditLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareAccessAuditLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) FileShareAccessAuditLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareAccessAuditLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) InternalValue() *TfWindowsFileSystem_AuditLogConfigurationProperty {
	var returns *TfWindowsFileSystem_AuditLogConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfWindowsFileSystem_AuditLogConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.TfWindowsFileSystem.AuditLogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference_Override(t TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.TfWindowsFileSystem.AuditLogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetAuditLogDestination(val *string) {
	if err := j.validateSetAuditLogDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditLogDestination",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetFileAccessAuditLogLevel(val *string) {
	if err := j.validateSetFileAccessAuditLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileAccessAuditLogLevel",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetFileShareAccessAuditLogLevel(val *string) {
	if err := j.validateSetFileShareAccessAuditLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileShareAccessAuditLogLevel",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetInternalValue(val *TfWindowsFileSystem_AuditLogConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ResetAuditLogDestination() {
	_jsii_.InvokeVoid(
		t,
		"resetAuditLogDestination",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ResetFileAccessAuditLogLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetFileAccessAuditLogLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ResetFileShareAccessAuditLogLevel() {
	_jsii_.InvokeVoid(
		t,
		"resetFileShareAccessAuditLogLevel",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

