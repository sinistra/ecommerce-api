# eco-api
This is an eCommerce REST API written in go.
A number of middlewares have been defined for logging, recovery & auth.

Auth is done via JWTs and processed by the middleware, so every call to a REST endpoint can be validated.
Endpoints are grouped, so they can optionally be subjected to auth or allowed to proceed without auth (ie to login).

The underlying db is Mongo and the services have been separated so that MongoDb calls can be individually tested and optimized.
The MVC pattern has been implemented. I'm yet to be convinced that the Go pattern of putting routes & handlers into an 'internal' folder is the best way to go.
IMO this provokes a pre-emptive conversation on microservice architecture. Once the architecture is agreed upon, then (maybe) the Go pattern is more appropriate.
