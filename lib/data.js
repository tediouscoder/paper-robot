import { getMetadata } from './utils'

const dataFilePath = 'src/data.json'

const get = async context => {
  let [owner, repoName] = getMetadata(context)

  let result = await context.github.getContents({
    owner: owner,
    repo: repoName,
    path: dataFilePath,
  })

  return [result.sha, JSON.parse(Buffer.from(result.content, 'base64'))]
}

const put = async (context, sha, data) => {
  let [owner, repoName] = getMetadata(context)

  let result = await context.github.updateFile({
    owner: owner,
    repo: repoName,
    path: dataFilePath,
    message: 'Update data.json via issue #' + context.payload.issue.id,
    content: Buffer.from(JSON.stringify(data)).toString('base64'),
    sha: sha,
  })
}