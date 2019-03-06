const nock = require('nock')
// Requiring our app implementation
const myProbotApp = require('..')
const { Probot } = require('probot')
// Requiring our fixtures
const issueOpened = require('./fixtures/issue.opened')

nock.disableNetConnect()

describe('paper app', () => {
  let probot

  beforeEach(() => {
    probot = new Probot({})
    // Load our app into probot
    const app = probot.load(myProbotApp)

    // just return a test token
    app.app = () => 'test'
  })

  test('add paper while issue opened', async () => {
    nock('https://api.github.com')
      .post('/app/installations/2/access_tokens')
      .reply(200, { token: 'test' })

    nock('https://api.github.com')
      .get('/repos/hiimbex/testing-things/git/refs/heads/master')
      .reply(200, { object: { sha: 'abc123' } })

    nock('https://api.github.com')
      .post('/repos/hiimbex/testing-things/git/refs', {
        ref: 'refs/heads/new-branch-9999',
        sha: 'abc123'
      })
      .reply(200)

    nock('https://api.github.com')
      .put('/repos/hiimbex/testing-things/contents/path/to/your/file.md', {
        branch: 'new-branch-9999',
        message: 'adds config file',
        content: 'TXkgbmV3IGZpbGUgaXMgYXdlc29tZSE='
      })
      .reply(200)

    nock('https://api.github.com')
      .post('/repos/hiimbex/testing-things/pulls', {
        title: 'Adding my file!',
        head: 'new-branch-9999',
        base: 'master',
        body: 'Adds my new file!',
        maintainer_can_modify: true
      })
      .reply(200)

    // Recieve a webhook event
    await probot.receive({ name: 'issues', payload: issueOpened })
  })
})
