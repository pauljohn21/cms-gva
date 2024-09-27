package global

{{- if .HasGlobal }}

import "github.com/pauljohn21/cms-gva/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}