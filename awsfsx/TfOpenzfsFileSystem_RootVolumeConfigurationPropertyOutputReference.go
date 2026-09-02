package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference interface {
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
	// Experimental.
	CopyTagsToSnapshots() interface{}
	// Experimental.
	SetCopyTagsToSnapshots(val interface{})
	// Experimental.
	CopyTagsToSnapshotsInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DataCompressionType() *string
	// Experimental.
	SetDataCompressionType(val *string)
	// Experimental.
	DataCompressionTypeInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfOpenzfsFileSystem_RootVolumeConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfOpenzfsFileSystem_RootVolumeConfigurationProperty)
	// Experimental.
	NfsExports() TfOpenzfsFileSystem_NfsExportsPropertyOutputReference
	// Experimental.
	NfsExportsInput() *TfOpenzfsFileSystem_NfsExportsProperty
	// Experimental.
	ReadOnly() interface{}
	// Experimental.
	SetReadOnly(val interface{})
	// Experimental.
	ReadOnlyInput() interface{}
	// Experimental.
	RecordSizeKib() *float64
	// Experimental.
	SetRecordSizeKib(val *float64)
	// Experimental.
	RecordSizeKibInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	UserAndGroupQuotas() TfOpenzfsFileSystem_UserAndGroupQuotasPropertyList
	// Experimental.
	UserAndGroupQuotasInput() interface{}
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
	PutNfsExports(value *TfOpenzfsFileSystem_NfsExportsProperty)
	// Experimental.
	PutUserAndGroupQuotas(value interface{})
	// Experimental.
	ResetCopyTagsToSnapshots()
	// Experimental.
	ResetDataCompressionType()
	// Experimental.
	ResetNfsExports()
	// Experimental.
	ResetReadOnly()
	// Experimental.
	ResetRecordSizeKib()
	// Experimental.
	ResetUserAndGroupQuotas()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference
type jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) CopyTagsToSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) CopyTagsToSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) DataCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) DataCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) InternalValue() *TfOpenzfsFileSystem_RootVolumeConfigurationProperty {
	var returns *TfOpenzfsFileSystem_RootVolumeConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) NfsExports() TfOpenzfsFileSystem_NfsExportsPropertyOutputReference {
	var returns TfOpenzfsFileSystem_NfsExportsPropertyOutputReference
	_jsii_.Get(
		j,
		"nfsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) NfsExportsInput() *TfOpenzfsFileSystem_NfsExportsProperty {
	var returns *TfOpenzfsFileSystem_NfsExportsProperty
	_jsii_.Get(
		j,
		"nfsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) RecordSizeKib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordSizeKib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) RecordSizeKibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordSizeKibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) UserAndGroupQuotas() TfOpenzfsFileSystem_UserAndGroupQuotasPropertyList {
	var returns TfOpenzfsFileSystem_UserAndGroupQuotasPropertyList
	_jsii_.Get(
		j,
		"userAndGroupQuotas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) UserAndGroupQuotasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAndGroupQuotasInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOpenzfsFileSystem.RootVolumeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference_Override(t TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.TfOpenzfsFileSystem.RootVolumeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetCopyTagsToSnapshots(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshots",
		val,
	)
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetDataCompressionType(val *string) {
	if err := j.validateSetDataCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCompressionType",
		val,
	)
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetInternalValue(val *TfOpenzfsFileSystem_RootVolumeConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetRecordSizeKib(val *float64) {
	if err := j.validateSetRecordSizeKibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordSizeKib",
		val,
	)
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) PutNfsExports(value *TfOpenzfsFileSystem_NfsExportsProperty) {
	if err := t.validatePutNfsExportsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putNfsExports",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) PutUserAndGroupQuotas(value interface{}) {
	if err := t.validatePutUserAndGroupQuotasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putUserAndGroupQuotas",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ResetCopyTagsToSnapshots() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyTagsToSnapshots",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ResetDataCompressionType() {
	_jsii_.InvokeVoid(
		t,
		"resetDataCompressionType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ResetNfsExports() {
	_jsii_.InvokeVoid(
		t,
		"resetNfsExports",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		t,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ResetRecordSizeKib() {
	_jsii_.InvokeVoid(
		t,
		"resetRecordSizeKib",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ResetUserAndGroupQuotas() {
	_jsii_.InvokeVoid(
		t,
		"resetUserAndGroupQuotas",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

