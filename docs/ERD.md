#### TODOES

## Requirements
1. Creating todo
    * POST /api/v1/todoes
    * request body 
        ```json
        {
            "name":"test todo",
            "description":"this todo is ...",
            "start_time":"2023-01-25T17:32:40Z+0500",
            "expiration_time":"2023-01-30T17:32:40Z+0500",
        }
        ```
    * response body
        ```json
        {
            "code": 201,
            "message":"success"
        }
        ```
2. Reading todo (by ID)
    * GET /api/v1/todoes?id=13
    * response body
        ```json
        {
            "id": 13,
            "name":"test todo",
            "description":"this todo is ...",
            "start_time":"2023-01-25T17:32:40Z+0500",
            "expiration_time":"2023-01-30T17:32:40Z+0500"
        }
        ```
2. Reading todoes (All todoes)
    * GET /api/v1/todoes
    * response body
        ```json
        {
        "code": 200,
        "message":"success",
        "todoes": [
            {
                "id": 1,
                "name":"test todo",
                "description":"this todo is ...",
                "start_time":"2023-01-25T17:32:40Z+0500",
                "expiration_time":"2023-01-30T17:32:40Z+0500"
            },
            {
                "id": 2,
                "name":"test todo",
                "description":"this todo is ...",
                "start_time":"2023-01-25T17:32:40Z+0500",
                "expiration_time":"2023-01-30T17:32:40Z+0500"
            }
        ]
        }
        ```
3. Updating todo
    * PUT /api/v1/todoes
    * request body 
        ```json
        {
            "name":"test todo",
            "description":"this todo is ...",
            "start_time":"2023-01-25T17:32:40Z+0500",
            "expiration_time":"2023-01-30T17:32:40Z+0500"
        }
        ```
    * response body
        ```json
        {
            "code": 200,
            "message":"success"
        }
        ```
4. Soft-deleting todo
    * DELETE /api/v1/todoes?id=13
    * response body
        ```json
        {
            "code": 200,
            "message":"success"
        }
        ```