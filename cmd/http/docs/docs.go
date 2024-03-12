// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "API Support",
            "email": "s.heidar@jabama.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/calendar": {
            "get": {
                "description": "Get",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calendar"
                ],
                "summary": "Get",
                "parameters": [
                    {
                        "type": "string",
                        "format": "YYYY-MM-DD",
                        "description": "Start date",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "format": "YYYY-MM-DD",
                        "description": "End date ",
                        "name": "end_date",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_adapter_handler_http.calendarGetResponseDto"
                        }
                    }
                }
            },
            "post": {
                "description": "Create",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calendar"
                ],
                "summary": "Create",
                "parameters": [
                    {
                        "description": "body",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/internal_adapter_handler_http.calendarCreateRequestDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internal_adapter_handler_http.calendarGetResponseDto"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "git.alibaba.ir_taraaz_salvation2_monorepo_price-service_internal_adapter_handler_http.calendarCreateRequestDto": {
            "type": "object",
            "required": [
                "days"
            ],
            "properties": {
                "days": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/git.alibaba.ir_taraaz_salvation2_monorepo_price-service_internal_adapter_handler_http.calendarDayDto"
                    }
                }
            }
        },
        "git.alibaba.ir_taraaz_salvation2_monorepo_price-service_internal_adapter_handler_http.calendarDayDto": {
            "type": "object",
            "required": [
                "date",
                "jalali",
                "type"
            ],
            "properties": {
                "date": {
                    "type": "string"
                },
                "isCustomHoliday": {
                    "type": "boolean"
                },
                "isPeak": {
                    "type": "boolean"
                },
                "jalali": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "base",
                        "weekend",
                        "holiday"
                    ]
                }
            }
        },
        "git.alibaba.ir_taraaz_salvation2_monorepo_price-service_internal_adapter_handler_http.calendarGetResponseDto": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "payload": {
                    "type": "object",
                    "properties": {
                        "days": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/git.alibaba.ir_taraaz_salvation2_monorepo_price-service_internal_adapter_handler_http.calendarDayDto"
                            }
                        }
                    }
                }
            }
        },
        "internal_adapter_handler_http.calendarCreateRequestDto": {
            "type": "object",
            "required": [
                "days"
            ],
            "properties": {
                "days": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/internal_adapter_handler_http.calendarDayDto"
                    }
                }
            }
        },
        "internal_adapter_handler_http.calendarDayDto": {
            "type": "object",
            "required": [
                "date",
                "jalali",
                "type"
            ],
            "properties": {
                "date": {
                    "type": "string"
                },
                "isCustomHoliday": {
                    "type": "boolean"
                },
                "isPeak": {
                    "type": "boolean"
                },
                "jalali": {
                    "type": "string"
                },
                "type": {
                    "type": "string",
                    "enum": [
                        "base",
                        "weekend",
                        "holiday"
                    ]
                }
            }
        },
        "internal_adapter_handler_http.calendarGetResponseDto": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "payload": {
                    "type": "object",
                    "properties": {
                        "days": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/internal_adapter_handler_http.calendarDayDto"
                            }
                        }
                    }
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Salvation2 - price service",
	Description:      "This is jabama Price service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
