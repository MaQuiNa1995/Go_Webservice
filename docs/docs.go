// Package docs GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import (
	"bytes"
	"encoding/json"
	"strings"
	"text/template"

	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "MaQuiNa1995",
            "url": "https://github.com/MaQuiNa1995",
            "email": "maquina1995@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/cancion": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Busca todos los registros de la tabla",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.CancionEntity"
                            }
                        }
                    }
                }
            },
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Crea un nuevo registro en Base de datos",
                "parameters": [
                    {
                        "description": "Dto con los datos necesarios para crear una entidad",
                        "name": "cancion",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CancionCreateDto"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    }
                }
            },
            "patch": {
                "produces": [
                    "application/json"
                ],
                "summary": "Actualiza un registro en Base de datos por id",
                "parameters": [
                    {
                        "description": "Dto con los datos necesarios para actualizar por id",
                        "name": "cancion",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CancionUpdateDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/cancion/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Buscar por Id en Base de datos",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id de la entidad a buscar",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.CancionEntity"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "summary": "Eliminar por Id en Base de datos",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id de la base de datos a eliminar",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "202": {
                        "description": ""
                    },
                    "204": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "model.CancionCreateDto": {
            "type": "object",
            "required": [
                "duracion",
                "genero",
                "nombre"
            ],
            "properties": {
                "duracion": {
                    "type": "number"
                },
                "genero": {
                    "type": "string"
                },
                "nombre": {
                    "type": "string"
                }
            }
        },
        "model.CancionEntity": {
            "type": "object",
            "properties": {
                "duracion": {
                    "type": "number"
                },
                "genero": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nombre": {
                    "type": "string"
                }
            }
        },
        "model.CancionUpdateDto": {
            "type": "object",
            "properties": {
                "duracion": {
                    "type": "number"
                },
                "genero": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "nombre": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
		"escape": func(v interface{}) string {
			// escape tabs
			str := strings.Replace(v.(string), "\t", "\\t", -1)
			// replace " with \", and if that results in \\", replace that with \\\"
			str = strings.Replace(str, "\"", "\\\"", -1)
			return strings.Replace(str, "\\\\\"", "\\\\\\\"", -1)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
