import { RequiredFiledMissing } from './errors'

const action = {
  ADD: 'add',
  UPDATE: 'update',
  REMOVE: 'remove',
  IGNORE: 'ignore'
}

const requiredFields = [
  'title',
  'url',
  'source',
  'year'
]
const optionalFields = [
  'terms'
]

// getMetadata will get repo's owner and name form context.
const getMetadata = context => {
  return [
    // repo's onwer.
    context.payload.repository.owner.login,
    // repo's name.
    context.payload.repository.name
  ]
}

// parseStatus will parse action that need to take from current issue label.
const parseStatus = labels => {
  // Issue without lgtm will be ignored.
  if (labels.length === 0 || !labels.includes('lgtm')) {
    return action.IGNORE
  }

  // Handle different actions.
  if (labels.includes('action/add')) {
    return action.ADD
  }
  if (labels.includes('action/update')) {
    return action.UPDATE
  }
  if (labels.includes('action/remove')) {
    return action.REMOVE
  }
}

// parseIssue will parse issue data.
const parseIssue = content => {
  let issueData = {}

  for (let v of content.split('\n')) {
    if (v.startsWith('>')) {
      continue
    }

    let [key, ...value] = v.split(':')
    issueData[key] = value.join()
  }

  let data = {}
  for (let v of requiredFields) {
    if (issueData[v] === undefined) {
      return RequiredFiledMissing
    }

    data[v] = issueData[v]
  }
  for (let v of optionalFields) {
    data[v] = issueData[v]
  }

  return data
}

module.exports = {
  action,

  getMetadata,

  parseStatus,
  parseIssue
}
