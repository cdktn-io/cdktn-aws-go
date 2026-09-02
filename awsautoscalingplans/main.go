// cdk-terrain bindings for the Auto Scaling Plans group of terraform-provider-aws 6.62.0
package awsautoscalingplans

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan",
		reflect.TypeOf((*TfScalingPlan)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "applicationSource", GoGetter: "ApplicationSource"},
			_jsii_.MemberProperty{JsiiProperty: "applicationSourceInput", GoGetter: "ApplicationSourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "markWriteOnlyAttribute", GoMethod: "MarkWriteOnlyAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putApplicationSource", GoMethod: "PutApplicationSource"},
			_jsii_.MemberMethod{JsiiMethod: "putScalingInstruction", GoMethod: "PutScalingInstruction"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "region", GoGetter: "Region"},
			_jsii_.MemberProperty{JsiiProperty: "regionInput", GoGetter: "RegionInput"},
			_jsii_.MemberMethod{JsiiMethod: "registerProviderFeatureUsage", GoMethod: "RegisterProviderFeatureUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegion", GoMethod: "ResetRegion"},
			_jsii_.MemberProperty{JsiiProperty: "scalingInstruction", GoGetter: "ScalingInstruction"},
			_jsii_.MemberProperty{JsiiProperty: "scalingInstructionInput", GoGetter: "ScalingInstructionInput"},
			_jsii_.MemberProperty{JsiiProperty: "scalingPlanVersion", GoGetter: "ScalingPlanVersion"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "with", GoMethod: "With"},
		},
		func() interface{} {
			j := jsiiProxy_TfScalingPlan{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.ApplicationSourceProperty",
		reflect.TypeOf((*TfScalingPlan_ApplicationSourceProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.ApplicationSourcePropertyOutputReference",
		reflect.TypeOf((*TfScalingPlan_ApplicationSourcePropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cloudformationStackArn", GoGetter: "CloudformationStackArn"},
			_jsii_.MemberProperty{JsiiProperty: "cloudformationStackArnInput", GoGetter: "CloudformationStackArnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putTagFilter", GoMethod: "PutTagFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resetCloudformationStackArn", GoMethod: "ResetCloudformationStackArn"},
			_jsii_.MemberMethod{JsiiMethod: "resetTagFilter", GoMethod: "ResetTagFilter"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "tagFilter", GoGetter: "TagFilter"},
			_jsii_.MemberProperty{JsiiProperty: "tagFilterInput", GoGetter: "TagFilterInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TfScalingPlan_ApplicationSourcePropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.CustomizedLoadMetricSpecificationProperty",
		reflect.TypeOf((*TfScalingPlan_CustomizedLoadMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.CustomizedLoadMetricSpecificationPropertyOutputReference",
		reflect.TypeOf((*TfScalingPlan_CustomizedLoadMetricSpecificationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dimensions", GoGetter: "Dimensions"},
			_jsii_.MemberProperty{JsiiProperty: "dimensionsInput", GoGetter: "DimensionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "metricName", GoGetter: "MetricName"},
			_jsii_.MemberProperty{JsiiProperty: "metricNameInput", GoGetter: "MetricNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDimensions", GoMethod: "ResetDimensions"},
			_jsii_.MemberMethod{JsiiMethod: "resetUnit", GoMethod: "ResetUnit"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberProperty{JsiiProperty: "statisticInput", GoGetter: "StatisticInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
		},
		func() interface{} {
			j := jsiiProxy_TfScalingPlan_CustomizedLoadMetricSpecificationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.CustomizedScalingMetricSpecificationProperty",
		reflect.TypeOf((*TfScalingPlan_CustomizedScalingMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.CustomizedScalingMetricSpecificationPropertyOutputReference",
		reflect.TypeOf((*TfScalingPlan_CustomizedScalingMetricSpecificationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dimensions", GoGetter: "Dimensions"},
			_jsii_.MemberProperty{JsiiProperty: "dimensionsInput", GoGetter: "DimensionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "metricName", GoGetter: "MetricName"},
			_jsii_.MemberProperty{JsiiProperty: "metricNameInput", GoGetter: "MetricNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDimensions", GoMethod: "ResetDimensions"},
			_jsii_.MemberMethod{JsiiMethod: "resetUnit", GoMethod: "ResetUnit"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "statistic", GoGetter: "Statistic"},
			_jsii_.MemberProperty{JsiiProperty: "statisticInput", GoGetter: "StatisticInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unit", GoGetter: "Unit"},
			_jsii_.MemberProperty{JsiiProperty: "unitInput", GoGetter: "UnitInput"},
		},
		func() interface{} {
			j := jsiiProxy_TfScalingPlan_CustomizedScalingMetricSpecificationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.PredefinedLoadMetricSpecificationProperty",
		reflect.TypeOf((*TfScalingPlan_PredefinedLoadMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.PredefinedLoadMetricSpecificationPropertyOutputReference",
		reflect.TypeOf((*TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "predefinedLoadMetricType", GoGetter: "PredefinedLoadMetricType"},
			_jsii_.MemberProperty{JsiiProperty: "predefinedLoadMetricTypeInput", GoGetter: "PredefinedLoadMetricTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceLabel", GoMethod: "ResetResourceLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resourceLabel", GoGetter: "ResourceLabel"},
			_jsii_.MemberProperty{JsiiProperty: "resourceLabelInput", GoGetter: "ResourceLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TfScalingPlan_PredefinedLoadMetricSpecificationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.PredefinedScalingMetricSpecificationProperty",
		reflect.TypeOf((*TfScalingPlan_PredefinedScalingMetricSpecificationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.PredefinedScalingMetricSpecificationPropertyOutputReference",
		reflect.TypeOf((*TfScalingPlan_PredefinedScalingMetricSpecificationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "predefinedScalingMetricType", GoGetter: "PredefinedScalingMetricType"},
			_jsii_.MemberProperty{JsiiProperty: "predefinedScalingMetricTypeInput", GoGetter: "PredefinedScalingMetricTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetResourceLabel", GoMethod: "ResetResourceLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resourceLabel", GoGetter: "ResourceLabel"},
			_jsii_.MemberProperty{JsiiProperty: "resourceLabelInput", GoGetter: "ResourceLabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TfScalingPlan_PredefinedScalingMetricSpecificationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.ScalingInstructionProperty",
		reflect.TypeOf((*TfScalingPlan_ScalingInstructionProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.ScalingInstructionPropertyList",
		reflect.TypeOf((*TfScalingPlan_ScalingInstructionPropertyList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_TfScalingPlan_ScalingInstructionPropertyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.ScalingInstructionPropertyOutputReference",
		reflect.TypeOf((*TfScalingPlan_ScalingInstructionPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customizedLoadMetricSpecification", GoGetter: "CustomizedLoadMetricSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "customizedLoadMetricSpecificationInput", GoGetter: "CustomizedLoadMetricSpecificationInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableDynamicScaling", GoGetter: "DisableDynamicScaling"},
			_jsii_.MemberProperty{JsiiProperty: "disableDynamicScalingInput", GoGetter: "DisableDynamicScalingInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxCapacity", GoGetter: "MaxCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "maxCapacityInput", GoGetter: "MaxCapacityInput"},
			_jsii_.MemberProperty{JsiiProperty: "minCapacity", GoGetter: "MinCapacity"},
			_jsii_.MemberProperty{JsiiProperty: "minCapacityInput", GoGetter: "MinCapacityInput"},
			_jsii_.MemberProperty{JsiiProperty: "predefinedLoadMetricSpecification", GoGetter: "PredefinedLoadMetricSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "predefinedLoadMetricSpecificationInput", GoGetter: "PredefinedLoadMetricSpecificationInput"},
			_jsii_.MemberProperty{JsiiProperty: "predictiveScalingMaxCapacityBehavior", GoGetter: "PredictiveScalingMaxCapacityBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "predictiveScalingMaxCapacityBehaviorInput", GoGetter: "PredictiveScalingMaxCapacityBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "predictiveScalingMaxCapacityBuffer", GoGetter: "PredictiveScalingMaxCapacityBuffer"},
			_jsii_.MemberProperty{JsiiProperty: "predictiveScalingMaxCapacityBufferInput", GoGetter: "PredictiveScalingMaxCapacityBufferInput"},
			_jsii_.MemberProperty{JsiiProperty: "predictiveScalingMode", GoGetter: "PredictiveScalingMode"},
			_jsii_.MemberProperty{JsiiProperty: "predictiveScalingModeInput", GoGetter: "PredictiveScalingModeInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomizedLoadMetricSpecification", GoMethod: "PutCustomizedLoadMetricSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "putPredefinedLoadMetricSpecification", GoMethod: "PutPredefinedLoadMetricSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "putTargetTrackingConfiguration", GoMethod: "PutTargetTrackingConfiguration"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomizedLoadMetricSpecification", GoMethod: "ResetCustomizedLoadMetricSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableDynamicScaling", GoMethod: "ResetDisableDynamicScaling"},
			_jsii_.MemberMethod{JsiiMethod: "resetPredefinedLoadMetricSpecification", GoMethod: "ResetPredefinedLoadMetricSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetPredictiveScalingMaxCapacityBehavior", GoMethod: "ResetPredictiveScalingMaxCapacityBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetPredictiveScalingMaxCapacityBuffer", GoMethod: "ResetPredictiveScalingMaxCapacityBuffer"},
			_jsii_.MemberMethod{JsiiMethod: "resetPredictiveScalingMode", GoMethod: "ResetPredictiveScalingMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetScalingPolicyUpdateBehavior", GoMethod: "ResetScalingPolicyUpdateBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduledActionBufferTime", GoMethod: "ResetScheduledActionBufferTime"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resourceId", GoGetter: "ResourceId"},
			_jsii_.MemberProperty{JsiiProperty: "resourceIdInput", GoGetter: "ResourceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "scalableDimension", GoGetter: "ScalableDimension"},
			_jsii_.MemberProperty{JsiiProperty: "scalableDimensionInput", GoGetter: "ScalableDimensionInput"},
			_jsii_.MemberProperty{JsiiProperty: "scalingPolicyUpdateBehavior", GoGetter: "ScalingPolicyUpdateBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "scalingPolicyUpdateBehaviorInput", GoGetter: "ScalingPolicyUpdateBehaviorInput"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledActionBufferTime", GoGetter: "ScheduledActionBufferTime"},
			_jsii_.MemberProperty{JsiiProperty: "scheduledActionBufferTimeInput", GoGetter: "ScheduledActionBufferTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNamespace", GoGetter: "ServiceNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "serviceNamespaceInput", GoGetter: "ServiceNamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetTrackingConfiguration", GoGetter: "TargetTrackingConfiguration"},
			_jsii_.MemberProperty{JsiiProperty: "targetTrackingConfigurationInput", GoGetter: "TargetTrackingConfigurationInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TfScalingPlan_ScalingInstructionPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.TagFilterProperty",
		reflect.TypeOf((*TfScalingPlan_TagFilterProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.TagFilterPropertyList",
		reflect.TypeOf((*TfScalingPlan_TagFilterPropertyList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_TfScalingPlan_TagFilterPropertyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.TagFilterPropertyOutputReference",
		reflect.TypeOf((*TfScalingPlan_TagFilterPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetValues", GoMethod: "ResetValues"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
			_jsii_.MemberProperty{JsiiProperty: "valuesInput", GoGetter: "ValuesInput"},
		},
		func() interface{} {
			j := jsiiProxy_TfScalingPlan_TagFilterPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.TargetTrackingConfigurationProperty",
		reflect.TypeOf((*TfScalingPlan_TargetTrackingConfigurationProperty)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.TargetTrackingConfigurationPropertyList",
		reflect.TypeOf((*TfScalingPlan_TargetTrackingConfigurationPropertyList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_TfScalingPlan_TargetTrackingConfigurationPropertyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlan.TargetTrackingConfigurationPropertyOutputReference",
		reflect.TypeOf((*TfScalingPlan_TargetTrackingConfigurationPropertyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "customizedScalingMetricSpecification", GoGetter: "CustomizedScalingMetricSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "customizedScalingMetricSpecificationInput", GoGetter: "CustomizedScalingMetricSpecificationInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableScaleIn", GoGetter: "DisableScaleIn"},
			_jsii_.MemberProperty{JsiiProperty: "disableScaleInInput", GoGetter: "DisableScaleInInput"},
			_jsii_.MemberProperty{JsiiProperty: "estimatedInstanceWarmup", GoGetter: "EstimatedInstanceWarmup"},
			_jsii_.MemberProperty{JsiiProperty: "estimatedInstanceWarmupInput", GoGetter: "EstimatedInstanceWarmupInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "predefinedScalingMetricSpecification", GoGetter: "PredefinedScalingMetricSpecification"},
			_jsii_.MemberProperty{JsiiProperty: "predefinedScalingMetricSpecificationInput", GoGetter: "PredefinedScalingMetricSpecificationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putCustomizedScalingMetricSpecification", GoMethod: "PutCustomizedScalingMetricSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "putPredefinedScalingMetricSpecification", GoMethod: "PutPredefinedScalingMetricSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomizedScalingMetricSpecification", GoMethod: "ResetCustomizedScalingMetricSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableScaleIn", GoMethod: "ResetDisableScaleIn"},
			_jsii_.MemberMethod{JsiiMethod: "resetEstimatedInstanceWarmup", GoMethod: "ResetEstimatedInstanceWarmup"},
			_jsii_.MemberMethod{JsiiMethod: "resetPredefinedScalingMetricSpecification", GoMethod: "ResetPredefinedScalingMetricSpecification"},
			_jsii_.MemberMethod{JsiiMethod: "resetScaleInCooldown", GoMethod: "ResetScaleInCooldown"},
			_jsii_.MemberMethod{JsiiMethod: "resetScaleOutCooldown", GoMethod: "ResetScaleOutCooldown"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scaleInCooldown", GoGetter: "ScaleInCooldown"},
			_jsii_.MemberProperty{JsiiProperty: "scaleInCooldownInput", GoGetter: "ScaleInCooldownInput"},
			_jsii_.MemberProperty{JsiiProperty: "scaleOutCooldown", GoGetter: "ScaleOutCooldown"},
			_jsii_.MemberProperty{JsiiProperty: "scaleOutCooldownInput", GoGetter: "ScaleOutCooldownInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetValue", GoGetter: "TargetValue"},
			_jsii_.MemberProperty{JsiiProperty: "targetValueInput", GoGetter: "TargetValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_TfScalingPlan_TargetTrackingConfigurationPropertyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktnComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktn/aws-auto-scaling-plans.TfScalingPlanConfig",
		reflect.TypeOf((*TfScalingPlanConfig)(nil)).Elem(),
	)
}
