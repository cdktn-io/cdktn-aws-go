//go:build no_runtime_type_checking

package glue

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsMlTransform_SchemaPropertyList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_AwsMlTransform_SchemaPropertyList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_AwsMlTransform_SchemaPropertyList) validateResolveParameters(context cdktn.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_AwsMlTransform_SchemaPropertyList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_AwsMlTransform_SchemaPropertyList) validateSetTerraformResourceParameters(val cdktn.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_AwsMlTransform_SchemaPropertyList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewAwsMlTransform_SchemaPropertyListParameters(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

