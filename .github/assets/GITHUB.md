# GitHub

## Packages

With the GitHub Package Registry API, you can manage packages for your GitHub repositories and organizations.

### List packages for a user

Lists all packages in a user's namespace for which the requesting user has access.

To use this endpoint, you must authenticate using an access token with the `packages:read` scope. If `package_type` is not `container`, your token must also include the `repo` scope.

**Endpoint**

```
/users/{username}/packages
```

**Parameters**

| Name | Type | In | Description |
|---|---|---|---|
| accept | string | header | Setting to application/vnd.github.v3+json is recommended. |
| package_type | string | query | The type of supported package. Can be one of npm, maven, rubygems, nuget, docker, or container. Packages in GitHub's Gradle registry have the type maven. Docker images pushed to GitHub's Container registry (ghcr.io) have the type container. You can use the type docker to find images that were pushed to GitHub's Docker registry (docker.pkg.github.com), even if these have now been migrated to the Container registry. |
| visibility | string | query | The selected visibility of the packages. Can be one of public, private, or internal. Only container package_types currently support internal visibility properly. For other ecosystems internal is synonymous with private. This parameter is optional and only filters an existing result set. |
| username | string | query | GitHub username |

**Response**

If 200, return a list with each container package that you have in your user account.

```json
[
  {
    "id": 197,
    "name": "hello_docker",
    "package_type": "container",
    "owner": {
      "login": "monalisa",
      "id": 9919,
      "node_id": "MDEyOk9yZ2FuaXphdGlvbjk5MTk=",
      "avatar_url": "https://avatars.monalisausercontent.com/u/9919?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/monalisa",
      "html_url": "https://github.com/github",
      "followers_url": "https://api.github.com/users/github/followers",
      "following_url": "https://api.github.com/users/github/following{/other_user}",
      "gists_url": "https://api.github.com/users/github/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/github/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/github/subscriptions",
      "organizations_url": "https://api.github.com/users/github/orgs",
      "repos_url": "https://api.github.com/users/github/repos",
      "events_url": "https://api.github.com/users/github/events{/privacy}",
      "received_events_url": "https://api.github.com/users/github/received_events",
      "type": "User",
      "site_admin": false
    },
    "version_count": 1,
    "visibility": "private",
    "url": "https://api.github.com/orgs/github/packages/container/hello_docker",
    "created_at": "2020-05-19T22:19:11Z",
    "updated_at": "2020-05-19T22:19:11Z",
    "html_url": "https://github.com/orgs/github/packages/container/package/hello_docker"
  },
  {
    "id": 198,
    "name": "goodbye_docker",
    "package_type": "container",
    "owner": {
      "login": "github",
      "id": 9919,
      "node_id": "MDEyOk9yZ2FuaXphdGlvbjk5MTk=",
      "avatar_url": "https://avatars.githubusercontent.com/u/9919?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/monalisa",
      "html_url": "https://github.com/github",
      "followers_url": "https://api.github.com/users/github/followers",
      "following_url": "https://api.github.com/users/github/following{/other_user}",
      "gists_url": "https://api.github.com/users/github/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/github/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/github/subscriptions",
      "organizations_url": "https://api.github.com/users/github/orgs",
      "repos_url": "https://api.github.com/users/github/repos",
      "events_url": "https://api.github.com/users/github/events{/privacy}",
      "received_events_url": "https://api.github.com/users/github/received_events",
      "type": "User",
      "site_admin": false
    },
    "version_count": 2,
    "visibility": "private",
    "url": "https://api.github.com/user/monalisa/packages/container/goodbye_docker",
    "created_at": "2020-05-20T22:19:11Z",
    "updated_at": "2020-05-20T22:19:11Z",
    "html_url": "https://github.com/user/monalisa/packages/container/package/goodbye_docker"
  }
]
```

Golang struct to render this output.

```Golang
type ContainerPackage struct {
	ID         int
	Name       string
	Owner      string
	Visibility string
	CreatedAt  time.Time
}
```

### Get all package versions for a package owned by a user

Returns all package versions for a public package owned by a specified user.

To use this endpoint, you must authenticate using an access token with the packages:read scope. If package_type is not container, your token must also include the repo scope.

**Endpoint**

```
/users/{username}/packages/{package_type}/{package_name}/versions
```

**Parameters**

| Name | Type | In | Description |
|---|---|---|---|
| accept | string | header | Setting to application/vnd.github.v3+json is recommended. |
| package_type | string | query | The type of supported package. Can be one of npm, maven, rubygems, nuget, docker, or container. Packages in GitHub's Gradle registry have the type maven. Docker images pushed to GitHub's Container registry (ghcr.io) have the type container. You can use the type docker to find images that were pushed to GitHub's Docker registry (docker.pkg.github.com), even if these have now been migrated to the Container registry. |
| package_name | string | query | The name of the package. |
| username | string | query | GitHub username |

**Response**

If 200, return a list with each container package that you have in your user account.

```json
[
  {
    "id": 3497268,
    "name": "0.3.0",
    "url": "https://api.github.com/users/octocat/packages/rubygems/octo-name/versions/3497268",
    "package_html_url": "https://github.com/octocat/octo-name-repo/packages/40201",
    "license": "MIT",
    "created_at": "2020-08-31T15:22:11Z",
    "updated_at": "2020-08-31T15:22:12Z",
    "description": "Project for octocats",
    "html_url": "https://github.com/octocat/octo-name-repo/packages/40201?version=0.3.0",
    "metadata": {
      "package_type": "rubygems"
    }
  },
  {
    "id": 387039,
    "name": "0.2.0",
    "url": "https://api.github.com/users/octocat/packages/rubygems/octo-name/versions/387039",
    "package_html_url": "https://github.com/octocat/octo-name-repo/packages/40201",
    "license": "MIT",
    "created_at": "2019-12-01T20:49:29Z",
    "updated_at": "2019-12-01T20:49:30Z",
    "description": "Project for octocats",
    "html_url": "https://github.com/octocat/octo-name-repo/packages/40201?version=0.2.0",
    "metadata": {
      "package_type": "rubygems"
    }
  },
  {
    "id": 169770,
    "name": "0.1.0",
    "url": "https://api.github.com/users/octocat/packages/rubygems/octo-name/versions/169770",
    "package_html_url": "https://github.com/octocat/octo-name-repo/packages/40201",
    "license": "MIT",
    "created_at": "2019-10-20T14:17:14Z",
    "updated_at": "2019-10-20T14:17:15Z",
    "html_url": "https://github.com/octocat/octo-name-repo/packages/40201?version=0.1.0",
    "metadata": {
      "package_type": "rubygems"
    }
  }
]
```
