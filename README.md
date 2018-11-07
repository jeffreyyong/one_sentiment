# Demo application for ASR

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
- `PUT /asr/:phone_number/:id`: call VAPI with the NCCO for the ASR scenario
- `GET /result/:id`: send the current result for the `id`
- `POST /callback`: to receive VAPI callbacks