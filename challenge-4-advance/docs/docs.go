// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/book": {
            "get": {
                "description": "Get all Books on the system",
                "tags": [
                    "book"
                ],
                "summary": "Get all Books",
                "parameters": [
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "page",
                        "name": "page",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/microservices_challenge-4-advance_application_usecases_book.PaginationResultBook"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/book.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/book.MessageResponse"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "book"
                ],
                "summary": "Create New Book",
                "parameters": [
                    {
                        "description": "body data",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/book.NewBookRequest"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Accepted",
                        "schema": {
                            "$ref": "#/definitions/microservices_challenge-4-advance_domain_book.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/book.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/book.MessageResponse"
                        }
                    }
                }
            }
        },
        "/book/{book_id}": {
            "get": {
                "tags": [
                    "book"
                ],
                "summary": "Get books by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "id of book",
                        "name": "book_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/microservices_challenge-4-advance_domain_book.Book"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/book.MessageResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/book.MessageResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "book.MessageResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "book.NewBookRequest": {
            "type": "object",
            "required": [
                "author",
                "description",
                "title"
            ],
            "properties": {
                "author": {
                    "type": "string",
                    "example": "Roche"
                },
                "description": {
                    "type": "string",
                    "example": "Something"
                },
                "title": {
                    "type": "string",
                    "example": "Book"
                }
            }
        },
        "microservices_challenge-4-advance_application_usecases_book.PaginationResultBook": {
            "type": "object",
            "properties": {
                "current": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/microservices_challenge-4-advance_domain_book.Book"
                    }
                },
                "limit": {
                    "type": "integer"
                },
                "nextCursor": {
                    "type": "integer"
                },
                "numPages": {
                    "type": "integer"
                },
                "prevCursor": {
                    "type": "integer"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "microservices_challenge-4-advance_domain_book.Book": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string",
                    "example": "mr. author"
                },
                "createdAt": {
                    "type": "string"
                },
                "deletedAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string",
                    "example": "book description"
                },
                "id": {
                    "type": "integer"
                },
                "title": {
                    "type": "string",
                    "example": "book title"
                },
                "updatedAt": {
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
