# Pro Cris Server

## Steps

- 1 Install [air](https://github.com/cosmtrek/air)
- 2 Install [templ](https://templ.guide/quick-start/installation)
- 3 Use `make dev` to run the server in development mode.
- 4 Use `make test` to start runing unit tests.

## Endpoints

- GET /api/v1/checkhealth
- POST /api/v1/auth/login
- Middleware /api/v1/auth/ensure-authenticated
- POST /api/v1/student
- PUT /api/v1/student/:id
- PUT /api/v1/student/:id/picture
- GET /api/v1/student
- GET /api/v1/student/:id
- DELETE /api/v1/student/:id
- POST /api/v1/appointment
- PUT /api/v1/appointment/:id
- GET /api/v1/appointment
- GET /api/v1/appointment/:id
- DELETE /api/v1/appointment/:id
