// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://localhost:8000",
        "contact": {
            "name": "Ngô Hoài Phong",
            "url": "https://www.swagger.io/support",
            "email": "hoaiphong01012002@gmai.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/get/user": {
            "get": {
                "description": "Nhận thông tin chi tiết của người dùng hiện được xác thực",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Thực hiện tìm kiếm thông tin người dùng theo ID",
                "responses": {}
            }
        },
        "/api/v1/login/email": {
            "post": {
                "description": "Thực hiện chức năng đăng nhập bằng email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Đăng nhập người dùng",
                "responses": {}
            }
        },
        "/api/v1/login/username": {
            "post": {
                "description": "Thực hiện chức năng đăng nhập bằng username",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Đăng nhập người dùng",
                "responses": {}
            }
        },
        "/api/v1/register": {
            "post": {
                "description": "Hiển thị form đăng ký cho người dùng điền thông tin",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Đăng ký người dùng trên trang web",
                "responses": {}
            }
        },
        "/api/v1/update/user": {
            "put": {
                "description": "Cập nhật thông tin người dùng",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Cập nhật thông tin người dùng",
                "responses": {}
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
	Host:             "localhost:8000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Cỏ Studio API",
	Description:      "Đây là API của Cỏ Studio,\nViệc sử dụng API này phải có sự đồng ý của Mr. Phong\nSwagger chỉ phục vụ cho việc hiển thị các API được phép sử dụng\nNếu bạn muốn sử dụng phải thực hiện trên postman\nĐể có file test postman, bạn cần liên hệ người quản lý API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
