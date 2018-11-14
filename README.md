# Demo application for ASR

## Demo

A demo of the demo app is available: http://loadtest.crete.npe:3003/

## Sequence

```sequence
Client->Backend: GET /
Backend-->Client: 200 OK/index.htm
Client->Backend: POST /asr/:id
Backend-->Client: 200 OK
Client->Backend: GET /result/:phone_number/:id
Backend-->Client: 200 OK/{}
Client->Backend: GET /result/:id
Backend-->Client: 200 OK/{}
Note right of Client: ...
Client->Backend: GET /result/:id
Backend-->Client: 200 OK/{"result": "Some text"}

```

## Endpoints

The application will have the following endpoints:

- `GET /`: send `index.html`
- `POST /asr`: call VAPI with the NCCO for the ASR scenario
- `GET /result/:id`: send the current result for the `id`
- `POST /callback`: to receive VAPI callbacks
- `GET /ncco`: return the NCCO to trigger ASR

## Running
```bash
go build ./src/app
export TOKEN=yourJWTtoken
env HOST='host:3003' VAPI_ENDPOINT='http://vapiURL/calls/' ./app
./app
```
