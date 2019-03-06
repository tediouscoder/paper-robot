import { action, parseStatus, parseIssue, getMetadata } from './utils'
import { get, put } from './data'
import { RequiredFiledMissing } from './errors'

const handler = app => {
  app.on('issues.opened', async context => {
    let [owner, repoName] = getMetadata(context)

    let issue = await context.github.issues.get({
      owner: owner,
      name: repoName,
      number: context.payload.issue.id
    })

    switch (parseStatus(issue.data.labels)) {
      case action.IGNORE:
        context.log.info('Issue %s do not have label "lgtm", ignore.', context.payload.issue.id)
        return
      case action.ADD:
        await addPaper(context, issue)
        break
      case action.UPDATE:
        // TODO
        break
      case action.REMOVE:
        // TODO
        break
    }
  })
}

// addPaper will parse issue and add paper into repo.
const addPaper = (context, issue) => {
  context.log.info('Recevied issue with content "%s".', issue.body)

  // Parse issue.
  let issueData = parseIssue(issue.body)
  if (issueData === RequiredFiledMissing) {
    // TODO: add issue comment.
    context.log.error('Parse issue data failed for %v', issueData)
    return
  }
  context.log.info('Parsed issue data %v.', issueData)

  // Update data.toml
  let [sha, data] = get(context)

  // TODO: check if duplicate
  data['updated_at'] = new Date.now()
  // TODO: sort data by time.
  data['papers'].push(issueData)

  put(context, sha, data)

  // Generate REAMDE.
}

module.exports = handler
