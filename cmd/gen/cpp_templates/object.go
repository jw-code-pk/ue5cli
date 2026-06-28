package cpp_templates

var (
	// Component Templates
	ObjectHTD = `
#pragma once

#include "CoreMinimal.h"
#include "UObject/NoExportTypes.h"
#include "{{.Name}}.generated.h"

UCLASS(Abstract, Blueprintable, EditInlineNew, CollapseCategories, HideCategories=(Object))
class {{.ApiName}} U{{.Name}} : public UObject
{
	GENERATED_BODY()

public:
	U{{.Name}}();
};`

	ObjectSTD = `
#include "{{.FolderPath}}{{.Name}}.h"

U{{.Name}}::U{{.Name}}()
{
}
`
)
