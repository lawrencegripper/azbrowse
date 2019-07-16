package handlers

import (
	"context"
)

func NewManagementPlaneSwaggerExpander() *ManagementPlaneSwaggerExpander {
	resources := getManagementPlaneResourceTypes() // TODO need to update code-gen  to use this name!
	return &ManagementPlaneSwaggerExpander{
		swaggerExpander: NewSwaggerResourceExpander(resources),
	}
}

// ManagementPlaneSwaggerExpander expands ARM resources
type ManagementPlaneSwaggerExpander struct {
	swaggerExpander *SwaggerResourceExpander
}

// Name returns the name of the expander
func (e *ManagementPlaneSwaggerExpander) Name() string {
	return "ManagementPlaneSwaggerExpander"
}

// DoesExpand checks if this is an RG
func (e *ManagementPlaneSwaggerExpander) DoesExpand(ctx context.Context, currentItem *TreeNode) (bool, error) {
	return e.swaggerExpander.DoesExpand(ctx, currentItem)
}

// Expand returns Resources in the RG
func (e *ManagementPlaneSwaggerExpander) Expand(ctx context.Context, currentItem *TreeNode) ExpanderResult {
	return e.swaggerExpander.Expand(ctx, currentItem)
}
