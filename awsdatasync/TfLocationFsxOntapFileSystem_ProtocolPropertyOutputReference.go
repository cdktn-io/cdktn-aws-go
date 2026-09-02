package awsdatasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdatasync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdatasync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference interface {
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
	InternalValue() *TfLocationFsxOntapFileSystem_ProtocolProperty
	// Experimental.
	SetInternalValue(val *TfLocationFsxOntapFileSystem_ProtocolProperty)
	// Experimental.
	Nfs() TfLocationFsxOntapFileSystem_NfsPropertyOutputReference
	// Experimental.
	NfsInput() *TfLocationFsxOntapFileSystem_NfsProperty
	// Experimental.
	Smb() TfLocationFsxOntapFileSystem_SmbPropertyOutputReference
	// Experimental.
	SmbInput() *TfLocationFsxOntapFileSystem_SmbProperty
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
	PutNfs(value *TfLocationFsxOntapFileSystem_NfsProperty)
	// Experimental.
	PutSmb(value *TfLocationFsxOntapFileSystem_SmbProperty)
	// Experimental.
	ResetNfs()
	// Experimental.
	ResetSmb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference
type jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) InternalValue() *TfLocationFsxOntapFileSystem_ProtocolProperty {
	var returns *TfLocationFsxOntapFileSystem_ProtocolProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) Nfs() TfLocationFsxOntapFileSystem_NfsPropertyOutputReference {
	var returns TfLocationFsxOntapFileSystem_NfsPropertyOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) NfsInput() *TfLocationFsxOntapFileSystem_NfsProperty {
	var returns *TfLocationFsxOntapFileSystem_NfsProperty
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) Smb() TfLocationFsxOntapFileSystem_SmbPropertyOutputReference {
	var returns TfLocationFsxOntapFileSystem_SmbPropertyOutputReference
	_jsii_.Get(
		j,
		"smb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) SmbInput() *TfLocationFsxOntapFileSystem_SmbProperty {
	var returns *TfLocationFsxOntapFileSystem_SmbProperty
	_jsii_.Get(
		j,
		"smbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfLocationFsxOntapFileSystem_ProtocolPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-datasync.TfLocationFsxOntapFileSystem.ProtocolPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference_Override(t TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-datasync.TfLocationFsxOntapFileSystem.ProtocolPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference)SetInternalValue(val *TfLocationFsxOntapFileSystem_ProtocolProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) PutNfs(value *TfLocationFsxOntapFileSystem_NfsProperty) {
	if err := t.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNfs",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) PutSmb(value *TfLocationFsxOntapFileSystem_SmbProperty) {
	if err := t.validatePutSmbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSmb",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		t,
		"resetNfs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) ResetSmb() {
	_jsii_.InvokeVoid(
		t,
		"resetSmb",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

