package datasync

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktn-io/cdktn-aws-go/datasync/jsii"

	"github.com/cdktn-io/cdktn-aws-go/datasync/internal"
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

// Experimental.
type AwsTask_TaskReportConfigPropertyOutputReference interface {
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
	InternalValue() *AwsTask_TaskReportConfigProperty
	// Experimental.
	SetInternalValue(val *AwsTask_TaskReportConfigProperty)
	// Experimental.
	OutputType() *string
	// Experimental.
	SetOutputType(val *string)
	// Experimental.
	OutputTypeInput() *string
	// Experimental.
	ReportLevel() *string
	// Experimental.
	SetReportLevel(val *string)
	// Experimental.
	ReportLevelInput() *string
	// Experimental.
	ReportOverrides() AwsTask_ReportOverridesPropertyOutputReference
	// Experimental.
	ReportOverridesInput() *AwsTask_ReportOverridesProperty
	// Experimental.
	S3Destination() AwsTask_S3DestinationPropertyOutputReference
	// Experimental.
	S3DestinationInput() *AwsTask_S3DestinationProperty
	// Experimental.
	S3ObjectVersioning() *string
	// Experimental.
	SetS3ObjectVersioning(val *string)
	// Experimental.
	S3ObjectVersioningInput() *string
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
	PutReportOverrides(value *AwsTask_ReportOverridesProperty)
	// Experimental.
	PutS3Destination(value *AwsTask_S3DestinationProperty)
	// Experimental.
	ResetOutputType()
	// Experimental.
	ResetReportLevel()
	// Experimental.
	ResetReportOverrides()
	// Experimental.
	ResetS3ObjectVersioning()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AwsTask_TaskReportConfigPropertyOutputReference
type jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) InternalValue() *AwsTask_TaskReportConfigProperty {
	var returns *AwsTask_TaskReportConfigProperty
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) OutputType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) OutputTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"outputTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) ReportLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) ReportLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reportLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) ReportOverrides() AwsTask_ReportOverridesPropertyOutputReference {
	var returns AwsTask_ReportOverridesPropertyOutputReference
	_jsii_.Get(
		j,
		"reportOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) ReportOverridesInput() *AwsTask_ReportOverridesProperty {
	var returns *AwsTask_ReportOverridesProperty
	_jsii_.Get(
		j,
		"reportOverridesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) S3Destination() AwsTask_S3DestinationPropertyOutputReference {
	var returns AwsTask_S3DestinationPropertyOutputReference
	_jsii_.Get(
		j,
		"s3Destination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) S3DestinationInput() *AwsTask_S3DestinationProperty {
	var returns *AwsTask_S3DestinationProperty
	_jsii_.Get(
		j,
		"s3DestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) S3ObjectVersioning() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersioning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) S3ObjectVersioningInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"s3ObjectVersioningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


// Experimental.
func NewAwsTask_TaskReportConfigPropertyOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) AwsTask_TaskReportConfigPropertyOutputReference {
	_init_.Initialize()

	if err := validateNewAwsTask_TaskReportConfigPropertyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference{}

	_jsii_.Create(
		"@cdktn/aws-datasync.AwsTask.TaskReportConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

// Experimental.
func NewAwsTask_TaskReportConfigPropertyOutputReference_Override(a AwsTask_TaskReportConfigPropertyOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/aws-datasync.AwsTask.TaskReportConfigPropertyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference)SetInternalValue(val *AwsTask_TaskReportConfigProperty) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference)SetOutputType(val *string) {
	if err := j.validateSetOutputTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"outputType",
		val,
	)
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference)SetReportLevel(val *string) {
	if err := j.validateSetReportLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reportLevel",
		val,
	)
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference)SetS3ObjectVersioning(val *string) {
	if err := j.validateSetS3ObjectVersioningParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"s3ObjectVersioning",
		val,
	)
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) PutReportOverrides(value *AwsTask_ReportOverridesProperty) {
	if err := a.validatePutReportOverridesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putReportOverrides",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) PutS3Destination(value *AwsTask_S3DestinationProperty) {
	if err := a.validatePutS3DestinationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putS3Destination",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) ResetOutputType() {
	_jsii_.InvokeVoid(
		a,
		"resetOutputType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) ResetReportLevel() {
	_jsii_.InvokeVoid(
		a,
		"resetReportLevel",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) ResetReportOverrides() {
	_jsii_.InvokeVoid(
		a,
		"resetReportOverrides",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) ResetS3ObjectVersioning() {
	_jsii_.InvokeVoid(
		a,
		"resetS3ObjectVersioning",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (a *jsiiProxy_AwsTask_TaskReportConfigPropertyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

