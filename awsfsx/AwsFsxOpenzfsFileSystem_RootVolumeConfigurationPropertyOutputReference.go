package awsfsx

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsfsx/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsfsx/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference interface {
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
	InternalValue() *AwsFsxOpenzfsFileSystem_RootVolumeConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsFsxOpenzfsFileSystem_RootVolumeConfigurationProperty)
	// Experimental.
	NfsExports() AwsFsxOpenzfsFileSystem_NfsExportsPropertyOutputReference
	// Experimental.
	NfsExportsInput() *AwsFsxOpenzfsFileSystem_NfsExportsProperty
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
	UserAndGroupQuotas() AwsFsxOpenzfsFileSystem_UserAndGroupQuotasPropertyList
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
	PutNfsExports(value *AwsFsxOpenzfsFileSystem_NfsExportsProperty)
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

// The jsii proxy struct for AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference
type jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) CopyTagsToSnapshots() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshots",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) CopyTagsToSnapshotsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"copyTagsToSnapshotsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) DataCompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) DataCompressionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataCompressionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) InternalValue() *AwsFsxOpenzfsFileSystem_RootVolumeConfigurationProperty {
	var returns *AwsFsxOpenzfsFileSystem_RootVolumeConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) NfsExports() AwsFsxOpenzfsFileSystem_NfsExportsPropertyOutputReference {
	var returns AwsFsxOpenzfsFileSystem_NfsExportsPropertyOutputReference
	_jsii_.Get(
		j,
		"nfsExports",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) NfsExportsInput() *AwsFsxOpenzfsFileSystem_NfsExportsProperty {
	var returns *AwsFsxOpenzfsFileSystem_NfsExportsProperty
	_jsii_.Get(
		j,
		"nfsExportsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ReadOnly() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ReadOnlyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"readOnlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) RecordSizeKib() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordSizeKib",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) RecordSizeKibInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"recordSizeKibInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) UserAndGroupQuotas() AwsFsxOpenzfsFileSystem_UserAndGroupQuotasPropertyList {
	var returns AwsFsxOpenzfsFileSystem_UserAndGroupQuotasPropertyList
	_jsii_.Get(
		j,
		"userAndGroupQuotas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) UserAndGroupQuotasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"userAndGroupQuotasInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxOpenzfsFileSystem.RootVolumeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference_Override(a AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-fsx.AwsFsxOpenzfsFileSystem.RootVolumeConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetCopyTagsToSnapshots(val interface{}) {
	if err := j.validateSetCopyTagsToSnapshotsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyTagsToSnapshots",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetDataCompressionType(val *string) {
	if err := j.validateSetDataCompressionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataCompressionType",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetInternalValue(val *AwsFsxOpenzfsFileSystem_RootVolumeConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetReadOnly(val interface{}) {
	if err := j.validateSetReadOnlyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readOnly",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetRecordSizeKib(val *float64) {
	if err := j.validateSetRecordSizeKibParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recordSizeKib",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) PutNfsExports(value *AwsFsxOpenzfsFileSystem_NfsExportsProperty) {
	if err := a.validatePutNfsExportsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putNfsExports",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) PutUserAndGroupQuotas(value interface{}) {
	if err := a.validatePutUserAndGroupQuotasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putUserAndGroupQuotas",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ResetCopyTagsToSnapshots() {
	_jsii_.InvokeVoid(
		a,
		"resetCopyTagsToSnapshots",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ResetDataCompressionType() {
	_jsii_.InvokeVoid(
		a,
		"resetDataCompressionType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ResetNfsExports() {
	_jsii_.InvokeVoid(
		a,
		"resetNfsExports",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ResetReadOnly() {
	_jsii_.InvokeVoid(
		a,
		"resetReadOnly",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ResetRecordSizeKib() {
	_jsii_.InvokeVoid(
		a,
		"resetRecordSizeKib",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ResetUserAndGroupQuotas() {
	_jsii_.InvokeVoid(
		a,
		"resetUserAndGroupQuotas",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsFsxOpenzfsFileSystem_RootVolumeConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

