const { Pact } = require('@pact-foundation/pact')
const addInteractions = require('../pact_support/addInteractions')
const Form = require('./form')
let provider

describe('Consumer is tested', () => {
    beforeAll(done => {
        provider = new Pact({
            consumer: "ResultsRenderer",
            provider: "ResultsAPI",
            port: 2202,
            dir: process.cwd() + '/pacts/',
            spec: 2
        });

        provider.setup()
            .then(() => addInteractions(provider))
            .then(done)
    })

    it('Returns responses for a form', done => {
        const formId = 'abcdef'
        Form.getResponses(formId)
            .then(responses => {
                provider.verify()
                expect(responses.title).toBe('My Awesome Form')
            })
            .then(done)
    })

    afterAll(() => provider.finalize())
})