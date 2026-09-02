package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference interface {
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
	InternalValue() *TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationProperty)
	// Experimental.
	NetbiosName() *string
	// Experimental.
	SetNetbiosName(val *string)
	// Experimental.
	NetbiosNameInput() *string
	// Experimental.
	SelfManagedActiveDirectoryConfiguration() TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference
	// Experimental.
	SelfManagedActiveDirectoryConfigurationInput() *TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty
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
	PutSelfManagedActiveDirectoryConfiguration(value *TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty)
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

// The jsii proxy struct for TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference
type jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) InternalValue() *TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationProperty {
	var returns *TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) NetbiosName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netbiosName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) NetbiosNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"netbiosNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) SelfManagedActiveDirectoryConfiguration() TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference {
	var returns TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"selfManagedActiveDirectoryConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) SelfManagedActiveDirectoryConfigurationInput() *TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty {
	var returns *TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty
	_jsii_.Get(
		j,
		"selfManagedActiveDirectoryConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOntapStorageVirtualMachine.ActiveDirectoryConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference_Override(t TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOntapStorageVirtualMachine.ActiveDirectoryConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference)SetInternalValue(val *TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference)SetNetbiosName(val *string) {
	if err := j.validateSetNetbiosNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"netbiosName",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) PutSelfManagedActiveDirectoryConfiguration(value *TfOntapStorageVirtualMachine_SelfManagedActiveDirectoryConfigurationProperty) {
	if err := t.validatePutSelfManagedActiveDirectoryConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSelfManagedActiveDirectoryConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) ResetNetbiosName() {
	_jsii_.InvokeVoid(
		t,
		"resetNetbiosName",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) ResetSelfManagedActiveDirectoryConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSelfManagedActiveDirectoryConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfOntapStorageVirtualMachine_ActiveDirectoryConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

