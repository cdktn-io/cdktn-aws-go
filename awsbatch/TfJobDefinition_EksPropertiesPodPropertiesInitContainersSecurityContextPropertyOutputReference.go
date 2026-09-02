package awsbatch

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsbatch/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsbatch/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference interface {
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
	InternalValue() *TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty
	// Experimental.
	SetInternalValue(val *TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty)
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

// The jsii proxy struct for TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference
type jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) AllowPrivilegeEscalation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPrivilegeEscalation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) AllowPrivilegeEscalationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowPrivilegeEscalationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) InternalValue() *TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty {
	var returns *TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) Privileged() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privileged",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) PrivilegedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privilegedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ReadOnlyRootFileSystem() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyRootFileSystem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ReadOnlyRootFileSystemInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyRootFileSystemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) RunAsGroup() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) RunAsGroupInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) RunAsNonRoot() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsNonRoot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) RunAsNonRootInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runAsNonRootInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) RunAsUser() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) RunAsUserInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"runAsUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-batch.TfJobDefinition.EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference_Override(t TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-batch.TfJobDefinition.EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetAllowPrivilegeEscalation(val interface{}) {
	if err := j.validateSetAllowPrivilegeEscalationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowPrivilegeEscalation",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetInternalValue(val *TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetPrivileged(val interface{}) {
	if err := j.validateSetPrivilegedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privileged",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetReadOnlyRootFileSystem(val interface{}) {
	if err := j.validateSetReadOnlyRootFileSystemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnlyRootFileSystem",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetRunAsGroup(val *float64) {
	if err := j.validateSetRunAsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsGroup",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetRunAsNonRoot(val interface{}) {
	if err := j.validateSetRunAsNonRootParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsNonRoot",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetRunAsUser(val *float64) {
	if err := j.validateSetRunAsUserParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runAsUser",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ResetAllowPrivilegeEscalation() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowPrivilegeEscalation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ResetPrivileged() {
	_jsii_.InvokeVoid(
		t,
		"resetPrivileged",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ResetReadOnlyRootFileSystem() {
	_jsii_.InvokeVoid(
		t,
		"resetReadOnlyRootFileSystem",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ResetRunAsGroup() {
	_jsii_.InvokeVoid(
		t,
		"resetRunAsGroup",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ResetRunAsNonRoot() {
	_jsii_.InvokeVoid(
		t,
		"resetRunAsNonRoot",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ResetRunAsUser() {
	_jsii_.InvokeVoid(
		t,
		"resetRunAsUser",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfJobDefinition_EksPropertiesPodPropertiesInitContainersSecurityContextPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

