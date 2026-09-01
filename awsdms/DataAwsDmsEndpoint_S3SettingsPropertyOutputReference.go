package awsdms

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awsdms/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awsdms/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type DataAwsDmsEndpoint_S3SettingsPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	AddColumnName() cdktn.IResolvable
	// Experimental.
	BucketFolder() *string
	// Experimental.
	BucketName() *string
	// Experimental.
	CannedAclForObjects() *string
	// Experimental.
	CdcInsertsAndUpdates() cdktn.IResolvable
	// Experimental.
	CdcInsertsOnly() cdktn.IResolvable
	// Experimental.
	CdcMaxBatchInterval() *float64
	// Experimental.
	CdcMinFileSize() *float64
	// Experimental.
	CdcPath() *string
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
	CompressionType() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	CsvDelimiter() *string
	// Experimental.
	CsvNoSupValue() *string
	// Experimental.
	CsvNullValue() *string
	// Experimental.
	CsvRowDelimiter() *string
	// Experimental.
	DataFormat() *string
	// Experimental.
	DataPageSize() *float64
	// Experimental.
	DatePartitionDelimiter() *string
	// Experimental.
	DatePartitionEnabled() cdktn.IResolvable
	// Experimental.
	DatePartitionSequence() *string
	// Experimental.
	DictPageSizeLimit() *float64
	// Experimental.
	EnableStatistics() cdktn.IResolvable
	// Experimental.
	EncodingType() *string
	// Experimental.
	EncryptionMode() *string
	// Experimental.
	ExternalTableDefinition() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	GlueCatalogGeneration() cdktn.IResolvable
	// Experimental.
	IgnoreHeaderRows() *float64
	// Experimental.
	IgnoreHeadersRow() *float64
	// Experimental.
	IncludeOpForFullLoad() cdktn.IResolvable
	// Experimental.
	InternalValue() *DataAwsDmsEndpoint_S3SettingsProperty
	// Experimental.
	SetInternalValue(val *DataAwsDmsEndpoint_S3SettingsProperty)
	// Experimental.
	MaxFileSize() *float64
	// Experimental.
	ParquetTimestampInMillisecond() cdktn.IResolvable
	// Experimental.
	ParquetVersion() *string
	// Experimental.
	PreserveTransactions() cdktn.IResolvable
	// Experimental.
	Rfc4180() cdktn.IResolvable
	// Experimental.
	RowGroupLength() *float64
	// Experimental.
	ServerSideEncryptionKmsKeyId() *string
	// Experimental.
	ServiceAccessRoleArn() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	TimestampColumnName() *string
	// Experimental.
	UseCsvNoSupValue() cdktn.IResolvable
	// Experimental.
	UseTaskStartTimeForFullLoadTimestamp() cdktn.IResolvable
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAwsDmsEndpoint_S3SettingsPropertyOutputReference
type jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) AddColumnName() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"addColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) BucketFolder() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketFolder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) BucketName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bucketName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) CannedAclForObjects() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cannedAclForObjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) CdcInsertsAndUpdates() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"cdcInsertsAndUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) CdcInsertsOnly() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"cdcInsertsOnly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) CdcMaxBatchInterval() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMaxBatchInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) CdcMinFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"cdcMinFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) CdcPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cdcPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) CompressionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"compressionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) CsvDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) CsvNoSupValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNoSupValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) CsvNullValue() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvNullValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) CsvRowDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"csvRowDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) DataFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) DataPageSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataPageSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) DatePartitionDelimiter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionDelimiter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) DatePartitionEnabled() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"datePartitionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) DatePartitionSequence() *string {
	var returns *string
	_jsii_.Get(
		j,
		"datePartitionSequence",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) DictPageSizeLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dictPageSizeLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) EnableStatistics() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"enableStatistics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) EncodingType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encodingType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) EncryptionMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) ExternalTableDefinition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"externalTableDefinition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) GlueCatalogGeneration() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"glueCatalogGeneration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) IgnoreHeaderRows() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreHeaderRows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) IgnoreHeadersRow() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ignoreHeadersRow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) IncludeOpForFullLoad() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"includeOpForFullLoad",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) InternalValue() *DataAwsDmsEndpoint_S3SettingsProperty {
	var returns *DataAwsDmsEndpoint_S3SettingsProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) MaxFileSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxFileSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) ParquetTimestampInMillisecond() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"parquetTimestampInMillisecond",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) ParquetVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"parquetVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) PreserveTransactions() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"preserveTransactions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) Rfc4180() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"rfc4180",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) RowGroupLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rowGroupLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) ServerSideEncryptionKmsKeyId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serverSideEncryptionKmsKeyId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) ServiceAccessRoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccessRoleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) TimestampColumnName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timestampColumnName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) UseCsvNoSupValue() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"useCsvNoSupValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) UseTaskStartTimeForFullLoadTimestamp() cdktn.IResolvable {
	var returns cdktn.IResolvable
	_jsii_.Get(
		j,
		"useTaskStartTimeForFullLoadTimestamp",
		&returns,
	)
	return returns
}


// Experimental.
func NewDataAwsDmsEndpoint_S3SettingsPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAwsDmsEndpoint_S3SettingsPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewDataAwsDmsEndpoint_S3SettingsPropertyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-dms.DataAwsDmsEndpoint.S3SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

// Experimental.
func NewDataAwsDmsEndpoint_S3SettingsPropertyOutputReference_Override(d DataAwsDmsEndpoint_S3SettingsPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-dms.DataAwsDmsEndpoint.S3SettingsPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference)SetInternalValue(val *DataAwsDmsEndpoint_S3SettingsProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAwsDmsEndpoint_S3SettingsPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

