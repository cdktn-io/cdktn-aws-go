package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference interface {
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
	InternalValue() *AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationProperty)
	// Experimental.
	NetbiosName() *string
	// Experimental.
	SetNetbiosName(val *string)
	// Experimental.
	NetbiosNameInput() *string
	// Experimental.
	SelfManagedActiveDirectoryConfiguration() AwsFsxOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference
	// Experimental.
	SelfManagedActiveDirectoryConfigurationInput() *AwsFsxOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty
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
	PutSelfManagedActiveDirectoryConfiguration(value *AwsFsxOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty)
	// Experimental.
	ResetNetbiosName()
	// Experimental.
	ResetSelfManagedActiveDirectoryConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference
type jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) InternalValue() *AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationProperty {
	var returns *AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) NetbiosName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netbiosName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) NetbiosNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netbiosNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) SelfManagedActiveDirectoryConfiguration() AwsFsxOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference {
	var returns AwsFsxOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"selfManagedActiveDirectoryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) SelfManagedActiveDirectoryConfigurationInput() *AwsFsxOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty {
	var returns *AwsFsxOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty
	_jsii_.Get(
		j,
		"selfManagedActiveDirectoryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxOntapStorageVirtualMachine.ActiveDirectoryConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference_Override(a AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxOntapStorageVirtualMachine.ActiveDirectoryConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference)SetInternalValue(val *AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference)SetNetbiosName(val *string) {
	if err := j.validateSetNetbiosNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netbiosName",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) PutSelfManagedActiveDirectoryConfiguration(value *AwsFsxOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty) {
	if err := a.validatePutSelfManagedActiveDirectoryConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSelfManagedActiveDirectoryConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) ResetNetbiosName() {
	_jsii_.InvokeVoid(
		a,
		"resetNetbiosName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) ResetSelfManagedActiveDirectoryConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSelfManagedActiveDirectoryConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFsxOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

