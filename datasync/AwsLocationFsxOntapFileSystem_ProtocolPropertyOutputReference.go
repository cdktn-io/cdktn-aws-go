package datasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/datasync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/datasync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference interface {
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
	InternalValue() *AwsLocationFsxOntapFileSystem_ProtocolProperty
	// Experimental.
	SetInternalValue(val *AwsLocationFsxOntapFileSystem_ProtocolProperty)
	// Experimental.
	Nfs() AwsLocationFsxOntapFileSystem_NfsPropertyOutputReference
	// Experimental.
	NfsInput() *AwsLocationFsxOntapFileSystem_NfsProperty
	// Experimental.
	Smb() AwsLocationFsxOntapFileSystem_SmbPropertyOutputReference
	// Experimental.
	SmbInput() *AwsLocationFsxOntapFileSystem_SmbProperty
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
	PutNfs(value *AwsLocationFsxOntapFileSystem_NfsProperty)
	// Experimental.
	PutSmb(value *AwsLocationFsxOntapFileSystem_SmbProperty)
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

// The jsii proxy struct for AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference
type jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) InternalValue() *AwsLocationFsxOntapFileSystem_ProtocolProperty {
	var returns *AwsLocationFsxOntapFileSystem_ProtocolProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) Nfs() AwsLocationFsxOntapFileSystem_NfsPropertyOutputReference {
	var returns AwsLocationFsxOntapFileSystem_NfsPropertyOutputReference
	_jsii_.Get(
		j,
		"nfs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) NfsInput() *AwsLocationFsxOntapFileSystem_NfsProperty {
	var returns *AwsLocationFsxOntapFileSystem_NfsProperty
	_jsii_.Get(
		j,
		"nfsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) Smb() AwsLocationFsxOntapFileSystem_SmbPropertyOutputReference {
	var returns AwsLocationFsxOntapFileSystem_SmbPropertyOutputReference
	_jsii_.Get(
		j,
		"smb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) SmbInput() *AwsLocationFsxOntapFileSystem_SmbProperty {
	var returns *AwsLocationFsxOntapFileSystem_SmbProperty
	_jsii_.Get(
		j,
		"smbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-datasync.AwsLocationFsxOntapFileSystem.ProtocolPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference_Override(a AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-datasync.AwsLocationFsxOntapFileSystem.ProtocolPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference)SetInternalValue(val *AwsLocationFsxOntapFileSystem_ProtocolProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) PutNfs(value *AwsLocationFsxOntapFileSystem_NfsProperty) {
	if err := a.validatePutNfsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNfs",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) PutSmb(value *AwsLocationFsxOntapFileSystem_SmbProperty) {
	if err := a.validatePutSmbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSmb",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) ResetNfs() {
	_jsii_.InvokeVoid(
		a,
		"resetNfs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) ResetSmb() {
	_jsii_.InvokeVoid(
		a,
		"resetSmb",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsLocationFsxOntapFileSystem_ProtocolPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

