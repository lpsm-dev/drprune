## ➤ Description <a name = "description"></a>

### Considerations about GitHub

- A user can have a list of packages.
  - Each package have a type: [container, maven, npm].
  - Each package have versions.
    - Each version of a package can have a name.
    - Each version of a package can be tagged or not.

### Considerations about GitLab

- GitLab uses a modularized organization with group and project concepts.
  - Group is a collection on projects and sub-groups.
  - Project is a git repository.
- GitLab organize te package idea between:
  - Package Registry
  - Container Registry
  - Infrastructure Registry
  - Dependency Proxy.
- In this case we focus in Container Registry.
  - You can view the Container Registry for a project or group.
  - Go to Packages & Registries > Container Registry.
  - Only members of the project or group can access a private project’s Container Registry.
  - Images follow this naming convention: `<registry URL>/<namespace>/<project>/<image>`
  - A namespace is a name of group or project.

**Project**

```txt
> Location: registry.gitlab.com/example/api/chip/develop
  > Tags count: 7
> Location: registry.gitlab.com/example/api/chip/develop/cache
  > Tags count: 2
> Location: registry.gitlab.com/example/api/chip/cache
  > Tags count: 17
> Location: registry.gitlab.com/example/api/chip
  > Tags count: 29
> Location: registry.gitlab.com/example/api/chip/feature-audit
  > Tags count: 88
> Location: registry.gitlab.com/example/api/chip/feature-audit/cache
  > Tags count: 46
```

**Group**

```txt
> Location: registry.gitlab.com/example/one/api/recarga/develop/cache
  > Tags count: 88
> Location: registry.gitlab.com/example/one/api/chip/develop
  > Tags count: 12
> Location: registry.gitlab.com/example/one/api/auth/feature-sms-send
  > Tags count: 2
> Location: registry.gitlab.com/example/one/api/auth/feature-sms-send/cache
  > Tags count: 4
> Location: registry.gitlab.com/example/one/api/user/feature-add-users
  > Tags count: 10
> Location: registry.gitlab.com/example/two/api/chip/develop
  > Tags count: 423
> Location: registry.gitlab.com/example/two/api/chip
  > Tags count: 663
> Location: registry.gitlab.com/example/two/api/portability
  > Tags count: 98
> Location: registry.gitlab.com/example/two/api/portability/cache
  > Tags count: 4
> Location: registry.gitlab.com/example/two/api/recharge/cache
  > Tags count: 16
```
