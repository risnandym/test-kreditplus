// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

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
        "/kredit-plus/customer/login": {
            "post": {
                "description": "Logging in to get jwt token to access admin or user api by roles.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login.",
                "parameters": [
                    {
                        "description": "the body to login a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contract.LoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/kredit-plus/customer/profile": {
            "post": {
                "security": [
                    {
                        "kreditplus-token": []
                    }
                ],
                "description": "Save Customer Profile.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Customer"
                ],
                "summary": "Create Profile.",
                "parameters": [
                    {
                        "description": "the body to create a new Profile",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contract.ProfileInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/contract.ProfileInput"
                        }
                    }
                }
            }
        },
        "/kredit-plus/customer/register": {
            "post": {
                "description": "registering a user from public access.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register a User.",
                "parameters": [
                    {
                        "description": "the body to register a user",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contract.RegisterInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/news": {
            "post": {
                "security": [
                    {
                        "kreditplus-token": []
                    }
                ],
                "description": "Creating a new Transaction.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transaction"
                ],
                "summary": "Create New Transaction. (Admin only)",
                "parameters": [
                    {
                        "description": "the body to create a new transaction",
                        "name": "Body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/contract.TransactionInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/contract.TransactionOutput"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "contract.LoginInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "contract.ProfileInput": {
            "type": "object",
            "required": [
                "date_of_birth",
                "full_name",
                "ktp_image",
                "legal_name",
                "nik",
                "place_of_birth",
                "salary",
                "selfie_image"
            ],
            "properties": {
                "auth_id": {
                    "type": "integer"
                },
                "date_of_birth": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "ktp_image": {
                    "type": "string"
                },
                "legal_name": {
                    "type": "string"
                },
                "nik": {
                    "type": "string"
                },
                "place_of_birth": {
                    "type": "string"
                },
                "salary": {
                    "type": "number"
                },
                "selfie_image": {
                    "type": "string"
                }
            }
        },
        "contract.RegisterInput": {
            "type": "object",
            "required": [
                "email",
                "password",
                "phone"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "contract.TransactionInput": {
            "type": "object",
            "properties": {
                "admin_fee": {
                    "type": "number"
                },
                "asset_id": {
                    "type": "integer"
                },
                "auth_id": {
                    "type": "integer"
                },
                "contract_number": {
                    "type": "string"
                },
                "installment_amount": {
                    "type": "number"
                },
                "installment_period": {
                    "type": "integer"
                },
                "interest_amount": {
                    "type": "number"
                },
                "otr_amount": {
                    "type": "number"
                },
                "sales_channel": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "contract.TransactionOutput": {
            "type": "object",
            "properties": {
                "admin_fee": {
                    "type": "number"
                },
                "asset_id": {
                    "type": "integer"
                },
                "auth_id": {
                    "type": "integer"
                },
                "contract_number": {
                    "type": "string"
                },
                "installment_amount": {
                    "type": "number"
                },
                "installment_period": {
                    "type": "integer"
                },
                "interest_amount": {
                    "type": "number"
                },
                "limit": {
                    "$ref": "#/definitions/entities.Limit"
                },
                "otr_amount": {
                    "type": "number"
                },
                "sales_channel": {
                    "type": "string"
                },
                "uuid": {
                    "type": "string"
                }
            }
        },
        "entities.Limit": {
            "type": "object",
            "properties": {
                "auth_id": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "deleted_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "tenor1": {
                    "type": "number"
                },
                "tenor2": {
                    "type": "number"
                },
                "tenor3": {
                    "type": "number"
                },
                "tenor4": {
                    "type": "number"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "kreditplus-token": {
            "description": "Value is: \"Bearer {access_token}\", where access_token is retrieved from cms-service/v1/login",
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
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Kredit+",
	Description:      "This is a Kredit+ test.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
