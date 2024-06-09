// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/v1/admin/recalculate": {
            "post": {
                "description": "Start recalculating matches",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Recalculate matches",
                "responses": {
                    "200": {
                        "description": "recalculation done",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/mmr/matches": {
            "get": {
                "description": "Get all matches",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Matches"
                ],
                "summary": "Get matches",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/view.MatchDetails"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Submit a match for MMR calculation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Matches"
                ],
                "summary": "Submit a match",
                "parameters": [
                    {
                        "description": "Match object",
                        "name": "match",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.Match"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "match submitted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/v1/profile": {
            "get": {
                "description": "Get profile details of the authenticated user",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Get profile",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.ProfileDetails"
                        }
                    }
                }
            }
        },
        "/v1/profile/claim": {
            "post": {
                "description": "Claims a user by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Claim user",
                "parameters": [
                    {
                        "description": "User to be claimed",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.ClaimUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.ProfileDetails"
                        }
                    }
                }
            }
        },
        "/v1/stats/leaderboard": {
            "get": {
                "description": "Get leaderboard stats including wins, loses, and MMR of users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Leaderboard"
                ],
                "summary": "Get leaderboard stats",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/repos.LeaderboardEntry"
                            }
                        }
                    }
                }
            }
        },
        "/v1/stats/player-history": {
            "get": {
                "description": "Get player history including MMR and date",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Statistics"
                ],
                "summary": "Get player history",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "userId",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Start date",
                        "name": "start",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "End date",
                        "name": "end",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/view.PlayerHistoryDetails"
                            }
                        }
                    }
                }
            }
        },
        "/v1/users": {
            "get": {
                "description": "Lists all users",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "List users",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/view.UserDetails"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Creates a new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/view.UserDetails"
                        }
                    }
                }
            }
        },
        "/v1/users/search": {
            "get": {
                "description": "Searches users by name",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Search users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name to search for",
                        "name": "query",
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
                                "$ref": "#/definitions/view.UserDetails"
                            }
                        }
                    }
                }
            }
        },
        "/v2/mmr/matches": {
            "get": {
                "description": "Get all matches",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Matches"
                ],
                "summary": "Get matches",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/view.MatchDetailsV2"
                            }
                        }
                    }
                }
            },
            "post": {
                "description": "Submit a match for MMR calculation",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Matches"
                ],
                "summary": "Submit a match",
                "parameters": [
                    {
                        "description": "Match object",
                        "name": "match",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/view.MatchV2"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "match submitted",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "repos.LeaderboardEntry": {
            "type": "object",
            "properties": {
                "loses": {
                    "type": "integer"
                },
                "mmr": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "wins": {
                    "type": "integer"
                }
            }
        },
        "view.ClaimUser": {
            "type": "object",
            "required": [
                "userId"
            ],
            "properties": {
                "userId": {
                    "type": "integer"
                }
            }
        },
        "view.CreateUser": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "displayName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "view.Match": {
            "type": "object",
            "required": [
                "team1",
                "team2"
            ],
            "properties": {
                "team1": {
                    "$ref": "#/definitions/view.MatchTeam"
                },
                "team2": {
                    "$ref": "#/definitions/view.MatchTeam"
                }
            }
        },
        "view.MatchDetails": {
            "type": "object",
            "required": [
                "date",
                "team1",
                "team2"
            ],
            "properties": {
                "date": {
                    "type": "string"
                },
                "mmrCalculations": {
                    "$ref": "#/definitions/view.MatchMMRCalculationDetails"
                },
                "team1": {
                    "$ref": "#/definitions/view.MatchTeam"
                },
                "team2": {
                    "$ref": "#/definitions/view.MatchTeam"
                }
            }
        },
        "view.MatchDetailsV2": {
            "type": "object",
            "required": [
                "date",
                "team1",
                "team2"
            ],
            "properties": {
                "date": {
                    "type": "string"
                },
                "mmrCalculations": {
                    "$ref": "#/definitions/view.MatchMMRCalculationDetails"
                },
                "team1": {
                    "$ref": "#/definitions/view.MatchTeamV2"
                },
                "team2": {
                    "$ref": "#/definitions/view.MatchTeamV2"
                }
            }
        },
        "view.MatchMMRCalculationDetails": {
            "type": "object",
            "required": [
                "team1",
                "team2"
            ],
            "properties": {
                "team1": {
                    "$ref": "#/definitions/view.MatchMMRCalculationTeam"
                },
                "team2": {
                    "$ref": "#/definitions/view.MatchMMRCalculationTeam"
                }
            }
        },
        "view.MatchMMRCalculationTeam": {
            "type": "object",
            "required": [
                "player1MMRDelta",
                "player2MMRDelta"
            ],
            "properties": {
                "player1MMRDelta": {
                    "type": "integer"
                },
                "player2MMRDelta": {
                    "type": "integer"
                }
            }
        },
        "view.MatchTeam": {
            "type": "object",
            "required": [
                "member1",
                "member2",
                "score"
            ],
            "properties": {
                "member1": {
                    "type": "string"
                },
                "member2": {
                    "type": "string"
                },
                "score": {
                    "type": "integer"
                }
            }
        },
        "view.MatchTeamV2": {
            "type": "object",
            "required": [
                "member1",
                "member2",
                "score"
            ],
            "properties": {
                "member1": {
                    "type": "integer"
                },
                "member2": {
                    "type": "integer"
                },
                "score": {
                    "type": "integer"
                }
            }
        },
        "view.MatchV2": {
            "type": "object",
            "required": [
                "team1",
                "team2"
            ],
            "properties": {
                "team1": {
                    "$ref": "#/definitions/view.MatchTeamV2"
                },
                "team2": {
                    "$ref": "#/definitions/view.MatchTeamV2"
                }
            }
        },
        "view.PlayerHistoryDetails": {
            "type": "object",
            "required": [
                "date",
                "mmr",
                "name",
                "userId"
            ],
            "properties": {
                "date": {
                    "type": "string"
                },
                "mmr": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "view.ProfileDetails": {
            "type": "object",
            "properties": {
                "userId": {
                    "type": "integer"
                }
            }
        },
        "view.UserDetails": {
            "type": "object",
            "required": [
                "name",
                "userId"
            ],
            "properties": {
                "displayName": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
