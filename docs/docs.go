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
        "/v1/login": {
            "post": {
                "description": "Create user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "login"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "User credencial",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userHandler.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userService.LoginUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/v1/register_types": {
            "get": {
                "description": "Get all register types",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "register type"
                ],
                "summary": "Get all regiter types",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/registerTypeService.RegisterTypeResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/v1/transactions": {
            "get": {
                "description": "Get all transactions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Get all transactions",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/transactionService.TransactionResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            },
            "post": {
                "description": "Create transaction and send to redis queue to be aproved",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Create transaction",
                "parameters": [
                    {
                        "description": "Transaction data",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/transactionHandler.CreateTransactionRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/transactionService.CreateTransactionResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/v1/transactions/{id}": {
            "get": {
                "description": "Get transaction by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Get transaction by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "transaction id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/transactionService.GetByIdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            },
            "delete": {
                "description": "Cancel transaction and send to redis queue to be aproved",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "transaction"
                ],
                "summary": "Cancel transaction",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "transaction id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/v1/user_types": {
            "get": {
                "description": "Get all user types",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user type"
                ],
                "summary": "Get all user types",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/userTypeService.UserTypeResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/v1/user_types/{id}": {
            "get": {
                "description": "Get user type by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user type"
                ],
                "summary": "Get user type by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user type id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userTypeService.GetByIdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "description": "Get all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/userService.UserResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            },
            "post": {
                "description": "Create user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/userHandler.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userService.CreateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/v1/users/{id}": {
            "get": {
                "description": "Get users by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get users by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userService.GetByIdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/v1/wallets": {
            "get": {
                "description": "Get all wallets",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get all wallets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/walletService.WalletResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        },
        "/v1/wallets/{id}": {
            "get": {
                "description": "Get wallet by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "wallet"
                ],
                "summary": "Get wallet by id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "wallet id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/userTypeService.GetByIdResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "registerTypeService.RegisterTypeResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "transactionHandler.CreateTransactionRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "payee_id": {
                    "type": "integer"
                },
                "payer_id": {
                    "type": "integer"
                }
            }
        },
        "transactionService.CreateTransactionResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "payee_id": {
                    "type": "integer"
                },
                "payer_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "transaction_id": {
                    "type": "integer"
                }
            }
        },
        "transactionService.GetByIdResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "payee_id": {
                    "type": "integer"
                },
                "payer_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                },
                "transaction_id": {
                    "type": "integer"
                }
            }
        },
        "transactionService.TransactionResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "payee_id": {
                    "type": "integer"
                },
                "payer_id": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "userHandler.CreateUserRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "register_number": {
                    "type": "integer"
                },
                "register_type_id": {
                    "type": "integer"
                },
                "user_type_id": {
                    "type": "integer"
                }
            }
        },
        "userHandler.LoginRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "userService.CreateUserResponse": {
            "type": "object",
            "properties": {
                "expiring_in": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "userService.GetByIdResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "register_number": {
                    "type": "integer"
                },
                "register_type_id": {
                    "type": "integer"
                },
                "user_type_id": {
                    "type": "integer"
                },
                "wallet_id": {
                    "type": "integer"
                }
            }
        },
        "userService.LoginUserResponse": {
            "type": "object",
            "properties": {
                "expiring_in": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "userService.UserResponse": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "wallet_id": {
                    "type": "integer"
                }
            }
        },
        "userTypeService.GetByIdResponse": {
            "type": "object",
            "properties": {
                "user_type": {
                    "type": "string"
                },
                "user_type_id": {
                    "type": "integer"
                }
            }
        },
        "userTypeService.UserTypeResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "walletService.WalletResponse": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Technical Challenge Q2bank",
	Description:      "A transaction service API in Go using Gin framework, Redis for queue works and Postgres as relational bank",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
