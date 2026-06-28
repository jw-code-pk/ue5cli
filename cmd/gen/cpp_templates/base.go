package cpp_templates

var (
	BaseHTD = `
#pragma once

#include "CoreMinimal.h"

class {{ .ApiName }} {{ .Name }} {
public:
  explicit {{ .Name }}();
  virtual ~{{ .Name }}();
};

	`

	BaseSTD = `
#include "{{ .FolderPath }}{{ .Name }}.h"

{{ .Name }}::{{ .Name }}() {}

{{ .Name}}::~{{ .Name }}() {}
	`
)
