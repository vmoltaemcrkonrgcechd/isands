// Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/category": {
            "get": {
                "tags": [
                    "категория"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.DictionaryEntry"
                            }
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "категория"
                ],
                "parameters": [
                    {
                        "description": "категория",
                        "name": "dto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.DictionaryEntryDTO"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/category/{id}": {
            "delete": {
                "tags": [
                    "категория"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "идентификатор",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "patch": {
                "tags": [
                    "категория"
                ],
                "parameters": [
                    {
                        "description": "категория",
                        "name": "dto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.DictionaryEntryDTO"
                        }
                    },
                    {
                        "type": "string",
                        "description": "идентификатор",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/country": {
            "get": {
                "tags": [
                    "страна"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.DictionaryEntry"
                            }
                        }
                    }
                }
            },
            "post": {
                "tags": [
                    "страна"
                ],
                "parameters": [
                    {
                        "description": "страна",
                        "name": "dto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.DictionaryEntryDTO"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/country/{id}": {
            "delete": {
                "tags": [
                    "страна"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "идентификатор",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "patch": {
                "tags": [
                    "страна"
                ],
                "parameters": [
                    {
                        "description": "страна",
                        "name": "dto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.DictionaryEntryDTO"
                        }
                    },
                    {
                        "type": "string",
                        "description": "идентификатор",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "entity.DictionaryEntry": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entity.DictionaryEntryDTO": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
