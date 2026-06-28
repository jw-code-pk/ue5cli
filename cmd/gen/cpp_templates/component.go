package cpp_templates

var (
	// Component Templates
	ComponentHTD = `
#pragma once

#include "CoreMinimal.h"
#include "Components/ActorComponent.h"
#include "{{.Name}}.generated.h"

UCLASS( ClassGroup=(Custom), meta=(BlueprintSpawnableComponent) )
class {{.ApiName}} U{{.Name}} : public UActorComponent
{
	GENERATED_BODY()

public:	
	U{{.Name}}();

protected:
	virtual void BeginPlay() override;

public:	
	virtual void TickComponent(float DeltaTime, ELevelTick TickType, FActorComponentTickFunction* ThisTickFunction) override;
};
	`

	ComponentSTD = `
#include "{{.FolderPath}}{{.Name}}.h"

U{{.Name}}::U{{.Name}}()
{
}


void U{{.Name}}::BeginPlay()
{
	Super::BeginPlay();
}


void U{{.Name}}::TickComponent(float DeltaTime, ELevelTick TickType, FActorComponentTickFunction* ThisTickFunction)
{
	Super::TickComponent(DeltaTime, TickType, ThisTickFunction);
}
	`
)
