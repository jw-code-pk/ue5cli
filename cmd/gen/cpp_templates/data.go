package cpp_templates

var (
	// Component Templates
	DataHTD = `
#pragma once

#include "Engine/DataAsset.h"
#include "{{.Name}}.generated.h"

UCLASS(BlueprintType)
class {{.ApiName}} U{{.Name}} : public UDataAsset
{
	GENERATED_BODY()
};
`

	DataSTD = `
#include "{{.FolderPath}}{{.Name}}.h"
`
)
