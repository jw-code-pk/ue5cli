package cpp_templates

var (
	// Actor Templates
	ActorHTD = `
#pragma once

#include "CoreMinimal.h"
#include "GameFramework/Actor.h"
#include "{{ .Name }}.generated.h" 

UCLASS()
class {{ .ApiName }} A{{ .Name }} : public AActor
{
	GENERATED_BODY()
			
public:
	A{{ .Name }} ();
};
	`

	ActorSTD = `
#include "{{ .FolderPath }}{{ .Name }}.h"

A{{ .Name }}::A{{ .Name }}() {}
  `
)
