package awssagemakerai

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awssagemakerai/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awssagemakerai/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference interface {
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
	DataCatalogConfig() AwsSagemakerFeatureGroup_DataCatalogConfigPropertyOutputReference
	// Experimental.
	DataCatalogConfigInput() *AwsSagemakerFeatureGroup_DataCatalogConfigProperty
	// Experimental.
	DisableGlueTableCreation() interface{}
	// Experimental.
	SetDisableGlueTableCreation(val interface{})
	// Experimental.
	DisableGlueTableCreationInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsSagemakerFeatureGroup_OfflineStoreConfigProperty
	// Experimental.
	SetInternalValue(val *AwsSagemakerFeatureGroup_OfflineStoreConfigProperty)
	// Experimental.
	S3StorageConfig() AwsSagemakerFeatureGroup_S3StorageConfigPropertyOutputReference
	// Experimental.
	S3StorageConfigInput() *AwsSagemakerFeatureGroup_S3StorageConfigProperty
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
	PutDataCatalogConfig(value *AwsSagemakerFeatureGroup_DataCatalogConfigProperty)
	// Experimental.
	PutS3StorageConfig(value *AwsSagemakerFeatureGroup_S3StorageConfigProperty)
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

// The jsii proxy struct for AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference
type jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) DataCatalogConfig() AwsSagemakerFeatureGroup_DataCatalogConfigPropertyOutputReference {
	var returns AwsSagemakerFeatureGroup_DataCatalogConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"dataCatalogConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) DataCatalogConfigInput() *AwsSagemakerFeatureGroup_DataCatalogConfigProperty {
	var returns *AwsSagemakerFeatureGroup_DataCatalogConfigProperty
	_jsii_.Get(
		j,
		"dataCatalogConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) DisableGlueTableCreation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableGlueTableCreation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) DisableGlueTableCreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableGlueTableCreationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) InternalValue() *AwsSagemakerFeatureGroup_OfflineStoreConfigProperty {
	var returns *AwsSagemakerFeatureGroup_OfflineStoreConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) S3StorageConfig() AwsSagemakerFeatureGroup_S3StorageConfigPropertyOutputReference {
	var returns AwsSagemakerFeatureGroup_S3StorageConfigPropertyOutputReference
	_jsii_.Get(
		j,
		"s3StorageConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) S3StorageConfigInput() *AwsSagemakerFeatureGroup_S3StorageConfigProperty {
	var returns *AwsSagemakerFeatureGroup_S3StorageConfigProperty
	_jsii_.Get(
		j,
		"s3StorageConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) TableFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) TableFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerFeatureGroup.OfflineStoreConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference_Override(a AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-sagemaker-ai.AwsSagemakerFeatureGroup.OfflineStoreConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetDisableGlueTableCreation(val interface{}) {
	if err := j.validateSetDisableGlueTableCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableGlueTableCreation",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetInternalValue(val *AwsSagemakerFeatureGroup_OfflineStoreConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetTableFormat(val *string) {
	if err := j.validateSetTableFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tableFormat",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) PutDataCatalogConfig(value *AwsSagemakerFeatureGroup_DataCatalogConfigProperty) {
	if err := a.validatePutDataCatalogConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putDataCatalogConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) PutS3StorageConfig(value *AwsSagemakerFeatureGroup_S3StorageConfigProperty) {
	if err := a.validatePutS3StorageConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3StorageConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) ResetDataCatalogConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetDataCatalogConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) ResetDisableGlueTableCreation() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableGlueTableCreation",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) ResetTableFormat() {
	_jsii_.InvokeVoid(
		a,
		"resetTableFormat",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsSagemakerFeatureGroup_OfflineStoreConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

