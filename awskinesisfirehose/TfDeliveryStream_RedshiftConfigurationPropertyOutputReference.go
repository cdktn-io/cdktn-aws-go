package awskinesisfirehose

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisfirehose/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type TfDeliveryStream_RedshiftConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	CloudwatchLoggingOptions() TfDeliveryStream_RedshiftConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	// Experimental.
	CloudwatchLoggingOptionsInput() *TfDeliveryStream_RedshiftConfigurationCloudwatchLoggingOptionsProperty
	// Experimental.
	ClusterJdbcurl() *string
	// Experimental.
	SetClusterJdbcurl(val *string)
	// Experimental.
	ClusterJdbcurlInput() *string
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
	CopyOptions() *string
	// Experimental.
	SetCopyOptions(val *string)
	// Experimental.
	CopyOptionsInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	DataTableColumns() *string
	// Experimental.
	SetDataTableColumns(val *string)
	// Experimental.
	DataTableColumnsInput() *string
	// Experimental.
	DataTableName() *string
	// Experimental.
	SetDataTableName(val *string)
	// Experimental.
	DataTableNameInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *TfDeliveryStream_RedshiftConfigurationProperty
	// Experimental.
	SetInternalValue(val *TfDeliveryStream_RedshiftConfigurationProperty)
	// Experimental.
	Password() *string
	// Experimental.
	SetPassword(val *string)
	// Experimental.
	PasswordInput() *string
	// Experimental.
	ProcessingConfiguration() TfDeliveryStream_RedshiftConfigurationProcessingConfigurationPropertyOutputReference
	// Experimental.
	ProcessingConfigurationInput() *TfDeliveryStream_RedshiftConfigurationProcessingConfigurationProperty
	// Experimental.
	RetryDuration() *float64
	// Experimental.
	SetRetryDuration(val *float64)
	// Experimental.
	RetryDurationInput() *float64
	// Experimental.
	RoleArn() *string
	// Experimental.
	SetRoleArn(val *string)
	// Experimental.
	RoleArnInput() *string
	// Experimental.
	S3BackupConfiguration() TfDeliveryStream_RedshiftConfigurationS3BackupConfigurationPropertyOutputReference
	// Experimental.
	S3BackupConfigurationInput() *TfDeliveryStream_RedshiftConfigurationS3BackupConfigurationProperty
	// Experimental.
	S3BackupMode() *string
	// Experimental.
	SetS3BackupMode(val *string)
	// Experimental.
	S3BackupModeInput() *string
	// Experimental.
	S3Configuration() TfDeliveryStream_RedshiftConfigurationS3ConfigurationPropertyOutputReference
	// Experimental.
	S3ConfigurationInput() *TfDeliveryStream_RedshiftConfigurationS3ConfigurationProperty
	// Experimental.
	SecretsManagerConfiguration() TfDeliveryStream_RedshiftConfigurationSecretsManagerConfigurationPropertyOutputReference
	// Experimental.
	SecretsManagerConfigurationInput() *TfDeliveryStream_RedshiftConfigurationSecretsManagerConfigurationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	Username() *string
	// Experimental.
	SetUsername(val *string)
	// Experimental.
	UsernameInput() *string
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
	PutCloudwatchLoggingOptions(value *TfDeliveryStream_RedshiftConfigurationCloudwatchLoggingOptionsProperty)
	// Experimental.
	PutProcessingConfiguration(value *TfDeliveryStream_RedshiftConfigurationProcessingConfigurationProperty)
	// Experimental.
	PutS3BackupConfiguration(value *TfDeliveryStream_RedshiftConfigurationS3BackupConfigurationProperty)
	// Experimental.
	PutS3Configuration(value *TfDeliveryStream_RedshiftConfigurationS3ConfigurationProperty)
	// Experimental.
	PutSecretsManagerConfiguration(value *TfDeliveryStream_RedshiftConfigurationSecretsManagerConfigurationProperty)
	// Experimental.
	ResetCloudwatchLoggingOptions()
	// Experimental.
	ResetCopyOptions()
	// Experimental.
	ResetDataTableColumns()
	// Experimental.
	ResetPassword()
	// Experimental.
	ResetProcessingConfiguration()
	// Experimental.
	ResetRetryDuration()
	// Experimental.
	ResetS3BackupConfiguration()
	// Experimental.
	ResetS3BackupMode()
	// Experimental.
	ResetSecretsManagerConfiguration()
	// Experimental.
	ResetUsername()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TfDeliveryStream_RedshiftConfigurationPropertyOutputReference
type jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) CloudwatchLoggingOptions() TfDeliveryStream_RedshiftConfigurationCloudwatchLoggingOptionsPropertyOutputReference {
	var returns TfDeliveryStream_RedshiftConfigurationCloudwatchLoggingOptionsPropertyOutputReference
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) CloudwatchLoggingOptionsInput() *TfDeliveryStream_RedshiftConfigurationCloudwatchLoggingOptionsProperty {
	var returns *TfDeliveryStream_RedshiftConfigurationCloudwatchLoggingOptionsProperty
	_jsii_.Get(
		j,
		"cloudwatchLoggingOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ClusterJdbcurl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterJdbcurl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ClusterJdbcurlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clusterJdbcurlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) CopyOptions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyOptions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) CopyOptionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"copyOptionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) DataTableColumns() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTableColumns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) DataTableColumnsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTableColumnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) DataTableName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTableName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) DataTableNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataTableNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) InternalValue() *TfDeliveryStream_RedshiftConfigurationProperty {
	var returns *TfDeliveryStream_RedshiftConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ProcessingConfiguration() TfDeliveryStream_RedshiftConfigurationProcessingConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_RedshiftConfigurationProcessingConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"processingConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ProcessingConfigurationInput() *TfDeliveryStream_RedshiftConfigurationProcessingConfigurationProperty {
	var returns *TfDeliveryStream_RedshiftConfigurationProcessingConfigurationProperty
	_jsii_.Get(
		j,
		"processingConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) RetryDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) RetryDurationInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"retryDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) RoleArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) RoleArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) S3BackupConfiguration() TfDeliveryStream_RedshiftConfigurationS3BackupConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_RedshiftConfigurationS3BackupConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3BackupConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) S3BackupConfigurationInput() *TfDeliveryStream_RedshiftConfigurationS3BackupConfigurationProperty {
	var returns *TfDeliveryStream_RedshiftConfigurationS3BackupConfigurationProperty
	_jsii_.Get(
		j,
		"s3BackupConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) S3BackupMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) S3BackupModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3BackupModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) S3Configuration() TfDeliveryStream_RedshiftConfigurationS3ConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_RedshiftConfigurationS3ConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Configuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) S3ConfigurationInput() *TfDeliveryStream_RedshiftConfigurationS3ConfigurationProperty {
	var returns *TfDeliveryStream_RedshiftConfigurationS3ConfigurationProperty
	_jsii_.Get(
		j,
		"s3ConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) SecretsManagerConfiguration() TfDeliveryStream_RedshiftConfigurationSecretsManagerConfigurationPropertyOutputReference {
	var returns TfDeliveryStream_RedshiftConfigurationSecretsManagerConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"secretsManagerConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) SecretsManagerConfigurationInput() *TfDeliveryStream_RedshiftConfigurationSecretsManagerConfigurationProperty {
	var returns *TfDeliveryStream_RedshiftConfigurationSecretsManagerConfigurationProperty
	_jsii_.Get(
		j,
		"secretsManagerConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewTfDeliveryStream_RedshiftConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) TfDeliveryStream_RedshiftConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewTfDeliveryStream_RedshiftConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.RedshiftConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewTfDeliveryStream_RedshiftConfigurationPropertyOutputReference_Override(t TfDeliveryStream_RedshiftConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-firehose.TfDeliveryStream.RedshiftConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetClusterJdbcurl(val *string) {
	if err := j.validateSetClusterJdbcurlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterJdbcurl",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetCopyOptions(val *string) {
	if err := j.validateSetCopyOptionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"copyOptions",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetDataTableColumns(val *string) {
	if err := j.validateSetDataTableColumnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTableColumns",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetDataTableName(val *string) {
	if err := j.validateSetDataTableNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataTableName",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetInternalValue(val *TfDeliveryStream_RedshiftConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetRetryDuration(val *float64) {
	if err := j.validateSetRetryDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"retryDuration",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetRoleArn(val *string) {
	if err := j.validateSetRoleArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArn",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetS3BackupMode(val *string) {
	if err := j.validateSetS3BackupModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3BackupMode",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) PutCloudwatchLoggingOptions(value *TfDeliveryStream_RedshiftConfigurationCloudwatchLoggingOptionsProperty) {
	if err := t.validatePutCloudwatchLoggingOptionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putCloudwatchLoggingOptions",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) PutProcessingConfiguration(value *TfDeliveryStream_RedshiftConfigurationProcessingConfigurationProperty) {
	if err := t.validatePutProcessingConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putProcessingConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) PutS3BackupConfiguration(value *TfDeliveryStream_RedshiftConfigurationS3BackupConfigurationProperty) {
	if err := t.validatePutS3BackupConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3BackupConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) PutS3Configuration(value *TfDeliveryStream_RedshiftConfigurationS3ConfigurationProperty) {
	if err := t.validatePutS3ConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putS3Configuration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) PutSecretsManagerConfiguration(value *TfDeliveryStream_RedshiftConfigurationSecretsManagerConfigurationProperty) {
	if err := t.validatePutSecretsManagerConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"putSecretsManagerConfiguration",
		[]interface{}{value},
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ResetCloudwatchLoggingOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetCloudwatchLoggingOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ResetCopyOptions() {
	_jsii_.InvokeVoid(
		t,
		"resetCopyOptions",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ResetDataTableColumns() {
	_jsii_.InvokeVoid(
		t,
		"resetDataTableColumns",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		t,
		"resetPassword",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ResetProcessingConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetProcessingConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ResetRetryDuration() {
	_jsii_.InvokeVoid(
		t,
		"resetRetryDuration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ResetS3BackupConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetS3BackupConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ResetS3BackupMode() {
	_jsii_.InvokeVoid(
		t,
		"resetS3BackupMode",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ResetSecretsManagerConfiguration() {
	_jsii_.InvokeVoid(
		t,
		"resetSecretsManagerConfiguration",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		t,
		"resetUsername",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (t *jsiiProxy_TfDeliveryStream_RedshiftConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

