package batch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/batch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/batch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AllowPrivilegeEscalation() interface{}
	// Experimental.
	SetAllowPrivilegeEscalation(val interface{})
	// Experimental.
	AllowPrivilegeEscalationInput() interface{}
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
	InternalValue() *AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty
	// Experimental.
	SetInternalValue(val *AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty)
	// Experimental.
	Privileged() interface{}
	// Experimental.
	SetPrivileged(val interface{})
	// Experimental.
	PrivilegedInput() interface{}
	// Experimental.
	ReadOnlyRootFileSystem() interface{}
	// Experimental.
	SetReadOnlyRootFileSystem(val interface{})
	// Experimental.
	ReadOnlyRootFileSystemInput() interface{}
	// Experimental.
	RunAsGroup() *float64
	// Experimental.
	SetRunAsGroup(val *float64)
	// Experimental.
	RunAsGroupInput() *float64
	// Experimental.
	RunAsNonRoot() interface{}
	// Experimental.
	SetRunAsNonRoot(val interface{})
	// Experimental.
	RunAsNonRootInput() interface{}
	// Experimental.
	RunAsUser() *float64
	// Experimental.
	SetRunAsUser(val *float64)
	// Experimental.
	RunAsUserInput() *float64
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
	ResetAllowPrivilegeEscalation()
	// Experimental.
	ResetPrivileged()
	// Experimental.
	ResetReadOnlyRootFileSystem()
	// Experimental.
	ResetRunAsGroup()
	// Experimental.
	ResetRunAsNonRoot()
	// Experimental.
	ResetRunAsUser()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference
type jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) AllowPrivilegeEscalation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPrivilegeEscalation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) AllowPrivilegeEscalationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPrivilegeEscalationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) InternalValue() *AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty {
	var returns *AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) Privileged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) PrivilegedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ReadOnlyRootFileSystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyRootFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ReadOnlyRootFileSystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyRootFileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) RunAsGroup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) RunAsGroupInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) RunAsNonRoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsNonRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) RunAsNonRootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsNonRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) RunAsUser() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) RunAsUserInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-batch.AwsJobDefinition.EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference_Override(a AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-batch.AwsJobDefinition.EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetAllowPrivilegeEscalation(val interface{}) {
	if err := j.validateSetAllowPrivilegeEscalationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPrivilegeEscalation",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetInternalValue(val *AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetPrivileged(val interface{}) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetReadOnlyRootFileSystem(val interface{}) {
	if err := j.validateSetReadOnlyRootFileSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnlyRootFileSystem",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetRunAsGroup(val *float64) {
	if err := j.validateSetRunAsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsGroup",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetRunAsNonRoot(val interface{}) {
	if err := j.validateSetRunAsNonRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsNonRoot",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetRunAsUser(val *float64) {
	if err := j.validateSetRunAsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUser",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ResetAllowPrivilegeEscalation() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowPrivilegeEscalation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		a,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ResetReadOnlyRootFileSystem() {
	_jsii_.InvokeVoid(
		a,
		"resetReadOnlyRootFileSystem",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ResetRunAsGroup() {
	_jsii_.InvokeVoid(
		a,
		"resetRunAsGroup",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ResetRunAsNonRoot() {
	_jsii_.InvokeVoid(
		a,
		"resetRunAsNonRoot",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ResetRunAsUser() {
	_jsii_.InvokeVoid(
		a,
		"resetRunAsUser",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

