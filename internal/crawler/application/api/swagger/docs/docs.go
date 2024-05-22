// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
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
        "/crawl": {
            "post": {
                "description": "Crawl website",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Crawler"
                ],
                "summary": "Crawl website",
                "parameters": [
                    {
                        "description": "Website Info",
                        "name": "arg",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.WebSite"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Entry"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/sources": {
            "get": {
                "description": "get source by Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Source"
                ],
                "summary": "get sources",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Source"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "put": {
                "description": "Create Source",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Source"
                ],
                "summary": "Create Source",
                "parameters": [
                    {
                        "description": "Source Info",
                        "name": "arg",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Source"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Source"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            },
            "post": {
                "description": "Create Source",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Source"
                ],
                "summary": "Create Source",
                "parameters": [
                    {
                        "description": "Source Info",
                        "name": "arg",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Source"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Source"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        },
        "/sources/{id}": {
            "get": {
                "description": "get source by Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Source"
                ],
                "summary": "get an source",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Source"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {}
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {}
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {}
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Entry": {
            "type": "object",
            "additionalProperties": {}
        },
        "domain.Page": {
            "type": "object",
            "properties": {
                "page_events": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.PageEvent"
                    }
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "domain.PageEvent": {
            "type": "object",
            "properties": {
                "collectors": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.PageObject"
                    }
                },
                "enter_value": {
                    "type": "string"
                },
                "order": {
                    "type": "integer"
                },
                "selector": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "domain.PageObject": {
            "type": "object",
            "properties": {
                "key": {
                    "type": "string"
                },
                "page_objects": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.PageObject"
                    }
                },
                "regex_extract": {
                    "type": "string"
                },
                "selector": {
                    "type": "string"
                }
            }
        },
        "domain.Source": {
            "type": "object",
            "properties": {
                "data": {
                    "$ref": "#/definitions/domain.WebSite"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "domain.WebSite": {
            "type": "object",
            "properties": {
                "pages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Page"
                    }
                },
                "url": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BasicAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Crawler API",
	Description:      "This is a Crawler API server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
