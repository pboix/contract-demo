const Form = require('./form')

const formId = 'abcdef'

describe('Consumer is tested', () => {

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

    it('Returns responses for a form', done => {
        Form.getResponses(formId)
            .then(responses => {
                console.log(`IT RETURNS: ${responses}`)
                console.log(responses.title)
                expect(responses).toEqual(EXPECTED_BODY)
            })
            .then(done)
    })
})