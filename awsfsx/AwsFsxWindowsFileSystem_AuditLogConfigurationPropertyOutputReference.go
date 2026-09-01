package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference interface {
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
	InternalValue() *AwsFsxWindowsFileSystem_AuditLogConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsFsxWindowsFileSystem_AuditLogConfigurationProperty)
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

// The jsii proxy struct for AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference
type jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) AuditLogDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) AuditLogDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"auditLogDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) FileAccessAuditLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileAccessAuditLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) FileAccessAuditLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileAccessAuditLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) FileShareAccessAuditLogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareAccessAuditLogLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) FileShareAccessAuditLogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileShareAccessAuditLogLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) InternalValue() *AwsFsxWindowsFileSystem_AuditLogConfigurationProperty {
	var returns *AwsFsxWindowsFileSystem_AuditLogConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxWindowsFileSystem.AuditLogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference_Override(a AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxWindowsFileSystem.AuditLogConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetAuditLogDestination(val *string) {
	if err := j.validateSetAuditLogDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"auditLogDestination",
		val,
	)
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetFileAccessAuditLogLevel(val *string) {
	if err := j.validateSetFileAccessAuditLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileAccessAuditLogLevel",
		val,
	)
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetFileShareAccessAuditLogLevel(val *string) {
	if err := j.validateSetFileShareAccessAuditLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileShareAccessAuditLogLevel",
		val,
	)
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetInternalValue(val *AwsFsxWindowsFileSystem_AuditLogConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ResetAuditLogDestination() {
	_jsii_.InvokeVoid(
		a,
		"resetAuditLogDestination",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ResetFileAccessAuditLogLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetFileAccessAuditLogLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ResetFileShareAccessAuditLogLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetFileShareAccessAuditLogLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFsxWindowsFileSystem_AuditLogConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

