const { Pact } = require('@pact-foundation/pact')
const Form = require('./form')
const MOCK_SERVER_PORT = 2202;
const formId = 'abcdef'

describe('Consumer is tested', () => {
    const provider = new Pact({
        consumer: "ResultsRenderer",
        provider: "ResultsAPI",
        port: MOCK_SERVER_PORT,
        dir: process.cwd() + '/pacts/',
        spec: 2
    });

    const EXPECTED_BODY =
    {
        "title": "My Awesome Form",
        "form": [
            {
                "token": "21085286190ffad1248d17c4135ee56f",
                "answers": [
                    {
                        "field": {
                            "id": "hVONkQcnSNRj",
                            "type": "short_text",
                            "title": "What's your favourite cheese?"
                        },
                        "type": "text",
                        "answer": "Gouda"
                    },
                    {
                        "field": {
                            "id": "RUqkXSeXBXSd",
                            "type": "yes_no",
                            "title": "do you eat cheese everyday?"
                        },
                        "type": "boolean",
                        "answer": true
                    }
                ]
            }
        ]
    }

    beforeAll(done => {
        provider.setup()
            .then(() => {
                provider.addInteraction({
                    uponReceiving: "a request for form responses",
                    withRequest: {
                        method: "GET",
                        path: "/forms/abcdef/responses",
                        headers: { Accept: "application/json" }
                    },
                    willRespondWith: {
                        status: 200,
                        headers: { "Content-Type": "application/json" },
                        body: EXPECTED_BODY
                    }
                });
            })
            .then(done);
    })

    it('Returns responses for a form', done => {
        Form.getResponses(formId)
            .then(responses => {
                provider.verify()
                expect(responses.title).toBe('My Awesome Form')
            })
            .then(done)
    })

    afterAll(() => {
        return provider.finalize();
    })
})