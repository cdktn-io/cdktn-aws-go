package awskinesisanalyticsv2

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/jsii"

	"github.com/cdktn-io/cdktn-aws-go/awskinesisanalyticsv2/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference interface {
	cdktn.ComplexObject
	// Experimental.
	ApplicationCodeConfiguration() AwsKinesisanalyticsv2Application_ApplicationCodeConfigurationPropertyOutputReference
	// Experimental.
	ApplicationCodeConfigurationInput() *AwsKinesisanalyticsv2Application_ApplicationCodeConfigurationProperty
	// Experimental.
	ApplicationEncryptionConfiguration() AwsKinesisanalyticsv2Application_ApplicationEncryptionConfigurationPropertyOutputReference
	// Experimental.
	ApplicationEncryptionConfigurationInput() *AwsKinesisanalyticsv2Application_ApplicationEncryptionConfigurationProperty
	// Experimental.
	ApplicationSnapshotConfiguration() AwsKinesisanalyticsv2Application_ApplicationSnapshotConfigurationPropertyOutputReference
	// Experimental.
	ApplicationSnapshotConfigurationInput() *AwsKinesisanalyticsv2Application_ApplicationSnapshotConfigurationProperty
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
	EnvironmentProperties() AwsKinesisanalyticsv2Application_EnvironmentPropertiesPropertyOutputReference
	// Experimental.
	EnvironmentPropertiesInput() *AwsKinesisanalyticsv2Application_EnvironmentPropertiesProperty
	// Experimental.
	FlinkApplicationConfiguration() AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference
	// Experimental.
	FlinkApplicationConfigurationInput() *AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationProperty
	// Experimental.
	Fqn() *string
	// Experimental.
	InternalValue() *AwsKinesisanalyticsv2Application_ApplicationConfigurationProperty
	// Experimental.
	SetInternalValue(val *AwsKinesisanalyticsv2Application_ApplicationConfigurationProperty)
	// Experimental.
	RunConfiguration() AwsKinesisanalyticsv2Application_RunConfigurationPropertyOutputReference
	// Experimental.
	RunConfigurationInput() *AwsKinesisanalyticsv2Application_RunConfigurationProperty
	// Experimental.
	SqlApplicationConfiguration() AwsKinesisanalyticsv2Application_SqlApplicationConfigurationPropertyOutputReference
	// Experimental.
	SqlApplicationConfigurationInput() *AwsKinesisanalyticsv2Application_SqlApplicationConfigurationProperty
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	VpcConfiguration() AwsKinesisanalyticsv2Application_VpcConfigurationPropertyOutputReference
	// Experimental.
	VpcConfigurationInput() *AwsKinesisanalyticsv2Application_VpcConfigurationProperty
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
	PutApplicationCodeConfiguration(value *AwsKinesisanalyticsv2Application_ApplicationCodeConfigurationProperty)
	// Experimental.
	PutApplicationEncryptionConfiguration(value *AwsKinesisanalyticsv2Application_ApplicationEncryptionConfigurationProperty)
	// Experimental.
	PutApplicationSnapshotConfiguration(value *AwsKinesisanalyticsv2Application_ApplicationSnapshotConfigurationProperty)
	// Experimental.
	PutEnvironmentProperties(value *AwsKinesisanalyticsv2Application_EnvironmentPropertiesProperty)
	// Experimental.
	PutFlinkApplicationConfiguration(value *AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationProperty)
	// Experimental.
	PutRunConfiguration(value *AwsKinesisanalyticsv2Application_RunConfigurationProperty)
	// Experimental.
	PutSqlApplicationConfiguration(value *AwsKinesisanalyticsv2Application_SqlApplicationConfigurationProperty)
	// Experimental.
	PutVpcConfiguration(value *AwsKinesisanalyticsv2Application_VpcConfigurationProperty)
	// Experimental.
	ResetApplicationEncryptionConfiguration()
	// Experimental.
	ResetApplicationSnapshotConfiguration()
	// Experimental.
	ResetEnvironmentProperties()
	// Experimental.
	ResetFlinkApplicationConfiguration()
	// Experimental.
	ResetRunConfiguration()
	// Experimental.
	ResetSqlApplicationConfiguration()
	// Experimental.
	ResetVpcConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference
type jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ApplicationCodeConfiguration() AwsKinesisanalyticsv2Application_ApplicationCodeConfigurationPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_ApplicationCodeConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"applicationCodeConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ApplicationCodeConfigurationInput() *AwsKinesisanalyticsv2Application_ApplicationCodeConfigurationProperty {
	var returns *AwsKinesisanalyticsv2Application_ApplicationCodeConfigurationProperty
	_jsii_.Get(
		j,
		"applicationCodeConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ApplicationEncryptionConfiguration() AwsKinesisanalyticsv2Application_ApplicationEncryptionConfigurationPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_ApplicationEncryptionConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"applicationEncryptionConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ApplicationEncryptionConfigurationInput() *AwsKinesisanalyticsv2Application_ApplicationEncryptionConfigurationProperty {
	var returns *AwsKinesisanalyticsv2Application_ApplicationEncryptionConfigurationProperty
	_jsii_.Get(
		j,
		"applicationEncryptionConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ApplicationSnapshotConfiguration() AwsKinesisanalyticsv2Application_ApplicationSnapshotConfigurationPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_ApplicationSnapshotConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"applicationSnapshotConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ApplicationSnapshotConfigurationInput() *AwsKinesisanalyticsv2Application_ApplicationSnapshotConfigurationProperty {
	var returns *AwsKinesisanalyticsv2Application_ApplicationSnapshotConfigurationProperty
	_jsii_.Get(
		j,
		"applicationSnapshotConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) EnvironmentProperties() AwsKinesisanalyticsv2Application_EnvironmentPropertiesPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_EnvironmentPropertiesPropertyOutputReference
	_jsii_.Get(
		j,
		"environmentProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) EnvironmentPropertiesInput() *AwsKinesisanalyticsv2Application_EnvironmentPropertiesProperty {
	var returns *AwsKinesisanalyticsv2Application_EnvironmentPropertiesProperty
	_jsii_.Get(
		j,
		"environmentPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) FlinkApplicationConfiguration() AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"flinkApplicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) FlinkApplicationConfigurationInput() *AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationProperty {
	var returns *AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationProperty
	_jsii_.Get(
		j,
		"flinkApplicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) InternalValue() *AwsKinesisanalyticsv2Application_ApplicationConfigurationProperty {
	var returns *AwsKinesisanalyticsv2Application_ApplicationConfigurationProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) RunConfiguration() AwsKinesisanalyticsv2Application_RunConfigurationPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_RunConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"runConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) RunConfigurationInput() *AwsKinesisanalyticsv2Application_RunConfigurationProperty {
	var returns *AwsKinesisanalyticsv2Application_RunConfigurationProperty
	_jsii_.Get(
		j,
		"runConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) SqlApplicationConfiguration() AwsKinesisanalyticsv2Application_SqlApplicationConfigurationPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_SqlApplicationConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"sqlApplicationConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) SqlApplicationConfigurationInput() *AwsKinesisanalyticsv2Application_SqlApplicationConfigurationProperty {
	var returns *AwsKinesisanalyticsv2Application_SqlApplicationConfigurationProperty
	_jsii_.Get(
		j,
		"sqlApplicationConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) VpcConfiguration() AwsKinesisanalyticsv2Application_VpcConfigurationPropertyOutputReference {
	var returns AwsKinesisanalyticsv2Application_VpcConfigurationPropertyOutputReference
	_jsii_.Get(
		j,
		"vpcConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) VpcConfigurationInput() *AwsKinesisanalyticsv2Application_VpcConfigurationProperty {
	var returns *AwsKinesisanalyticsv2Application_VpcConfigurationProperty
	_jsii_.Get(
		j,
		"vpcConfigurationInput",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsKinesisanalyticsv2Application.ApplicationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference_Override(a AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-kinesis-analytics-v2.AwsKinesisanalyticsv2Application.ApplicationConfigurationPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference)SetInternalValue(val *AwsKinesisanalyticsv2Application_ApplicationConfigurationProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) PutApplicationCodeConfiguration(value *AwsKinesisanalyticsv2Application_ApplicationCodeConfigurationProperty) {
	if err := a.validatePutApplicationCodeConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApplicationCodeConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) PutApplicationEncryptionConfiguration(value *AwsKinesisanalyticsv2Application_ApplicationEncryptionConfigurationProperty) {
	if err := a.validatePutApplicationEncryptionConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApplicationEncryptionConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) PutApplicationSnapshotConfiguration(value *AwsKinesisanalyticsv2Application_ApplicationSnapshotConfigurationProperty) {
	if err := a.validatePutApplicationSnapshotConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putApplicationSnapshotConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) PutEnvironmentProperties(value *AwsKinesisanalyticsv2Application_EnvironmentPropertiesProperty) {
	if err := a.validatePutEnvironmentPropertiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnvironmentProperties",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) PutFlinkApplicationConfiguration(value *AwsKinesisanalyticsv2Application_FlinkApplicationConfigurationProperty) {
	if err := a.validatePutFlinkApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFlinkApplicationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) PutRunConfiguration(value *AwsKinesisanalyticsv2Application_RunConfigurationProperty) {
	if err := a.validatePutRunConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRunConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) PutSqlApplicationConfiguration(value *AwsKinesisanalyticsv2Application_SqlApplicationConfigurationProperty) {
	if err := a.validatePutSqlApplicationConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSqlApplicationConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) PutVpcConfiguration(value *AwsKinesisanalyticsv2Application_VpcConfigurationProperty) {
	if err := a.validatePutVpcConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putVpcConfiguration",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ResetApplicationEncryptionConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationEncryptionConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ResetApplicationSnapshotConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationSnapshotConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ResetEnvironmentProperties() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironmentProperties",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ResetFlinkApplicationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetFlinkApplicationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ResetRunConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetRunConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ResetSqlApplicationConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetSqlApplicationConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ResetVpcConfiguration() {
	_jsii_.InvokeVoid(
		a,
		"resetVpcConfiguration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsKinesisanalyticsv2Application_ApplicationConfigurationPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

