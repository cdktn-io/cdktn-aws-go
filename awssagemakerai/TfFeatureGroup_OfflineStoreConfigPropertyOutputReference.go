package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfFeatureGroup_OfflineStoreConfigPropertyOutputReference interface {
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
	DataCatalogConfig() TfFeatureGroup_DataCatalogConfigPropertyOutputReference
	// Experimental.
	DataCatalogConfigInput() *TfFeatureGroup_DataCatalogConfigProperty
	// Experimental.
	DisableGlueTableCreation() interface{}
	// Experimental.
	SetDisableGlueTableCreation(val interface{})
	// Experimental.
	DisableGlueTableCreationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfFeatureGroup_OfflineStoreConfigProperty
	// Experimental.
	SetInternalValue(val *TfFeatureGroup_OfflineStoreConfigProperty)
	// Experimental.
	S3StorageConfig() TfFeatureGroup_S3StorageConfigPropertyOutputReference
	// Experimental.
	S3StorageConfigInput() *TfFeatureGroup_S3StorageConfigProperty
	// Experimental.
	TableFormat() *string
	// Experimental.
	SetTableFormat(val *string)
	// Experimental.
	TableFormatInput() *string
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
	PutDataCatalogConfig(value *TfFeatureGroup_DataCatalogConfigProperty)
	// Experimental.
	PutS3StorageConfig(value *TfFeatureGroup_S3StorageConfigProperty)
	// Experimental.
	ResetDataCatalogConfig()
	// Experimental.
	ResetDisableGlueTableCreation()
	// Experimental.
	ResetTableFormat()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfFeatureGroup_OfflineStoreConfigPropertyOutputReference
type jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) DataCatalogConfig() TfFeatureGroup_DataCatalogConfigPropertyOutputReference {
	var returns TfFeatureGroup_DataCatalogConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"dataCatalogConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) DataCatalogConfigInput() *TfFeatureGroup_DataCatalogConfigProperty {
	var returns *TfFeatureGroup_DataCatalogConfigProperty
	_jsii_.Get(
		j,
		"dataCatalogConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) DisableGlueTableCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableGlueTableCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) DisableGlueTableCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableGlueTableCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) InternalValue() *TfFeatureGroup_OfflineStoreConfigProperty {
	var returns *TfFeatureGroup_OfflineStoreConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) S3StorageConfig() TfFeatureGroup_S3StorageConfigPropertyOutputReference {
	var returns TfFeatureGroup_S3StorageConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"s3StorageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) S3StorageConfigInput() *TfFeatureGroup_S3StorageConfigProperty {
	var returns *TfFeatureGroup_S3StorageConfigProperty
	_jsii_.Get(
		j,
		"s3StorageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) TableFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) TableFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfFeatureGroup_OfflineStoreConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfFeatureGroup_OfflineStoreConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfFeatureGroup_OfflineStoreConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfFeatureGroup.OfflineStoreConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfFeatureGroup_OfflineStoreConfigPropertyOutputReference_Override(t TfFeatureGroup_OfflineStoreConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.TfFeatureGroup.OfflineStoreConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetDisableGlueTableCreation(val interface{}) {
	if err := j.validateSetDisableGlueTableCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableGlueTableCreation",
		val,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetInternalValue(val *TfFeatureGroup_OfflineStoreConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetTableFormat(val *string) {
	if err := j.validateSetTableFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableFormat",
		val,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) PutDataCatalogConfig(value *TfFeatureGroup_DataCatalogConfigProperty) {
	if err := t.validatePutDataCatalogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putDataCatalogConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) PutS3StorageConfig(value *TfFeatureGroup_S3StorageConfigProperty) {
	if err := t.validatePutS3StorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3StorageConfig",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) ResetDataCatalogConfig() {
	_jsii_.InvokeVoid(
		t,
		"resetDataCatalogConfig",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) ResetDisableGlueTableCreation() {
	_jsii_.InvokeVoid(
		t,
		"resetDisableGlueTableCreation",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) ResetTableFormat() {
	_jsii_.InvokeVoid(
		t,
		"resetTableFormat",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfFeatureGroup_OfflineStoreConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

